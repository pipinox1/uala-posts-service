package content

import (
	"errors"
	"strings"
)

var ErrInvalidContentType = errors.New("content_type.invalid")

type ContentType string

const (
	TextContentType  ContentType = "text"
	ImageContentType ContentType = "image"
)

var allowedContentTypes = map[string]ContentType{
	TextContentType.String():  TextContentType,
	ImageContentType.String(): ImageContentType,
}

func NewContentType(contentType string) (ContentType, error) {
	if contentType, ok := allowedContentTypes[strings.ToLower(contentType)]; ok {
		return contentType, nil
	}

	return "", ErrInvalidContentType
}

func (c ContentType) String() string {
	return string(c)
}
