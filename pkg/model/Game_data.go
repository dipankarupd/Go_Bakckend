package model

import (
	"github.com/dipankarupd/go-football-betting/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Game struct {
	gorm.Model

	Team1  string `gorm:"" json:"team1"`
	Team2  string `gorm:"" json:"team2"`
	Winner string `gorm:"" json:"winner,omitempty"`
}

// Initialize the database connection
func init() {
	config.Connect()
	db = config.GetDB()

	// Disable table auto-migration for the Game struct
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Game{})
}

// Create a game and add to db
func CreateGame(game *Game) {
	db.Create(&game)
}

// GetGames retrieves all games
func GetGames() []Game {
	var games []Game
	db.Order("id desc").Limit(5).Find(&games)
	return games
}

// GetGame retrive a single game
func GetGameById(Id int64) (*Game, *gorm.DB) {
	var getGame Game
	// running WHERE command in mysql
	db := db.Where("id=?", Id).Find(&getGame)
	return &getGame, db
}

func UpdateDb(Id int, Winner string) {
	game := Game{}
	db.First(&game, Id)

	// Update the Winner field
	game.Winner = Winner
	db.Save(&game)
}
