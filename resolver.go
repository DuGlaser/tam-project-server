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

func (r *mutationResolver) CreateNote(ctx context.Context, title string) (*models.Note, error) {
	db := detabase.FetchConnection()
	defer db.Close()

	note := &models.Note{
		Title: title,
	}

	db.Create(note)

	return note, nil
}
func (r *mutationResolver) UpdateNote(ctx context.Context, id string, content string) (*models.Note, error) {
	db := detabase.FetchConnection()
	defer db.Close()

	note := models.Note{ID: id}
	db.First(&note)

	note.Content = content
	db.Model(&models.Note{}).Update(&note)
	return &note, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Note(ctx context.Context, id string) (*models.Note, error) {
	db := detabase.FetchConnection()
	defer db.Close()

	var note models.Note
	db.Preload("Message").Find(&note, id)

	return &note, nil
}

func (r *queryResolver) Notes(ctx context.Context) ([]*models.Note, error) {
	db := detabase.FetchConnection()
	defer db.Close()

	var note []*models.Note
	db.Preload("Message").Find(&note)

	return note, nil
}
