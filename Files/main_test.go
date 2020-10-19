package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMainCreatemeet(t *testing.T) {
	var message Meeting
	var part participant
	part.Name = "Yash Jain"
	part.Email = "yashjain289.mail@gmail.com"
	part.Rsvp = "Yes"
	message.Title = "Title"
	message.Participants = append(message.Participants, part)
	message.Starttime = "2021-09-01T09:52:12+05:30"
	message.Endtime = "2021-09-01T10:52:12+05:30"
	bytesRepresentation, _ := json.Marshal(message)
	resp, err := http.Post("http://localhost:4000/meetings", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		t.Error("Fail")
	}
	if resp == nil {
		t.Error("No response")
	}
}
func TestMaingetmeet(t *testing.T) {
	resp, err := http.Get("http://localhost:4000/meeting/?id=5f4dcb738b246dc74d8ecd44")
	if err != nil {
		t.Error("Fail")
	}
	if resp == nil {
		t.Error("No response")
	}
}
func TestMaingetparticipants(t *testing.T) {
	resp, err := http.Get("http://localhost:4000/articles/?participant=manast95@gmail.com")
	if err != nil {
		t.Error("Fail")
	}
	if resp == nil {
		t.Error("No response")
	}
}
func TestMaingetmeetintime(t *testing.T) {
	resp, err := http.Get("http://localhost:4000/meetings?start=2019-09-01T13:30:10+05:30&end=2021-09-01T14:30:10+05:30")
	if err != nil {
		t.Error("Fail")
	}
	if resp == nil {
		t.Error("No response")
	}
}

func BenchmarkMaingetmeet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		http.Get("http://localhost:4000/meeting/5f4dcc74fa1a4b2011daf69a")
	}
}

func BenchmarkMaingetparticipant(b *testing.B) {
	for n := 0; n < b.N; n++ {
		http.Get("http://localhost:4000/articles/?participant=arnavdixit127@gmail.com")
	}
}

func BenchmarkMaingettime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		http.Get("http://localhost:4000/meetings?start=2019-09-01T13:30:10+05:30&end=2021-09-01T14:30:10+05:30")
	}
}

func BenchmarkMaingetpost(b *testing.B) {
	var message Meeting
	var part participant
	part.Name = "Yash Jain"
	part.Email = "yashjain289.mail@gmail.com"
	part.Rsvp = "No"
	message.Title = "Title"
	message.Participants = append(message.Participants, part)
	message.Starttime = "2021-09-01T09:52:12+05:30"
	message.Endtime = "2021-09-01T10:52:12+05:30"

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error")
	}

	for n := 0; n < b.N; n++ {
		resp, err := http.Post("http://localhost:4000/meetings", "application/json", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			b.Error("Fail")
		}
		if resp == nil {
			b.Error("NO response")
		}
	}
}

func BenchmarkParticipantsBusy(b *testing.B) {
	var message Meeting
	var part participant
	part.Name = "Yash Jain"
	part.Email = "yashjain289.mail@gmail.com"
	part.Rsvp = "No"
	message.Title = "Title"
	message.Participants = append(message.Participants, part)
	message.Starttime = "2021-09-01T09:52:12+05:30"
	message.Endtime = "2021-09-01T10:52:12+05:30"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb+srv://yashjain:yashjain@cluster0.dnamu.mongodb.net/go_rest_api?w=majority")
	client, _ = mongo.Connect(ctx, clientOptions)
	for n := 0; n < b.N; n++ {
		ParticipantsBusy(message)
	}
}

func BenchmarkCheckParticipant(b *testing.B) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb+srv://yashjain:yashjain@cluster0.dnamu.mongodb.net/go_rest_api?w=majority")
	client, _ = mongo.Connect(ctx, clientOptions)
	for n := 0; n < b.N; n++ {
		_ = CheckParticipant("yashjain289.mail@gmail.com")
	}
}
