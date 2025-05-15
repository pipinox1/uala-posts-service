package content

import "errors"

var ErrImageContentEmptyUrl = errors.New("content.image.url_empty")

type ImageContentValidator struct{}

func (v *ImageContentValidator) Validate(content Content) error {
	if content.Url == nil {
		return ErrImageContentEmptyUrl
	}
	return nil
}
