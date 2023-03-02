package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Hiroya3/learning-graphql/db"
	"github.com/Hiroya3/learning-graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	assetsDir = "../assets/photo"
)

// PostPhoto is the resolver for the postPhoto field.
func (r *mutationResolver) PostPhoto(ctx context.Context, input model.PostPhotoInput) (*model.Photo, error) {
	now := time.Now()

	// photo情報をmongoに保存
	doc, err := r.DbClient.Database(dbName).Collection(photoCollection).InsertOne(ctx, &db.Photo{
		Name:        input.Name,
		Description: input.Description,
		Category:    string(*input.Category),
		CreatedAt:   now,
	})
	if err != nil {
		log.Printf("fail to postPhoto,%v", err)
		return nil, err
	}

	id := doc.InsertedID.(primitive.ObjectID).Hex()

	// fileをローカルに保存
	// ディレクトリの存在確認
	err = createFileAssetsDirIfNeed()
	if err != nil {
		log.Printf("fail to create assets dir,%v", err)
		return nil, err
	}

	filePath := fmt.Sprintf("%v/%v_%v", assetsDir, id, input.File.Filename)
	fileBytes := make([]byte, 0)
	for {
		i, err2 := input.File.File.Read(fileBytes)
		if err2 != nil {
			log.Printf("fail to read file,%v", err)
			return nil, err2
		}
		if i == 0 {
			break
		}
	}
	err = os.WriteFile(filePath, fileBytes, 0777)
	if err != nil {
		log.Printf("fail to write file,%v", err)
		return nil, err
	}

	result := &model.Photo{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
		Category:    *input.Category,
		Created:     now.String(),
	}

	// subしているuserにpubする
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	for _, ch := range r.PhotoChs {
		ch <- result
	}

	return result, nil
}

func createFileAssetsDirIfNeed() error {
	if _, err := os.Stat(assetsDir); os.IsNotExist(err) {
		// ./tempがなければ作成する
		err = os.Mkdir(assetsDir, 0777)
		if err != nil {
			return err
		}
	}

	return nil
}

// TagPhoto is the resolver for the tagPhoto field.
func (r *mutationResolver) TagPhoto(ctx context.Context, githubLogin string, photoID string) (*model.Photo, error) {
	panic(fmt.Errorf("not implemented: TagPhoto - tagPhoto"))
}

// GithubAuth is the resolver for the githubAuth field.
func (r *mutationResolver) GithubAuth(ctx context.Context, code string) (*model.AuthPayload, error) {
	return r.AuthService.GetAuth(ctx, code)
}

// AddFakeUsers is the resolver for the addFakeUsers field.
func (r *mutationResolver) AddFakeUsers(ctx context.Context, count *int) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: AddFakeUsers - addFakeUsers"))
}

// FakeUserAuth is the resolver for the fakeUserAuth field.
func (r *mutationResolver) FakeUserAuth(ctx context.Context, githubLogin string) (*model.AuthPayload, error) {
	panic(fmt.Errorf("not implemented: FakeUserAuth - fakeUserAuth"))
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Me - me"))
}

// TotalPhotos is the resolver for the totalPhotos field.
func (r *queryResolver) TotalPhotos(ctx context.Context) (int, error) {
	counts, err := r.DbClient.Database(dbName).Collection(photoCollection).CountDocuments(ctx, &bson.D{})
	if err != nil {
		log.Printf("fail to count totalPhotos,%v", err)
		return 0, err
	}
	return int(counts), nil
}

// AllPhotos is the resolver for the allPhotos field.
func (r *queryResolver) AllPhotos(ctx context.Context) ([]*model.Photo, error) {
	res, err := r.DbClient.Database(dbName).Collection(photoCollection).Find(ctx, &bson.D{})
	if err != nil {
		log.Printf("fail to find allPhotos,%v", err)
		return nil, err
	}

	result := make([]*model.Photo, 0)

	for res.Next(ctx) {
		var v db.Photo
		err = res.Decode(&v)
		if err != nil {
			log.Printf("fail to decode,%v", err)
			return nil, err
		}

		result = append(result, &model.Photo{
			ID:          v.Id,
			Name:        v.Name,
			URL:         v.URL,
			Description: v.Description,
			Category:    model.PhotoCategory(v.Category),
			PostedBy:    &model.User{}, // 現時点はnil
			TaggedUsers: nil,           // 現時点はnil
			Created:     v.CreatedAt.String(),
		})
	}

	err = res.All(ctx, &result)
	if err != nil {
		log.Printf("fail to bind allPhotos,%v", err)
		return nil, err
	}

	return result, nil
}

// Photo is the resolver for the Photo field.
func (r *queryResolver) Photo(ctx context.Context, id string) (*model.Photo, error) {
	panic(fmt.Errorf("not implemented: Photo - Photo"))
}

// TotalUsers is the resolver for the totalUsers field.
func (r *queryResolver) TotalUsers(ctx context.Context) (int, error) {
	count, err := r.DbClient.Database(dbName).Collection(userCollection).CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Printf("fail to count totalUsers,%v", err)
		return 0, err
	}

	return int(count), nil
}

// AllUsers is the resolver for the allUsers field.
func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	res, err := r.DbClient.Database(dbName).Collection(userCollection).Find(ctx, &bson.D{})
	if err != nil {
		log.Printf("fail to find allUsers,%v", err)
		return nil, err
	}

	result := make([]*model.User, 0)

	for res.Next(ctx) {
		var v db.User
		err = res.Decode(&v)
		if err != nil {
			log.Printf("fail to decode,%v", err)
			return nil, err
		}

		result = append(result, &model.User{
			GithubLogin:  v.GithubLogin,
			Name:         v.Name,
			Avatar:       v.Avatar,
			PostedPhotos: nil, // 現時点ではnil
			InPhotos:     nil, // 現時点ではnil
		})
	}

	err = res.All(ctx, &result)
	if err != nil {
		log.Printf("fail to bind allUsers,%v", err)
		return nil, err
	}

	return result, nil
}

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context, login string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - User"))
}

// NewPhoto is the resolver for the newPhoto field.
func (r *subscriptionResolver) NewPhoto(ctx context.Context, userID string) (<-chan *model.Photo, error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	// ユーザーからsub依頼が来たので、チャネルを追加してpub対象に含める

	// 登録済みであればエラー
	if _, ok := r.PhotoChs[userID]; ok {
		log.Printf(fmt.Errorf("%v is already subscribed", userID).Error())
		return nil, fmt.Errorf("%v is already subscribed", userID)
	}

	ch := make(chan *model.Photo)
	r.PhotoChs[userID] = ch

	go func() {
		// コネクションが終了（subscribeが終了したらチャネルを削除する）
		<-ctx.Done()
		r.Mutex.Lock()
		delete(r.PhotoChs, userID)
		r.Mutex.Unlock()
		log.Printf("%s has been unsubscribed", userID)
	}()

	log.Printf("%v is subscribed", userID)

	return ch, nil
}

// NewUser is the resolver for the newUser field.
func (r *subscriptionResolver) NewUser(ctx context.Context) (<-chan *model.User, error) {
	panic(fmt.Errorf("not implemented: NewUser - newUser"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
