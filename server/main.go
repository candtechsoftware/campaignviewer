package main

import (
	"log"
	"net/http"

	"github.com/candtechsoftware/campaignapp/controllers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	r.HandleFunc("/campaign", controllers.GetCampaign).Methods("GET")
	r.HandleFunc("/campaign", controllers.CreateCampaign).Methods("POST")
	r.HandleFunc("/campaign/{id}", controllers.UpdateCampaign).Methods("POST")
	r.HandleFunc("/campaign/{id}", controllers.DeleteCampaign).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", r))
}
