package main

import (
	stdlog "log"
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	log "github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

type Record struct {
	ID      int `json: "id"`
	message int `json: "message"`
}

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	r := mux.NewRouter()
	db := GetDBconn()
	var svc RecordService
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	repository, err := NewRepo(db, logger)
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}
	svc = NewService(repository, logger)

	addMessageHandler := httptransport.NewServer(
		makeCreateMessageEndpoint(svc),
		decodeAddMessageRequest,
		encodeResponse,
	)

	readMessageHandler := httptransport.NewServer(
		makeCreateMessageEndpoint(svc),
		decodeReadMessageRequest,
		encodeResponse,
	)

	updateMessageHandler := httptransport.NewServer(
		makeCreateMessageEndpoint(svc),
		decodeUpdateMessageRequest,
		encodeResponse,
	)

	deleteMessageHandler := httptransport.NewServer(
		makeCreateMessageEndpoint(svc),
		decodeDeleteMessageRequest,
		encodeResponse,
	)

	r.Handle("/message", addMessageHandler).Methods("POST")
	r.Handle("/message", readMessageHandler).Methods("GET")
	r.Handle("/message", updateMessageHandler).Methods("PUT")
	r.Handle("/message", deleteMessageHandler).Methods("DELETE")
	stdlog.Fatal(http.ListenAndServe(":8080", r))
}
