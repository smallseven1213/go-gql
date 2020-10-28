package main

import (
	"gql/graph"
	"gql/graph/generated"
	mysqldbmodel "gql/mysqldb/model"
	"gql/service"
	"gql/utils"
	"net/http"

	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
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
	db.AutoMigrate(&mysqldbmodel.Todo{})
}

const defaultPort = "9090"

func main() {
	db := utils.SQLDBInit()
	Migrate(db)

	router := chi.NewRouter()
	router.Use(service.Middleware())

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	err := http.ListenAndServe(":9090", router)
	if err != nil {
		panic(err)
	}
}
