package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/federated-covid-prediction/coordinator/actors"
	"github.com/rruzicic/federated-covid-prediction/peer/controllers"
	"github.com/rruzicic/federated-covid-prediction/peer/services"
)

func main() {
	actors.SetupLeaderCoordinator()
	go ginSetup()
}

func ginSetup() {
	address, _ := services.GetYourAddress()

	r := gin.New()
	r.POST("/coordinator-pid", controllers.HandleCoordinatorPID)
	r.Run(":" + strconv.Itoa(address.Port+1000))
}
