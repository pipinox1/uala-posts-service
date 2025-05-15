package content

import "errors"

var ErrTextContentEmpty = errors.New("content.text.url_empty")
var ErrTextContentTooLong = errors.New("content.text.too_long")

type TextContentValidator struct{}

func (v *TextContentValidator) Validate(content Content) error {
	if content.Text == nil || *content.Text == "" {
		return ErrTextContentEmpty
	}

	if len(*content.Text) > 1000 {
		return ErrTextContentTooLong
	}

	return nil
}
