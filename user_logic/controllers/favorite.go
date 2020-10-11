package controllers

import (
	"auth/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateFavorite(w http.ResponseWriter, r *http.Request) {

	favorite := &models.Favorite{}
	json.NewDecoder(r.Body).Decode(favorite)

	createdFavorite := db.Create(favorite)
	var errMessage = createdFavorite.Error

	if createdFavorite.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdFavorite)
}

func GetFavorite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var favorite models.Favorite
	db.First(&favorite, id)
	json.NewEncoder(w).Encode(&favorite)
}

func DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var favorite models.Favorite
	db.First(&favorite, id)
	db.Delete(&favorite)
	json.NewEncoder(w).Encode("Favorite deleted")
}

func GetAllFavorite(w http.ResponseWriter, r *http.Request) {
	var favorite []models.Favorite
	params := mux.Vars(r)
	var id = params["id"]
	if err := db.Where("User_ID = ?", id).Find(&favorite).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "User id not found"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	json.NewEncoder(w).Encode(favorite)
}

