package db

import "time"

// mongoのcollection定義

type Photo struct {
	Id           string    `bson:"_id,omitempty"`
	Name         string    `bson:"name"`
	URL          string    `bson:"url"`
	Description  *string   `bson:"description,omitempty"`
	Category     string    `bson:"category"`
	PostedByUser *User     `bson:"posted_by_user,omitempty"`
	TaggedUsers  []*User   `bson:"tagged_users,omitempty"`
	CreatedAt    time.Time `bson:"createdAt"`
}

type User struct {
	GithubLogin  string   `bson:"github_login"`
	Name         *string  `bson:"name,omitempty"`
	Avatar       *string  `bson:"avatar,omitempty"`
	PostedPhotos []*Photo `bson:"posted_photos,omitempty"`
	InPhotos     []*Photo `bson:"in_photos,omitempty"`
}
