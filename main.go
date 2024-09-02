package main

import (
	"blog-api/controllers"
	"blog-api/db"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "blog-api/db"
	"blog-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	// gin.SetMode(gin.ReleaseMode)
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	port := os.Getenv("SERVER_PORT")
	if port == ""{
		port = "8080"
	}

	db.InitDB()

	router := gin.Default()

	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"status":"ok",
			"message":"Server chal rha hai bhai",
		})
	})


	userController := &controllers.UserController{DB: db.DB}
	routes.UserRoutes(router, userController)

	server := &http.Server{
		Addr: ":" + port,
		Handler: router,
	}

	go func (){
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Cloud not start server: %v", err)
		}
	}()

	log.Printf("Server is running on port %s...", port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-quit
	log.Printf("shutting down server....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Printf("Server existing gracefully")

	// if err := router.Run(":"+port); err != nil {
	// 	log.Printf("Cloud not start server: %v", err)
	// } else {

	// 	log.Printf("Starting server on port %s...", port)
	// }
}