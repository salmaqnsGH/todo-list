package main

import (
	"fmt"
	"log"

	"github.com/salmaqnsGH/todo-list/activity"
	"github.com/salmaqnsGH/todo-list/handler"
	"github.com/salmaqnsGH/todo-list/todo"
	"github.com/salmaqnsGH/todo-list/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dsn := config.MySqlUser + ":" + config.MySqlPassword + "@tcp(" + config.MySqlHost + ":" + config.MySqlPort + ")/" + config.MySqlDBName + "?charset=utf8mb4&parseTime=True&loc=Local"
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
	router.DELETE("/activity-groups/:id", activityHandler.DeleteActivity)
	router.PATCH("/activity-groups/:id", activityHandler.UpdateActivity)

	todoRepository := todo.NewRepository(db)
	todoService := todo.NewService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService)

	router.GET("/todo-items", todoHandler.GetTodos)
	router.GET("/todo-items/:id", todoHandler.GetTodoById)
	router.POST("/todo-items", todoHandler.CreateTodo)
	router.DELETE("/todo-items/:id", todoHandler.DeleteTodo)
	router.PATCH("/todo-items/:id", todoHandler.UpdatedTodo)

	router.Run(":3030")
}
