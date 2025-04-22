package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/iv4n-t3a/fart-simulator/internal/experiments"
)

var CLI struct {
	Experiment string `short:"e" help:"Experiment to run"`
	Dim        int    `short:"d" help:"Dimensions"`
}

func main() {
	kong.Parse(&CLI)

	switch CLI.Experiment {
	case "":
		experiments.RunSimpleSimulation(CLI.Dim)
	case "simple-simulation":
		experiments.RunSimpleSimulation(CLI.Dim)
	case "visualisation":
		experiments.RunVisualisation(CLI.Dim)
	default:
		panic(fmt.Sprintf("Unknown experiment %s, running default", CLI.Experiment))
	}
}
