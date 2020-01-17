package tam_project_server

import (
	"context"

	"github.com/DuGlaser/tam-project-server/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

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

func (r *chatroomResolver) Message(ctx context.Context, obj *models.Chatroom) ([]*models.Message, error) {
	return []*models.Message{}, nil
}

type messageResolver struct{ *Resolver }

func (r *messageResolver) Text(ctx context.Context, obj *models.Message) (string, error) {
	return obj.Text, nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) PostMessage(ctx context.Context, text string, roomName string) (*models.Message, error) {

	message := &models.Message{
		Text: text,
		ID:   1, /*TODO:randomなintの数*/
	}
	_ = models.Chatroom{
		Name:    roomName,
		Message: *message,
	}

	return message, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Room(ctx context.Context, name string) (*models.Chatroom, error) {
	return &models.Chatroom{}, nil
}
