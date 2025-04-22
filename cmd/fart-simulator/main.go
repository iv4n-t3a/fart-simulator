package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/iv4n-t3a/fart-simulator/internal/experiments"
)

var CLI struct {
  Experiment string `short:"e" help:"Experiment to run"`
}

func main() {
	kong.Parse(&CLI)

	switch CLI.Experiment {
  case "":
		experiments.RunSimpleSimulation()
	case "simple-simulation":
		experiments.RunSimpleSimulation()
	case "visualisation":
		experiments.RunVisualisation()
	default:
		panic(fmt.Sprintf("Unknown experiment %s, running default", CLI.Experiment))
	}
}
