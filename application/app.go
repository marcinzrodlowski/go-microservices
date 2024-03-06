package application

import (
	"context"
	"fmt"
	"net/http"
)

// App is a struct holding types
type App struct {
	router http.Handler
}

// New is a constructor which returns pointer to App (we want to be able to modify it)
func New() *App {
	app := &App{ // creates a new instance of the App struct and takes the address of it
		router: loadRoutes(), // defines the routes for the router
	}
	return app
}

// Start method associated with App type
func (a *App) Start(ctx context.Context) error {
	server := &http.Server{ // creates a new http.Server instance and takes address of it
		Addr:    ":3000",
		Handler: a.router, // incoming HTTP requests will be routed according to the routes defined in the router field
	}

	err := server.ListenAndServe() // starts the server asynchronously
	if err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	return nil
}
