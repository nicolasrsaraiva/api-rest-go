package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicolasrsaraiva/api-rest-go/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", r))
}
