package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/federated-covid-prediction/coordinator/actors"
	"github.com/rruzicic/federated-covid-prediction/peer/controllers"
	"github.com/rruzicic/federated-covid-prediction/peer/services"
)

func main() {
	go ginSetup()
	actors.SetupLeaderCoordinator()

	fmt.Println("Press [TAB] when your agent reaches EXIT state")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\t')
}

func ginSetup() {
	address, _ := services.GetYourAddress()

	r := gin.New()
	r.POST("/coordinator-pid", controllers.HandleCoordinatorPID)
	r.Run(":" + strconv.Itoa(address.Port+1000))
}
