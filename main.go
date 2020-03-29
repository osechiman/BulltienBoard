package main

import "vspro/drivers/web/api"

func main() {
	api.Listen()
}

// func main() {
// 	mrouter := melody.New()
// 	mrouter.HandleMessage(func(s *melody.Session, msg []byte) {
// 		mrouter.Broadcast(msg)
// 		d, _ := strconv.Atoi(string(msg))
// 		battle.Participant[s] = User{battle.Participant[s].ID, battle.Participant[s].LifePoint - d}
// 		l := strconv.Itoa(battle.Participant[s].LifePoint)
// 		dm := []byte("user:" + battle.Participant[s].ID + " collect answer. remaining is " + l)
// 		mrouter.Broadcast(dm)
// 	})
// 	mrouter.HandleConnect(func(s *melody.Session) {
// 		u := User{uuid.New().String(), life}
// 		battle.Participant[s] = u
// 	})
// 	router.Run(":8080")
// }
