package db

import "time"

// mongoのcollection定義

type Photo struct {
	Name           string    `bson:"name"`
	URL            string    `bson:"url"`
	Description    *string   `bson:"description,omitempty"`
	Category       *string   `bson:"category,omitempty"`
	PostedByUserId *string   `bson:"posted_by_user_id,omitempty"`
	TaggedUsersId  []*string `bson:"tagged_users_id,omitempty"`
	CreatedAt      time.Time `bson:"createdAt"`
}
