package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/Hiroya3/learning-graphql/graph/model"
)

// PostPhoto is the resolver for the postPhoto field.
func (r *mutationResolver) PostPhoto(ctx context.Context, input model.PostPhotoInput) (*model.Photo, error) {
	panic(fmt.Errorf("not implemented: PostPhoto - postPhoto"))
}

// TagPhoto is the resolver for the tagPhoto field.
func (r *mutationResolver) TagPhoto(ctx context.Context, githubLogin string, photoID string) (*model.Photo, error) {
	panic(fmt.Errorf("not implemented: TagPhoto - tagPhoto"))
}

// GithubAuth is the resolver for the githubAuth field.
func (r *mutationResolver) GithubAuth(ctx context.Context, code string) (*model.AuthPayload, error) {
	panic(fmt.Errorf("not implemented: GithubAuth - githubAuth"))
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
	return 42, nil
}

// AllPhotos is the resolver for the allPhotos field.
func (r *queryResolver) AllPhotos(ctx context.Context) ([]*model.Photo, error) {
	panic(fmt.Errorf("not implemented: AllPhotos - allPhotos"))
}

// Photo is the resolver for the Photo field.
func (r *queryResolver) Photo(ctx context.Context, id string) (*model.Photo, error) {
	panic(fmt.Errorf("not implemented: Photo - Photo"))
}

// TotalUsers is the resolver for the totalUsers field.
func (r *queryResolver) TotalUsers(ctx context.Context) (int, error) {
	panic(fmt.Errorf("not implemented: TotalUsers - totalUsers"))
}

// AllUsers is the resolver for the allUsers field.
func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: AllUsers - allUsers"))
}

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context, login string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - User"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
