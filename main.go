package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/olahol/melody.v1"
	"net/http"
	"strconv"
)

// is not battle. which fast answer.
type Battle struct {
	ID          string
	Participant map[*melody.Session]User
}

type User struct {
	ID        string
	LifePoint int
}

var battle Battle = Battle{"id", map[*melody.Session]User{}}
var life int = 1000

func main() {
	router := gin.Default()
	mrouter := melody.New()

	router.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	router.GET("/ws", func(c *gin.Context) {
		mrouter.HandleRequest(c.Writer, c.Request)
	})

	mrouter.HandleMessage(func(s *melody.Session, msg []byte) {
		mrouter.Broadcast(msg)
		d, _ := strconv.Atoi(string(msg))
		battle.Participant[s] = User{battle.Participant[s].ID, battle.Participant[s].LifePoint - d}
		l := strconv.Itoa(battle.Participant[s].LifePoint)
		dm := []byte("user:" + battle.Participant[s].ID + " collect answer. remaining is " + l)
		mrouter.Broadcast(dm)
	})

	mrouter.HandleConnect(func(s *melody.Session) {
		u := User{uuid.New().String(), life}
		battle.Participant[s] = u
	})

	router.Run(":8080")
}
