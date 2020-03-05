package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4001"}
	r.Use(cors.New(config))
	r.GET("/v1/player", player)
	r.GET("/v1/jleague/j1/result", jleague)
	r.GET("/v1/jleague/j1/result/:section", resultSection)
	r.Run(":5000")
}

func player(c *gin.Context) {
	p, err := ioutil.ReadFile("./player.json")
	if err != nil {
		print(err)
	}
	var pl PLAYERS
	json.Unmarshal(p, &pl)
	c.JSON(200, pl)
}

//PLAYERS DO
type PLAYERS struct {
	PLAYERS []DATA `json:"player"`
}

//DATA DO
type DATA struct {
	I    int    `json:"u"`
	NAME string `json:"name"`
}

//RESULTS D
type RESULTS struct {
	RESULTS []RESULT `json:"result"`
}

//RESULT D
type RESULT struct {
	ID      int    `json:"id"`
	SEC     int    `json:"section"`
	HOME    string `json:"home"`
	HS      int    `json:"home_score"`
	AWAY    string `json:"away"`
	AS      int    `json:"away_score"`
	DATE    string `json:"date"`
	VISITOR string `json:"visitor"`
	KO      string `json:"ko"`
	STADIUM string `json:"stadium"`
}

func jleague(c *gin.Context) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=name password=pass dbname=jleague sslmode=disable")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT * FROM result.match_result")
	if err != nil {
		panic(err)
	}
	results := []RESULT{}
	result := RESULT{}
	for rows.Next() {
		rows.Scan(&result.ID, &result.SEC, &result.HOME, &result.HS, &result.AWAY, &result.AS, &result.DATE, &result.VISITOR, &result.KO, &result.STADIUM)
		results = append(results, result)
	}
	c.JSON(200, results)
}

func resultSection(c *gin.Context) {
	section := c.Param("section")
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=ken41 password=kimosken41 dbname=jleague sslmode=disable")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT * FROM result.match_result where section = $1", section)
	if err != nil {
		panic(err)
	}
	results := []RESULT{}
	result := RESULT{}
	for rows.Next() {
		rows.Scan(&result.ID, &result.SEC, &result.HOME, &result.HS, &result.AWAY, &result.AS, &result.DATE, &result.VISITOR, &result.KO, &result.STADIUM)
		results = append(results, result)
	}
	c.JSON(200, results)
}
