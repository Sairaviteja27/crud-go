package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-kit/log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "test@123"
	dbname   = "postgres"
)

var db *sql.DB

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func GetDBconn() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}

// Describes the accountservice for repository interaction
type Repository interface {
	addMessage(context.Context, messageRequest) error
}

// Creates and returns an instance
func NewRepo(db *sql.DB, logger log.Logger) (Repository, error) {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}, nil
}

func (repo *repo) addMessage(ctx context.Context, messageReq messageRequest) error {
	fmt.Printf("db is %#+v", repo.db)
	_, err := repo.db.ExecContext(ctx, "insert into records(message) values($1)", messageReq.Message)
	if err != nil {
		fmt.Println("Error occured inside addMessage in repo")
		return err
	} else {
		fmt.Println("Message Created:", messageReq.Message)
	}
	return nil
}

func (repo *repo) readMessage(ctx context.Context, messageReq messageRequest) error {
	fmt.Printf("db is %#+v", repo.db)
	err := repo.db.QueryRowContext(ctx, "select * from records").Scan(&messageReq.ID,
		&messageReq.Message)
	fmt.Println("err", err)
	if err != nil {

		return nil
	}
	fmt.Println("messageReq", messageReq)
	return nil
}
