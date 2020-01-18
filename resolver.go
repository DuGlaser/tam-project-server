package tam_project_server

import (
	"context"

	detabase "github.com/DuGlaser/tam-project-server/database"
	"github.com/DuGlaser/tam-project-server/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) PostMessage(ctx context.Context, text string, chatroomID string) (*models.Message, error) {
	db := detabase.FetchConnection()
	defer db.Close()

	message := &models.Message{
		Text:       text,
		ChatroomID: chatroomID,
	}

	db.Create(message)

	return message, nil
}
func (r *mutationResolver) CreateRoom(ctx context.Context, name string) (*models.Chatroom, error) {
	db := detabase.FetchConnection()
	defer db.Close()

	chatroom := &models.Chatroom{
		Name: name,
	}

	db.Create(chatroom)
	return chatroom, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Room(ctx context.Context, name string) (*models.Chatroom, error) {
	db := detabase.FetchConnection()
	defer db.Close()

	var chatroom models.Chatroom
	db.Preload("Message").First(&chatroom, name)

	return &chatroom, nil
}

func (r *queryResolver) Rooms(ctx context.Context) ([]*models.Chatroom, error) {
	db := detabase.FetchConnection()
	defer db.Close()

	var chatroom []*models.Chatroom
	db.Preload("Message").First(&chatroom)

	return chatroom, nil
}
