package entities

import "github.com/google/uuid"

type Tweet struct {
	ID string `json:id`
	Name string `json:name`
	Age uint8 `json:age`
	Profession string `json:profession`
	Description string `json:description`
}

func NewTweet() *Tweet {
	tweet := Tweet {
		ID: uuid.NewString(),
	}

	return &tweet
}