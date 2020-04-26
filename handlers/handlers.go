package handlers

import (
	"CSGORest/initializers"
	"io/ioutil"
	"net/http"
	"CSGORest/models"
	"encoding/json"
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

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading your POST request", 500)
		return
	}

	if err := json.Unmarshal(body, &match); err != nil {
		http.Error(w, "Error decoding JSON", 500)
		return
	}

	initializers.Db.Create(&match)
}

func UpdateMatchHandler(w http.ResponseWriter, r *http.Request) {
	var match models.Match
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
	initializers.Db.Where("id = ?", match.ID).Find(&dbMatch)
	dbMatch.Duration = match.Duration
	dbMatch.MapId = match.MapId
	dbMatch.Rounds = match.Rounds
	dbMatch.Score = match.Score
	dbMatch.ID = match.ID
	dbMatch.Players = match.Players
	initializers.Db.Save(&dbMatch)
}

func DeleteMatchHandler(w http.ResponseWriter, r *http.Request) {
	var match models.Match

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
	initializers.Db.Where("id = ?", match.ID).Find(&toDelete)
	initializers.Db.Delete(&toDelete)
}

func ListPlayersHandler(w http.ResponseWriter, r *http.Request) {
	var players []models.Player

	initializers.Db.Find(&players)

	err := json.NewEncoder(w).Encode(players)

	if err != nil {
		http.Error(w, "Error encoding data to JSON", 500)
	}
}
