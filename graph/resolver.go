package graph

import (
	"github.com/Hiroya3/learning-graphql/app/service/auth"
	"github.com/Hiroya3/learning-graphql/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthService auth.AuthService
	PhotoChs    map[string]chan *model.Photo // photoの追加時にpubするchannel
	Mutex       sync.Mutex                   // channelの追加・削除の排他制御
	DbClient    *mongo.Client
}

const (
	dbName          = "sample"
	photoCollection = "photo"
	userCollection  = "user"
)
