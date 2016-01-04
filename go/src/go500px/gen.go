package go500px

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"regexp"
	"text/template"
)

// walkPackages scans the current directory for files which require code generation.
// These go files are interfaces with annotations.
func walkPackages(root string) (*requestInfo, error) {
	fileset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fileset, root, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Failed to parse dir for packages. Reason: %v", err)
		return nil, err
	}

	var files []*ast.File
	for _, ast := range pkgs {
		for _, f := range ast.Files {
			files = append(files, f)
		}
	}

	config := &types.Config{
		Error: func(e error) {
			log.Print(e)
		},
		Importer: importer.Default(),
	}

	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}

	_, err = config.Check(root, fileset, files, info)
	if err != nil {
		log.Fatalf("One or more packages failed type checking: %v", err)
	}

	visitor := newFileASTVisitor(info)
	for _, f := range files {
		ast.Walk(visitor, f)
	}

	return visitor.reqInfo, nil
}

// generateBuilder takes the scanned files that meet the generation requirements and outputs the generated
// go file that implements the respective interface.
func generateBuilder(r *requestInfo) {
	r.Pkg = "go500px"

	funcMap := template.FuncMap{
		"ParamsList": getParamsList,
		"ParamName":  getParamName,
		"ParamKey":   getParamKey,
	}

	var builderTemplate = template.Must(template.New("builder").Funcs(funcMap).Parse(`/*
* CODE GENERATED AUTOMATICALLY WITH Go500px
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package {{.Pkg}}

import (
	"net/http"
	"net/url"
)

type {{ .Name }}Impl struct {
	queryParams url.Values
}

func New{{ .Name }}(url string) {{ .Name }} {
	return &{{ .Name }}Impl{}
}

{{ range $key, $value := .QueryParams }}
func (b *{{ $.Name }}Impl) {{ $key }}({{ $value.Type | ParamsList }}) {{ $.Name }} {
	b.queryParams.Add("{{ $value | ParamKey }}", {{ $value.Type | ParamName }})
	return b
}
{{ end }}

func (b *{{ .Name }}Impl) build() (*http.Request, error) {
	req, err := http.NewRequest("{{ .HttpMethod }}", "{{ .Api }}", nil)
	if err != nil {
		return nil, err
	}
	if len(b.queryParams) > 0 {
		req.URL.RawQuery = b.queryParams.Encode()
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}
`))
	var buf bytes.Buffer
	err := builderTemplate.Execute(&buf, r)
	if err != nil {
		log.Panicf("Failed to generate template: %v", err)
		return
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Panicf("Failed to generate template: %v", err)
		return
	}

	os.Stdout.Write(formatted)
}

func getParamKey(f *ast.Field) string {
	re := regexp.MustCompile(pattern)
	comment := f.Doc.Text()
	key := extractQueryParam(re, comment)
	if key == "" {
		log.Panicf("Must have query parameter defined.")
	}
	return key
}

func getParamName(e ast.Expr) string {
	p := e.(*ast.FuncType).Params
	if len(p.List) == 0 {
		log.Panicf("Cannot have a function declared without any parameters")
		return ""
	}
	return fmt.Sprintf("string(%s)", p.List[0].Names[0].Name)
}

func getParamsList(e ast.Expr) string {
	p := e.(*ast.FuncType).Params
	var s string
	for i := 0; i < len(p.List); i++ {
		f := p.List[i]
		s += fmt.Sprintf("%s %s", f.Names[0].Name, f.Type.(*ast.Ident).Name)
		if i > 0 {
			s += ","
		}
	}
	return s
}

const (
	query            string = "QUERY"
	body             string = "BODY"
	httpMethodGet    string = "GET"
	httpMethodPost   string = "POST"
	httpMethodPut    string = "PUT"
	httpMethodDelete string = "DELETE"
	httpMethodHead   string = "HEAD"

	pattern string = `@(\w+)\(\"(.*)\"\)`
)

type requestInfo struct {
	Pkg         string
	Name        string
	Api         string
	HttpMethod  string
	QueryParams map[string]*ast.Field
}

type fileASTVisitor struct {
	info     *types.Info
	reqInfo  *requestInfo
	re       *regexp.Regexp
	generate bool
}

func newFileASTVisitor(info *types.Info) *fileASTVisitor {
	return &fileASTVisitor{
		info: info,
		reqInfo: &requestInfo{
			QueryParams: make(map[string]*ast.Field),
		},
		re:       regexp.MustCompile(pattern),
		generate: false,
	}
}

func (v *fileASTVisitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return v
	}

	switch node.(type) {
	case *ast.File:
		v.generate = false
		break
	case *ast.TypeSpec:
		// Check if we are at the beginning of a request builder declaration
		// This must be an interface
		typeSpec := node.(*ast.TypeSpec)
		switch typeSpec.Type.(type) {
		case *ast.InterfaceType:
			v.reqInfo.Name = typeSpec.Name.Name
			break
		}
		break
	case *ast.InterfaceType:
		// Retain a mapping of interface methods to their fields which contain
		// the query parameter and argument name and type information to implement
		// the interface
		ifc := node.(*ast.InterfaceType)
		methods := ifc.Methods
		for _, f := range methods.List {
			if qp := extractQueryParam(v.re, f.Doc.List[0].Text); qp != "" {
				v.reqInfo.QueryParams[f.Names[0].Name] = f
			}
		}
		break
	case *ast.Comment:
		// Extract the HTTP Method and API from the Interface declaration
		comment := node.(*ast.Comment)
		if httpMethod, api := extractHttpMethod(v.re, comment.Text); httpMethod != "" {
			v.generate = true
			v.reqInfo.HttpMethod = httpMethod
			v.reqInfo.Api = api
		}
		break
	}

	return v
}

func extractHttpMethod(re *regexp.Regexp, s string) (string, string) {
	match := re.FindStringSubmatch(s)
	if len(match) == 3 && isAnnotationHttpMethod(match[1]) {
		return match[1], match[2]
	}
	return "", ""
}

func isAnnotationHttpMethod(s string) bool {
	return s == httpMethodGet || s == httpMethodPost || s == httpMethodPut || s == httpMethodHead || s == httpMethodDelete
}

func extractQueryParam(re *regexp.Regexp, s string) string {
	match := re.FindStringSubmatch(s)
	if len(match) == 3 && match[1] == query {
		return match[2]
	}
	return ""
}
