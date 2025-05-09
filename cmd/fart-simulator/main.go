package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/iv4n-t3a/fart-simulator/internal/experiments"
)

var CLI struct {
	Experiment string `short:"e" default:"simple-simulation" help:"Experiment to run"`
	Dim        int    `short:"d" default:"3"                 help:"Dimensions"`
}

func main() {
	kong.Parse(&CLI)

	switch CLI.Experiment {
	case "simple-simulation":
		experiments.RunSimpleSimulation(CLI.Dim)
	case "shrinking-container":
		experiments.RunShrinkingContainerSimulation(CLI.Dim)
	case "visualisation":
		experiments.RunVisualisation(CLI.Dim)
	case "adaptive-step":
		experiments.RunVisualisationWithAdaptiveStep(CLI.Dim)
	case "mixed-gas":
		experiments.RunMixedGasSimulation(CLI.Dim)
	case "hole":
		experiments.RunWallWithHole(CLI.Dim)
	default:
		panic(fmt.Sprintf("Unknown experiment %s, running default", CLI.Experiment))
	}
}
