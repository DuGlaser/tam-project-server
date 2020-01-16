package tam_project_server

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Chatroom() ChatroomResolver {
	return &chatroomResolver{r}
}
func (r *Resolver) Message() MessageResolver {
	return &messageResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type chatroomResolver struct{ *Resolver }

func (r *chatroomResolver) Message(ctx context.Context, obj *Chatroom) ([]*Message, error) {
	panic("not implemented")
}

type messageResolver struct{ *Resolver }

func (r *messageResolver) Text(ctx context.Context, obj *Message) (string, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) PostMessage(ctx context.Context, text string, roomName string) (*Message, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Room(ctx context.Context, name string) (*Chatroom, error) {
	panic("not implemented")
}
