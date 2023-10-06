package route

import (
	"gocrudapi/configs"
	porthandler "gocrudapi/handler/portHandler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run(bookHandler porthandler.BookInterface) {
	conf := configs.GetConfig()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", bookHandler.HomePage)
	router.HandleFunc("/api/book", bookHandler.Create).Methods("POST")
	router.HandleFunc("/api/book", bookHandler.GetAll).Methods("GET")
	router.HandleFunc("/api/book/update/{id}", bookHandler.Update).Methods("PUT")
	router.HandleFunc("/api/book/{id}", bookHandler.Delete).Methods("DELETE")
	router.HandleFunc("/api/book/{id}", bookHandler.Get).Methods("GET")

	server := http.Server{
		Addr:    conf.Host + ":" + conf.Port,
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Server cant running")
	}
}
