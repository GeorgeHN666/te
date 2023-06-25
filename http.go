package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

func NewServer(port string) error {

	srv := http.Server{
		Addr:              fmt.Sprintf(":%v", port),
		Handler:           Handler(),
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}

	fmt.Printf("Server started at port %v \n\n", port)

	return srv.ListenAndServe()
}

func Handler() http.Handler {

	mux := chi.NewRouter()

	mux.Get("/start", StartCampaing)
	mux.Get("/stop", StopCampaing)

	return mux
}

func StartCampaing(w http.ResponseWriter, r *http.Request) {

	campaingID := r.URL.Query().Get("id")
	setTimer := r.URL.Query().Get("set")
	timer, _ := strconv.Atoi(setTimer)

	go StartTimer(campaingID, int64(timer))

}

func StopCampaing(w http.ResponseWriter, r *http.Request) {

	campaingID := r.URL.Query().Get("id")

	CancelTimer(campaingID)

}
