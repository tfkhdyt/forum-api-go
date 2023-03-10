package main

import (
	"log"
	"os"

	"github.com/tfkhdyt/forum-api-go/config"
	"github.com/tfkhdyt/forum-api-go/domain"
	httpUserController "github.com/tfkhdyt/forum-api-go/user/controller/http"
	postgresUserRepo "github.com/tfkhdyt/forum-api-go/user/repository/postgres"
	userService "github.com/tfkhdyt/forum-api-go/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// db setup
	db, err := gorm.Open(postgres.Open(config.GetPostgresConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to db:", err.Error())
	}
	db.AutoMigrate(&domain.User{})

	// gin setup
	r := gin.New()
	ginMode := os.Getenv("GIN_MODE")
	r.Use(gin.Recovery())
	if ginMode != "release" {
		r.Use(gin.Logger())
	}

	userRepo := postgresUserRepo.New(db)
	userService := userService.New(userRepo)
	httpUserController.New(r, userService)

	log.Fatalln(r.Run(":3000"))
}
