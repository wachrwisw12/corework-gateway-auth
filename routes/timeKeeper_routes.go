package routes

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/wachrwisw12/corework-gateway-auth/handler/timekeeper_handler"
)

func SetupTimeKeeper(timeKeep fiber.Router) {
	timeKeep.Post("/timekeeperhistory", handler.Timekeeperhandler)
}
