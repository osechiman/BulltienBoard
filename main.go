package main

import (
	"vspro/adapters/middlewares/di"
	"vspro/drivers/web/api"
)

func main() {
	api.Listen(di.InitializeRouter())
}
