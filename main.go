package main

import (
	"fmt"
	"log"

	"todo-list/activity"
	"todo-list/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Koneksi database berhasil!")

	activityRepository := activity.NewRepository(db)
	activityService := activity.NewService(activityRepository)
	activityHandler := handler.NewActivityHandler(activityService)

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/activity-groups", activityHandler.GetActivities)
	router.GET("/activity-groups/:id", activityHandler.GetActivityById)
	router.POST("/activity-groups", activityHandler.CreateActivity)

	router.Run()
}
