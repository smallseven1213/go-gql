package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gql/graph/generated"
	"gql/graph/model"
	mysqldbmodel "gql/mysqldb/model"
	"gql/utils"
	"strconv"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	db := utils.GetSQLDB()
	todo := mysqldbmodel.Todo{Text: input.Text}

	db.Create(&todo)
	fmt.Printf("%+v\n", todo)

	return &model.Todo{
		ID:   strconv.FormatUint(uint64(todo.ID), 10),
		Text: todo.Text,
		Done: todo.Done,
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	db := utils.GetSQLDB()
	db.Find(&todos)

	var result []*model.Todo

	for _, todo := range todos {
		result = append(result, &model.Todo{
			ID:   strconv.FormatUint(uint64(todo.ID), 10),
			Text: todo.Text,
			Done: todo.Done,
		})
	}

	return result, nil
}

func (r *queryResolver) Todo(ctx context.Context, input *string) (*model.Todo, error) {
	db := utils.GetSQLDB()
	todo := mysqldbmodel.Todo{}
	db.First(&todo, input)

	return &model.Todo{
		ID:   strconv.FormatUint(uint64(todo.ID), 10),
		Text: todo.Text,
		Done: todo.Done,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	todos []mysqldbmodel.Todo
)
