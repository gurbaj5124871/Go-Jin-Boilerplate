package main

import (
	"context"
	"fmt"
	"go-gin-boilerplate/config"
	"go-gin-boilerplate/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	goEnv := os.Getenv("go_env")
	if goEnv == "" {
		goEnv = "dev"
	}
	conf, err := config.InitiliseConfig(goEnv)
	if err != nil {
		log.Fatalf("Error while initilising environment: %s\n", err)
		os.Exit(1)
	}

	r := routes.InitiliseRoutes()
	addr := fmt.Sprint(":", conf.Port)
	log.Println(addr)
	server := &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Start server to listen and handle gracefull shutdown
	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("... Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.GraceShutDown)*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Error:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("Server exiting")
		// Break connection to db with a timeout of 5 seconds here
	}
}
