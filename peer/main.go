package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/federated-covid-prediction/coordinator/actors"
	"github.com/rruzicic/federated-covid-prediction/peer/controllers"
)

func main() {
	actors.SetupCoordinator()
	go ginSetup()
}

func ginSetup() {
	r := gin.New()
	r.POST("/coordinator-pid", controllers.HandleCoordinatorPID)
	r.Run(":8080")
}
