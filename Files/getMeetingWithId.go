package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckMeetingwithID(id primitive.ObjectID) (Meeting, error) {
	var meet Meeting
	collection := client.Database("go_rest_api").Collection("Meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, Meeting{ID: id}).Decode(&meet)
	if meet.ID != id {
		err = errors.New("Error 400 found: ID not present")
	}
	return meet, err
}

func GetMeetingwithID(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		response.Header().Set("content-type", "application/json")
		fmt.Println(path.Base(request.URL.Path))
		id, _ := primitive.ObjectIDFromHex(path.Base(request.URL.Path))
		meetingwithID, err := CheckMeetingwithID(id)
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		json.NewEncoder(response).Encode(meetingwithID)
	}
}
