package controllers

import (
	"auth/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateHistory(w http.ResponseWriter, r *http.Request) {

	history := &models.History{}
	json.NewDecoder(r.Body).Decode(history)

	createdHistory := db.Create(history)
	var errMessage = createdHistory.Error

	if createdHistory.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdHistory)
}

func GetHistory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var history models.History
	db.First(&history, id)
	json.NewEncoder(w).Encode(&history)
}

func DeleteHistory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var history models.History
	db.First(&history, id)
	db.Delete(&history)
	json.NewEncoder(w).Encode("History deleted")
}

func GetAllHistory(w http.ResponseWriter, r *http.Request) {
	var history []models.History
	params := mux.Vars(r)
	var id = params["id"]
	if err := db.Where("User_ID = ?", id).Find(&history).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "User id not found"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	json.NewEncoder(w).Encode(history)
}

