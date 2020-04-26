package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Match struct {
	ID int `json:"id"`
	Score string `gorm:"type:varchar(45)" json:"score"`
	MapId uint `gorm:"not null" json:"map_id"`
	Duration uint `gorm:"not null" json:"duration"`
	Rounds []Round `json:"-"`
	Players []Player `gorm:"many2many:match_players" json:"-,omitempty"`
}

type Round struct {
	gorm.Model
	MatchId uint `gorm:"not null"`
	Winner string `gorm:"not null"`
	WinType string `gorm:"not null"`
	Duration time.Time `gorm:"not null"`
	RoundNum int `gorm:"not null"`
	Firearms []Firearm `gorm:"many2many:round_firearms"`
}

type Player struct {
	gorm.Model
	Nickname string `gorm:"not null"`
	Rank string
	Score string `gorm:"not null"`
}

type Firearm struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Damage int `gorm:"not null"`
	Price int `gorm:"not null"`
	Accuracy int `gorm:"not null"`
	FireRate int `gorm:"not null"`
	Side string `gorm:"not null"`
}

type Grenade struct {
	gorm.Model
	Name string `gorm:"not null"`
	Damage int `gorm:"not null"`
	Price int `gorm:"not null"`
	Range int `gorm:"not null"`
	TimeToExplode float32 `gorm:"not null"`
}

type Map struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Type string `gorm:"not null"`
}
