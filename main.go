package main

import (
	gateways "bulltienboard/adapters/gateways/maria_db"
	"bulltienboard/entities/valueobjects"
	"fmt"
)

func main() {
	db, _ := gateways.NewMariaDBRepository(
		"localhost",
		3306,
		"BulletinBoard",
		"root",
		"my-secret-pw",
	)
	row, _ := db.GetBulletinBoardByID(valueobjects.BulletinBoardID{})

	fmt.Printf("db: %v", db)
	fmt.Printf("row: %v", row)

	//api.Listen(di.InitializeRouter())
}
