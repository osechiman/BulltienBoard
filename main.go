package main

import (
	"bulltienboard/adapters/middlewares/di"
	"bulltienboard/drivers/web/api"
)

func main() {
	api.Listen(di.InitializeRouter())
}
