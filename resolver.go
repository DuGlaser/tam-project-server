package tam_project_server

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) PostMessage(ctx context.Context, text string, roomName string) (*Message, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Room(ctx context.Context, name string) (*Chatroom, error) {
	panic("not implemented")
}
func (r *queryResolver) Rooms(ctx context.Context) (*Chatroom, error) {
	panic("not implemented")
}
