package model

type Image interface {
	URL() string
	Size() Size
}
