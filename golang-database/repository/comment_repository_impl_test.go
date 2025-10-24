package repository

import (
	"context"
	"fmt"
	golangdatabase "golang-database"
	"golang-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "respository@email.com",
		Comment: "Respository Comment",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Comment: ", result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Get Comment: ", comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
