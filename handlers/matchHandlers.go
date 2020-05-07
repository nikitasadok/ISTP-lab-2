package handlers

import (
	"CSGORest/initializers"
	"CSGORest/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func ListMatchesHandler(w http.ResponseWriter, r *http.Request) {
	var matches []models.Match

	initializers.Db.Find(&matches)

	err := json.NewEncoder(w).Encode(matches)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", 500)
	}
}

func AddMatchHandler(w http.ResponseWriter, r *http.Request) {
	var match models.Match

	var err error

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading your POST request", 500)
		return
	}

	if err = json.Unmarshal(body, &match); err != nil {
		http.Error(w, "Error decoding JSON", 500)
		return
	}

	errors := initializers.Db.Create(&match).GetErrors()

	if errors != nil {
		http.Error(w, "Some data is not in valid format", 422)
		for _, err := range errors {
			fmt.Fprintf(w, err.Error())
			fmt.Fprintf(w, "\n")
		}
		return
	}

	var matches []models.Match

	initializers.Db.Find(&matches)

	err = json.NewEncoder(w).Encode(matches)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", 500)
	}

}

func UpdateMatchHandler(w http.ResponseWriter, r *http.Request) {
	var match models.Match
	vars := mux.Vars(r)
	id := vars["id"]
	var err error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error when reading your request", 500)
		return
	}

	err = json.Unmarshal(body, &match)
	if err != nil {
		http.Error(w, "Error decoding JSON", 500)
		return
	}

	var dbMatch models.Match
	initializers.Db.Where("id = ?", id).Find(&dbMatch)
	dbMatch.Duration = match.Duration
	dbMatch.MapId = match.MapId
	dbMatch.Score = match.Score
	dbMatch.Winner = match.Winner
	dbMatch.TournamentId = match.TournamentId

	errors := initializers.Db.Save(&dbMatch).GetErrors()

	if errors != nil {
		http.Error(w, "Some data is not in valid format", 422)
		for _, err := range errors {
			fmt.Fprintf(w, err.Error())
			fmt.Fprintf(w, "\n")
		}
		return
	}

	var matches []models.Match

	initializers.Db.Find(&matches)

	err = json.NewEncoder(w).Encode(matches)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", 500)
	}
}

func DeleteMatchHandler(w http.ResponseWriter, r *http.Request) {
	var match models.Match
	vars := mux.Vars(r)
	id := vars["id"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error when reading your request", 500)
		return
	}

	err = json.Unmarshal(body, &match)
	if err != nil {
		http.Error(w, "Error decoding JSON", 500)
		return
	}

	var toDelete models.Match
	initializers.Db.Where("id = ?", id).Find(&toDelete)
	initializers.Db.Delete(&toDelete)

	var matches []models.Match

	initializers.Db.Find(&matches)

	err = json.NewEncoder(w).Encode(matches)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", 500)
	}
}

func ListPlayersHandler(w http.ResponseWriter, r *http.Request) {
	var players []models.Player

	initializers.Db.Find(&players)

	err := json.NewEncoder(w).Encode(players)

	if err != nil {
		http.Error(w, "Error encoding data to JSON", 500)
	}
}
