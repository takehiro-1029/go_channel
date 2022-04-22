package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Server struct {
	chatService *ChatService
	channel     chan uuid.UUID
}

type CreateMessage struct {
	Body string `json:"body"`
}

func main() {
	r := chi.NewRouter()
	chatService := ChatService{
		mutex:    sync.RWMutex{},
		messages: map[uuid.UUID]Message{},
	}

	server := Server{
		chatService: &chatService,
		channel:     make(chan uuid.UUID, 5),
	}

	go func() {
		for v := range server.channel {
			messages, _ := server.chatService.FindMessages(&v)
			log.Println(messages)
		}
	}()

	r.Post("/messages", server.CreateMessageHandler)

	http.ListenAndServe(":8080", r)
}

func (server *Server) CreateMessageHandler(w http.ResponseWriter, req *http.Request) {
	var input CreateMessage

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 10000))
	if err != nil {
		httpError(w, err)
		return

	}
	if err := req.Body.Close(); err != nil {
		httpError(w, err)
		return
	}
	if err := json.Unmarshal(body, &input); err != nil {
		httpError(w, err)
		return
	}

	message, err := server.chatService.CreateMessage(input.Body)
	if err != nil {
		httpError(w, err)
		return
	}
	server.channel <- message.ID

	httpOk(w, message)
}

func httpError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err)
}

func httpOk(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
