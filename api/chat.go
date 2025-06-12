package main

import (
	"net/http"

	"github.com/atrastudhi/genai-chatbot/pkg/handlers"
)

var chatHandler *handlers.ChatHandler

func init() {
	chatHandler = handlers.NewChatHandler()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	chatHandler.HandleChatRequest(w, r)
}
