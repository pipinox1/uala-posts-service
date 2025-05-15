package content

import (
	"fmt"
)

type ContentValidator interface {
	Validate(content Content) error
}

type ContentFactory struct {
	validators map[ContentType]ContentValidator
}

func NewContentFactory() *ContentFactory {
	factory := &ContentFactory{
		validators: map[ContentType]ContentValidator{
			TextContentType:  &TextContentValidator{},
			ImageContentType: &ImageContentValidator{},
		},
	}
	return factory
}

type ContentBody struct {
	Text *string
	Url  *string
}

func (f *ContentFactory) CreateContent(contentTypeStr string, bodyContent ContentBody) (*Content, error) {
	contentType, err := NewContentType(contentTypeStr)
	if err != nil {
		return nil, err
	}

	content := Content{
		Type: contentType,
		Text: bodyContent.Text,
	}

	validator, exists := f.validators[contentType]
	if !exists {
		return nil, fmt.Errorf("no validator found for content type: %s", contentType)
	}

	if err := validator.Validate(content); err != nil {
		return nil, err
	}

	return &content, nil
}
