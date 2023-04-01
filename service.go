package main

import (
	"context"
	"errors"
	"fmt"

	stdlog "log"

	"github.com/go-kit/kit/log"
)

// RecordService provides operations on records.
type RecordService interface {
	addRecord(context.Context, messageRequest) (string, error)
	//updateRecord(string) (string, error)
	deleteRecord(context.Context, messageRequest) (string, error)
	readRecords(context.Context, messageRequest) (string, error)
}

func NewService(rep Repository, logger log.Logger) RecordService {
	return &recordService{
		repository: rep,
		logger:     logger,
	}
}

// recordService is a concrete implementation of RecordService
type recordService struct {
	repository Repository
	logger     log.Logger
}

func (s recordService) addRecord(ctx context.Context, messageReq messageRequest) (string, error) {
	err := s.repository.addMessage(ctx, messageReq)
	fmt.Println("resp", err)
	if err != nil {
		stdlog.Fatalf("An error occured while executing query: %v", err)
		return "", errors.New("Some Error Occured")
	}
	return "", nil
}
func (s recordService) deleteRecord(ctx context.Context, messageReq messageRequest) (string, error) {

	_, err := db.Exec("INSERT into records VALUES ($1)", messageReq.Message)
	if err != nil {
		stdlog.Fatalf("An error occured while executing query: %v", err)
		return "", errors.New("Some Error Occured")
	}
	return "", nil
}

func (s recordService) readRecords(ctx context.Context, messageReq messageRequest) (string, error) {

	_, err := db.Exec("INSERT into records VALUES ($1)", messageReq.Message)
	if err != nil {
		stdlog.Fatalf("An error occured while executing query: %v", err)
		return "", errors.New("Some Error Occured")
	}
	return "", nil
}
