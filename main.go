package main

import (
	"Vservice/internal/repository"
	"Vservice/internal/repository/dbmodel"
	"Vservice/internal/shared"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	fmt.Println("started")

	ConnString := "postgres://postgres:daniil@localhost:5432/postgres"
	//"postgres://postgres:1234@localhost:5432/testdb"

	ctxPool, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := pgxpool.New(ctxPool, ConnString)
	if err != nil {
		panic("not connected")
	}

	repo := repository.NewRepo(pool)

	//hhtp recieved
	/*
		ctxCreate, cancel := context.WithTimeout(context.Background(), time.Second*5)
		repo.CreateTable(ctxCreate)

		//http recieved
		ctxInsert, cancel := context.WithTimeout(context.Background(), time.Second*5)
		repo.InsertData(ctxInsert)
	*/

	//http recieved

	bookUpdateParams := dbmodel.BookUpdateParams{
		Id:       1,
		Title:    shared.Some("New title"),
		Author:   shared.None[string](),
		NumPages: shared.Null[int](),
		Rating:   shared.Some[float64](4.98),
	}
	ctxUpdate, cancel := context.WithTimeout(context.Background(), time.Second*5)
	repo.UpdateData_tx(ctxUpdate, bookUpdateParams)

	fmt.Println("ended")
}
