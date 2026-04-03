package main

import (
	"Vservice/internal/db"
	"Vservice/internal/repository"
	"Vservice/internal/shared"
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
	/*
		ctxInsert, cancel := context.WithTimeout(context.Background(), time.Second*5)
		_, err = repo.InsertBook(ctxInsert, db.InsertBookParams{
			Title:    shared.Ptr("newBook1"),
			Author:   shared.Ptr("NewAuthor1"),
			NumPages: nil,
			Rating:   shared.Ptr(4.25),
		})
		if err != nil {
			fmt.Print(err)
		}
	*/
	ctxTransaction, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = repo.WithTx(ctxTransaction, func(r *repository.Repo) error {
		_, err := r.InsertBook(ctxTransaction, db.InsertBookParams{
			Title:    shared.Ptr("newBook2"),
			Author:   shared.Ptr("NewAuthor2"),
			NumPages: nil,
			Rating:   shared.Ptr(2.25),
		})
		if err != nil {
			return err
		}
		_, err = r.InsertBook(ctxTransaction, db.InsertBookParams{
			Title:    shared.Ptr("newBook3"),
			Author:   shared.Ptr("NewAuthor3"),
			NumPages: shared.Ptr(int32(250)),
			Rating:   nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ended")
}
