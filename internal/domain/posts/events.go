package posts

import "encoding/json"

const (
	PostCreatedTopic        = "post.created"
	PostContentUpdatedTopic = "post.updated_content"
)

type PostCreatedEvent struct {
	ID       string `json:"id"`
	AuthorID string `json:"author_id"`
}

func (p PostCreatedEvent) Key() string {
	return p.ID
}

func (p PostCreatedEvent) Topic() string {
	return PostCreatedTopic
}

func (p PostCreatedEvent) Payload() []byte {
	payload, _ := json.Marshal(p)
	return payload
}

func NewPostCreatedEvent(post *Post) PostCreatedEvent {
	return PostCreatedEvent{ID: post.ID, AuthorID: post.AuthorId}
}

type PostContentUpdatedEvent struct {
	ID string
}

func (p *PostContentUpdatedEvent) Key() string {
	return p.ID
}

func (p *PostContentUpdatedEvent) Topic() string {
	return PostContentUpdatedTopic
}

func (p *PostContentUpdatedEvent) Payload() []byte {
	json, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return json
}

func NewPostContentUpdatedEvent(post *Post) PostContentUpdatedEvent {
	return PostContentUpdatedEvent{ID: post.ID}
}
