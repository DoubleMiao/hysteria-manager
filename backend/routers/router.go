package routers

import (
    "github.com/gorilla/mux"
    "hysteria-manager/handlers"
)

func SetupRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/instances", handlers.CreateInstance).Methods("POST")
    router.HandleFunc("/instances/{id}", handlers.GetInstance).Methods("GET")
    router.HandleFunc("/instances/{id}", handlers.UpdateInstance).Methods("PUT")
    router.HandleFunc("/instances/{id}", handlers.DeleteInstance).Methods("DELETE")
    return router
}
