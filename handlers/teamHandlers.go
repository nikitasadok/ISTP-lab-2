package handlers

import (
	"CSGORest/initializers"
	"CSGORest/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func ListOneTeamHandler(w http.ResponseWriter, r *http.Request) {
	var team models.Team
	vars := mux.Vars(r)
	id := vars["id"]

	initializers.Db.Where("id = ?", id).Find(&team)

	err := json.NewEncoder(w).Encode(team)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", 500)
	}
}
