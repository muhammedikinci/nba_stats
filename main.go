package main

import (
	"nba_stats/repository"
	"nba_stats/simulation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	generateData()

	e := echo.New()
	repository := repository.NewRepository()
	sim := simulation.NewSimulation(repository)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
	}))

	e.POST("/simulation/start", sim.StartSimulate)
	e.GET("/simulation/current-round", sim.GetCurrentRound)
	e.GET("/simulation/get-all-rounds", sim.GetAllRounds)
	e.DELETE("/simulation/stop", sim.StopSimulation)
	e.GET("/simulation", sim.GetState)

	e.Logger.Fatal(e.Start(":3456"))
}
