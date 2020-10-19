package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Application Running")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb+srv://yashjain:yashjain@cluster0.dnamu.mongodb.net/go_rest_api?w=majority")
	client, _ = mongo.Connect(ctx, clientOptions)

	http.HandleFunc("/meetings", MeetingHandler)
	http.HandleFunc("/articles/", GetParticipants)
	http.HandleFunc("/meeting/", GetMeetingwithID)
	http.ListenAndServe(":4000", nil)
}
