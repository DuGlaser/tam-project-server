package main

import (
	"github.com/99designs/gqlgen/handler"
	tam_project_server "github.com/DuGlaser/tam-project-server"
	database "github.com/DuGlaser/tam-project-server/database"
	models "github.com/DuGlaser/tam-project-server/models"
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.GraphQL(tam_project_server.NewExecutableSchema(tam_project_server.Config{Resolvers: &tam_project_server.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	db := database.FetchConnection()
	db.AutoMigrate(&models.Chatroom{}, &models.Message{})
	defer db.Close()

	// Setting up Gin
	r := gin.Default()

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://sample.com"}
	// r.Use(cors.New(config))

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
