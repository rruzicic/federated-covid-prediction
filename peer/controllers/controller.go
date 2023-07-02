package controllers

import (
	"log"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/gin-gonic/gin"
)

func HandleCoordinatorPID(ctx *gin.Context) {
	var coordinatorPID actor.PID
	if err := ctx.ShouldBindJSON(&coordinatorPID); err != nil {
		log.Println("Could not bind pid from json. Error: ", err)
		ctx.JSON(400, "Bad Request")
		return
	}
}
