package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/Hiroya3/learning-graphql/app/service/auth"
	"github.com/Hiroya3/learning-graphql/db"
	"github.com/Hiroya3/learning-graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Hiroya3/learning-graphql/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// mongoDBのclient作成
	fmt.Println(os.Getenv("MONGO_URI"))
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if deferErr := client.Disconnect(ctx); deferErr != nil {
			log.Fatal(deferErr)
		}
	}()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success!!")
	}

	// dbとcollectionの作成
	_ = client.Database(db.DbName).RunCommand(ctx, bson.D{{"create", db.PhotoCollection}})
	_ = client.Database(db.DbName).RunCommand(ctx, bson.D{{"create", db.UserCollection}})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			AuthService: auth.NewAuthService(),
			PhotoChs:    map[string]chan *model.Photo{},
			Mutex:       sync.Mutex{},
			DbClient:    client,
		},
	}))

	// websocketの登録
	srv.AddTransport(&transport.Websocket{})

	// 複雑度の許容設定
	// ref:https://github.com/99designs/gqlgen/blob/master/docs/content/reference/complexity.md
	srv.Use(extension.FixedComplexityLimit(5))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", http.TimeoutHandler(srv, 5000*time.Microsecond, ""))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
