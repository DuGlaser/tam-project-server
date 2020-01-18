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

func (r *mutationResolver) PostMessage(ctx context.Context, text string, roomName string) (*models.Message, error) {
	db := detabase.FetchConnection()
	defer db.Close()

	message := &models.Message{
		Text: text,
		ID:   1, /*TODO:randomなintの数*/
	}

	db.Create(message)

	return message, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Room(ctx context.Context, name string) (*models.Chatroom, error) {
	var chatroom *models.Chatroom
	//TODO:DBと接続して情報を取得
	return chatroom, nil
}
func (r *queryResolver) Rooms(ctx context.Context) (*models.Chatroom, error) {
	var chatroom *models.Chatroom
	//TODO:DBと接続して情報を取得

	return chatroom, nil
}
