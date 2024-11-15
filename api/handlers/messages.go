package handlers

import (
	"api/models"
	"api/postgres"
	"encoding/json"
	"io"
	"net/http"
)

func AddMessagesHandlers(mux *http.ServeMux) {
	mux.HandleFunc("GET /chat", func(writer http.ResponseWriter, request *http.Request) {
		messages, err := postgres.GetChatHistory()
		if err != nil {
			io.WriteString(writer, err.Error())
			return
		}

		marshalled, err := json.Marshal(models.History{
			Messages: messages,
		})
		if err != nil {
			io.WriteString(writer, err.Error())
			return
		}

		writer.Write(marshalled)
	})

	mux.HandleFunc("POST /chat", func(writer http.ResponseWriter, request *http.Request) {
		var message models.Message
		err := json.NewDecoder(request.Body).Decode(&message)
		if err != nil {
			io.WriteString(writer, err.Error())
			return
		}

		err = postgres.UploadText(message)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			io.WriteString(writer, err.Error())
			return
		}

		writer.WriteHeader(http.StatusCreated)
	})
}
