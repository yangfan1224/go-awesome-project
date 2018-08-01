package main

import (
		"net/http"
	"github.com/gorilla/mux"
	"log"
	"github.com/yangfan1224/go-awesome-project/muxhandler"
	)

func main(){
	//channel.StartServer()
	topicHanlder := muxhandler.NewTopicHandler()
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/topic", topicHanlder.HandleAdd).Methods("POST")
	r.HandleFunc("/topic/{id}", topicHanlder.HandleGet).Methods("GET")
	r.HandleFunc("/topic", topicHanlder.HandleGetAll).Methods("GET")
	r.HandleFunc("/topic", topicHanlder.HandleModify).Methods("PUT")
	r.HandleFunc("/topic/{id}", topicHanlder.HandleDelete).Methods("DELETE")

	// Bind to a port and pass our router in

	log.Fatal(http.ListenAndServe(":8000", r))
}
