package go500px

//go:generate $GOPATH/src/github.com/jsaund/gorest/gorest -input post_comment_request_builder.go -output post_comment_request_builder_impl.go -pkg go500px

// @POST_FORM("/photos/{id}/comments")
type PostCommentRequestBuilder interface {
	// @PATH("id")
	PhotoID(id string) PostCommentRequestBuilder

	// @FIELD("body")
	Body(body string) PostCommentRequestBuilder

	// @SYNC("PostCommentResponse")
	Run() (PostCommentResponse, error)

	// @ASYNC("PostCommentCallback")
	RunAsync(callback PostCommentCallback)
}
