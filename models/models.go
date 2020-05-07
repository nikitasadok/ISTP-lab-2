package models

import (
	"errors"
	//"errors"
	//"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm"
	"time"
)

type Tournament struct {
	ID        int
	Name      string     `gorm:"not null;unique"`
	StartDate *time.Time `gorm:"not null"`
	EndDate   *time.Time `gorm:"not null"`
	Teams     []Team     `gorm:"many2many:tournament_teams"`
	Matches   []Match
}

type Team struct {
	ID             int    `json:"id"`
	Name           string `gorm:"not null;unique"`
	FoundationYear uint   `gorm:"not null" json:"foundation_year"`
	Players        []Player
	Matches        []Match      `gorm:"many2many:team_matches"`
	Tournaments    []Tournament `gorm:"many2many:tournament_teams"`
}

type Match struct {
	ID           int     `json:"id"`
	Score        string  `gorm:"type:varchar(45)" json:"score"`
	MapId        uint    `gorm:"not null" json:"map_id"`
	Duration     uint    `gorm:"not null" json:"duration"`
	TournamentId uint    `gorm:"not null"`
	Winner       string  `gorm:"type:varchar(45)" json:"winner"`
	Rounds       []Round `json:"-"`
	Teams        []Team  `gorm:"many2many:team_matches"`
}

func (match Match) Validate(db *gorm.DB) {
	if match.Score == "" {
		db.AddError(errors.New("the score cannot be empty"))
	}
	if match.MapId == 0 {
		db.AddError(errors.New("the MapId cannot be empty"))
	}
	if match.Duration == 0 {
		db.AddError(errors.New("the duration cannot be zero"))
	}
	if match.Winner == "" {
		db.AddError(errors.New("match can't have no winner"))
	}
	if match.TournamentId == 0 {
		db.AddError(errors.New("Match cannot exist without a tournament"))
	}
}

type Round struct {
	ID       int
	MatchId  uint      `gorm:"not null"`
	Winner   string    `gorm:"not null"`
	WinType  string    `gorm:"not null"`
	Duration time.Time `gorm:"not null"`
	RoundNum int       `gorm:"not null"`
}

type Player struct {
	ID       int
	TeamId   uint   `gorm:"not null"`
	Nickname string `gorm:"not null"`
	Rank     string `gorm:"not null"`
	Score    string `gorm:"not null"`
}

type Map struct {
	ID   int
	Name string `gorm:"unique;not null"`
	Type string `gorm:"not null"`
}
