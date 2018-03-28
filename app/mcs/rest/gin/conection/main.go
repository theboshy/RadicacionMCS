package main

import (
	"context"
	"github.com/gin-gonic/gin"
	//"fmt"
	"net/http"
	"time"
	//"github.com/fvbock/endless"
	//"log"
	"../serviceMapping"
	"fmt"
	"log"
	"os"
	"os/signal"
	"github.com/gin-contrib/cors"
)

const serverPort = ":8283"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	return router
}

func main() {
	router := SetupRouter()
	//-acors origin
	router.Use(cors.Default())
	serviceMapping.MapRouterGroup(router)
	//router.Run(serverPort)

	srv := &http.Server{
		Addr:    serverPort,
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//log.Fatalf("listen: %s\n", err)
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Apagando servidor ...")
	//log.Println("Apagando servidor ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
		fmt.Printf("Servidor apagado")
	}

	//endless.ListenAndServe(":8956", router)
}
