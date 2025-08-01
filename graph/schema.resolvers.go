package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.76

import (
	"context"
	"fmt"

	"github.com/mdzahid786/golang-graphql/db"
	"github.com/mdzahid786/golang-graphql/graph/model"
	"github.com/mdzahid786/golang-graphql/helper"
	"golang.org/x/crypto/bcrypt"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	bcrypt, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user := &model.User{Name: input.Name, Email: input.Email, Password: string(bcrypt)}
	if err := db.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := db.DB.Preload("Todos").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Login operation on the user
func (r *queryResolver) Login(ctx context.Context, email string, password string) (*model.User, error) {
	var user *model.User
	err := db.DB.Where("email=?", email).First(&user).Error
	if err != nil {
		return user, fmt.Errorf("invalid email or user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("password does not match")
	}
	token, err := helper.GenerateJWT(user.ID)
	if err != nil {
		return &model.User{}, err
	}
	user.Token = token
	return user, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	userID, err := helper.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("access denied: %w", err)
	}

	var user model.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
