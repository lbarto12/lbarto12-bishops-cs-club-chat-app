package main

import (
	"api/handlers"
	"api/postgres"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file: ", err)
		return
	}

	host, ok := os.LookupEnv("HOST")
	if !ok {
		log.Fatal("HOST environment variable unset")
		return
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("PORT environment variable unset")
		return
	}

	err = postgres.InitDB()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
		return
	}
	log.Println("Connected to database..")

	finalAddress := fmt.Sprintf("%s:%s", host, port)

	mux := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://207.162.100.14:5173", "*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST"},
	})

	handler := c.Handler(mux)

	handlers.AddMessagesHandlers(mux)
	handlers.AddLiveHandler(mux)

	log.Println("Listening on", finalAddress)
	log.Fatal(http.ListenAndServe(finalAddress, handler))
}
