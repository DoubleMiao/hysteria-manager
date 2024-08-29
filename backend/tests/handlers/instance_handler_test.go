package handlers

import (
    "bytes"
    "encoding/json"
    "hysteria-manager/models"
    "hysteria-manager/services"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestCreateInstance(t *testing.T) {
    instance := models.Instance{Name: "test-instance", Port: 8080}
    jsonValue, _ := json.Marshal(instance)
    req, err := http.NewRequest("POST", "/instances", bytes.NewBuffer(jsonValue))
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(CreateInstance)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }

    var createdInstance models.Instance
    json.Unmarshal(rr.Body.Bytes(), &createdInstance)
    if createdInstance.Name != instance.Name {
        t.Errorf("handler returned unexpected body: got %v want %v", createdInstance.Name, instance.Name)
    }
}
