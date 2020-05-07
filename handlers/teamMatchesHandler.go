package handlers

import (
	"CSGORest/initializers"
	"CSGORest/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func ListTeamsForMatch(w http.ResponseWriter, r *http.Request) {
	var teams []models.Team
	vars := mux.Vars(r)
	id := vars["id"]

	var teamIds []int64
	initializers.Db.Table("team_matches").Select("team_id").Where("match_id = ?", id).
		Pluck("team_id", &teamIds)

	var team models.Team
	for _, val := range teamIds {
		print(val)
		initializers.Db.Where("id = ?", val).Find(&team)
		teams = append(teams, team)
		team = models.Team{}
	}
	err := json.NewEncoder(w).Encode(teams)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", 500)
	}
}
