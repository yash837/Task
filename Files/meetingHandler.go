package main

import "net/http"

func MeetingHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		CreateMeeting(response, request)
	}
	if request.Method == "GET" {
		GetMeetingwithTime(response, request)
	}
}
