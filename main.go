package main

import (
	"encoding/csv"
	"flag"
	"log"
	"net/http"
	"os"

	"pokemons/config"
	"pokemons/controller"
	"pokemons/router"
	"pokemons/service/api"
	serviceCSV "pokemons/service/csv"
	"pokemons/usecase"

	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

// Exit Errors
const (
	ExitErrorLoadingConfig = iota
	ExitErrorCreatingLogger
	ExitErrorLoadingCSVFile
	ExitErrorCreatingPokeonClient
)

func main() {
	var configFile string
	flag.StringVar(
		&configFile,
		"public-config-file",
		"config.yml",
		"Path to public config file",
	)
	flag.Parse()

	// Read the configFile
	config, err := config.Load(configFile)
	if err != nil {
		log.Fatal("Failed to load config: %w", err)
		os.Exit(ExitErrorLoadingConfig)
	}

	// Create Logger
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatal("Failed creating logger: %w", err)
		os.Exit(ExitErrorCreatingLogger)
	}

	logger := logrus.New()
	logger.SetLevel(level)
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{}

	// Create Reader
	rf, err := os.Open(config.Pokedex)
	if err != nil {
		os.Exit(ExitErrorLoadingCSVFile)
	}
	defer rf.Close()

	// Create Writter
	wf, err := os.OpenFile(config.Pokedex, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		os.Exit(ExitErrorLoadingCSVFile)
	}
	defer wf.Close()

	csvw := csv.NewWriter(wf)

	// Create Pokemon Client
	pokemonClient, err := api.New(logger, config.Services.PokemonAPI.Host,
		config.Services.PokemonAPI.Timeout, csvw)
	if err != nil {
		logger.WithError(err).
			Error("Creating Pokemon Client")
		os.Exit(ExitErrorCreatingPokeonClient)
	}

	// Create  csv service
	csvClient := serviceCSV.New(logger, rf, csvw)

	// Create Usecase
	usecase := usecase.New(pokemonClient, csvClient)

	// Create Controller
	controller := controller.New(usecase, logger, render.New())

	// Create Router
	router := router.New(controller)

	log.Fatal(http.ListenAndServe(":3030", router))
}
