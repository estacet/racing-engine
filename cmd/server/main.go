package main

import (
	"context"
	"log"
	"net/http"

	"github.com/greeflas/racing-engine-backend/pkg/db"

	"github.com/greeflas/racing-engine-backend/pkg/validator"

	"github.com/greeflas/racing-engine-backend/internal/repository"
	"github.com/greeflas/racing-engine-backend/internal/service"

	"github.com/greeflas/racing-engine-backend/internal/handler"
	"github.com/greeflas/racing-engine-backend/pkg/server"
)

func main() {
	ctx := context.Background()

	validate, err := validator.New()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	raceRepository := repository.NewRaceRepository(conn)
	driverRepository := repository.NewDriverRepository(conn)

	raceService := service.NewRaceService(raceRepository)
	driverService := service.NewDriverService(driverRepository)

	mux := http.NewServeMux()

	raceCRUDHandler := handler.NewRaceCRUDHandler(raceService, validate)
	raceCRUDHandler.RegisterRoutes(mux)

	driverCRUDHandler := handler.NewDriverCRUDHandler(driverService, validate)
	driverCRUDHandler.RegisterRoutes(mux)

	apiServer := server.NewAPIServer(mux)
	if err := apiServer.Start(); err != nil {
		log.Panic(err)
	}
}
