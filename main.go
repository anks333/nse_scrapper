package main

import (
	"fmt"
	"net/http"
	"nse_scrapper/app"
	"nse_scrapper/handlers"
	"nse_scrapper/routes"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	if err := app.LoadConfig("env/dev.yaml"); err != nil {
		panic(err)
	}

	now := time.Now()

	router := mux.NewRouter()

	router.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, fmt.Sprintf(`
		Date: %v
		Version: %v
		`, now, app.Version))
	})

	router.HandleFunc("/nse", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	registerRoutes(router)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", app.Config.ServerPort),
		Handler: router,
	}

	fmt.Println("Listening on Port ", app.Config.ServerPort)
	panic(server.ListenAndServe())

}

func registerRoutes(r *mux.Router) {
	nseHandler := handlers.NewNSEHandlers()
	routes.NewNSERoutes(r, nseHandler).RegisterRoutes()
}
