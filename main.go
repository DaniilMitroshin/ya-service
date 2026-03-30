package main

import (
	"Vservice/internal/repository"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	fmt.Println("started")

	ConnString := os.Getenv("CONN_STRING")
	fmt.Println(ConnString)
	//"postgres://postgres:1234@localhost:5432/testdb"

	ctxPool, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := pgxpool.New(ctxPool, ConnString)
	if err != nil {
		panic("not connected")
	}

	repo := repository.NewRepo(pool)

	//hhtp recieved
	/* create and fill
	ctxCreate, cancel := context.WithTimeout(context.Background(), time.Second*5)
	repo.CreateTable(ctxCreate)

	//http recieved
	ctxInsert, cancel := context.WithTimeout(context.Background(), time.Second*5)
	repo.InsertData(ctxInsert)
	*/

	//http recieved
	/*update
	bookUpdateParams := dbmodel.BookUpdateParams{
		Id:       1,
		Title:    shared.Some("New title"),
		Author:   shared.None[string](),
		NumPages: shared.Null[int](),
		Rating:   shared.Some[float64](4.98),
	}
	ctxUpdate, cancel := context.WithTimeout(context.Background(), time.Second*5)
	repo.UpdateData_tx(ctxUpdate, bookUpdateParams)
	*/
	/*
		ctxGet, cancel := context.WithTimeout(context.Background(), time.Second*5)
		res, err := repo.SelectById(ctxGet, 1)
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Print(res)
		}*/

	ctxDelete, cancel := context.WithTimeout(context.Background(), time.Second*5)
	err = repo.DropTable(ctxDelete)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("ended")
}
