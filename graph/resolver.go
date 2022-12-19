package graph

import (
	"github.com/Hiroya3/learning-graphql/app/service/auth"
	"github.com/Hiroya3/learning-graphql/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PhotoId     int
	Photos      []*model.Photo
	AuthService auth.AuthService
	DbClient    *mongo.Client
}
