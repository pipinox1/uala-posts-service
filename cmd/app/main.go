package main

import (
	"fmt"
	httpServer "net/http"
	"os"
	"os/signal"
	"syscall"
	"uala-posts-service/cmd/http"
	"uala-posts-service/config"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error loading config file: %w", err))
	}

	dependencies, err := config.BuildDependencies(*cfg)
	if err != nil {
		panic(fmt.Errorf("fatal error building dependencies: %w", err))
	}

	router := http.SetupRouterAndRoutes(cfg, dependencies)
	go func() {
		fmt.Println("starting server on port: ", cfg.Port)
		if err := httpServer.ListenAndServe(
			fmt.Sprintf(":%s", cfg.Port),
			router,
		); err != nil {
			fmt.Println("error starting server")
			panic(err)
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit

	// Gracefully shutdown
	dependencies.EventPublisher.Close()
}
