package handlers

import (
    "encoding/json"
    "hysteria-manager/models"
    "hysteria-manager/services"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
)

func CreateInstance(w http.ResponseWriter, r *http.Request) {
    var instance models.Instance
    if err := json.NewDecoder(r.Body).Decode(&instance); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    err := services.CreateInstance(&instance)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(instance)
}

func GetInstance(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    instance, err := services.GetInstance(uint(id))
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(instance)
}

// 类似地实现 UpdateInstance 和 DeleteInstance
