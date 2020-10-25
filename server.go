package main

import (
	"gql/graph"
	"gql/graph/generated"
	"gql/mysqldb/models"
	"gql/utils"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func Migrate(db *gorm.DB) {
	/*
		users.AutoMigrate()
		db.AutoMigrate(&articles.ArticleModel{})
		db.AutoMigrate(&articles.TagModel{})
		db.AutoMigrate(&articles.FavoriteModel{})
		db.AutoMigrate(&articles.ArticleUserModel{})
		db.AutoMigrate(&articles.CommentModel{})
	*/
	db.AutoMigrate(&dbmodel.Todo{})
}

const defaultPort = "9090"

func main() {
	db := utils.Init()
	Migrate(db)
	// defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
