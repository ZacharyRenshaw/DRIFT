package initializemodel

import (
	"drift/modules/actuarialloader"
	"drift/modules/chromosomeloader"
	"drift/modules/maploader"
	"drift/modules/paramloader"
	"drift/modules/save"
	"drift/types"
	"fmt"
	"math"
)

// Initializes the model based on the configuration files.
func InitializeModel(configRoot string, mapRoot string) (*types.Model, error) {
	model := &types.Model{
		Parameters:     make(map[string]float64),
		PlotFlags:      make(map[string]bool),
		ChromosomeArms: make(map[int]map[int][]int),
		DeathRisk:      make(map[int]float64),
		CumulativeProb: make(map[int]float64),
		FreeParameters: make(map[string]int),
		Map:            make(map[int]map[int]int),
	}

	// Attempt to load each config file. Failure will be fatal.
	err := paramloader.LoadParameters(model, configRoot)
	if err != nil {
		return nil, err
	}
	err = chromosomeloader.LoadChromosomeArms(model, configRoot)
	if err != nil {
		return nil, err
	}
	err = actuarialloader.LoadActuarialTable(model, configRoot)
	if err != nil {
		return nil, err
	}
	err = maploader.LoadMap(model, mapRoot)
	if err != nil {
		return nil, err
	}

	// Calculate derived values
	model.Parameters["mu_sig_figs"] = math.Pow(1, model.Parameters["mu_sig_figs"])

	// Prepare output files
	save.SaveHeaders(model.ModelName)

	// Initialize free parameters
	model.FreeParameters["indID"] = 0         // Starting ID for individuals
	model.FreeParameters["seed"] = -1         // No seed initially
	model.FreeParameters["last_pop_size"] = 0 // Required for growth rate calculations
	model.FreeParameters["mutID"] = 0         // Starting ID for mutations

	//PrintModel(model)                       // For doublechecking purposes

	return model, nil
}

func PrintModel(model *types.Model) {
	fmt.Println("Model Name:", model.ModelName)
	fmt.Println("Parameters:", model.Parameters)
	fmt.Println("Free Parameters:", model.FreeParameters)
	fmt.Println("Plot Flags:", model.PlotFlags)
	fmt.Println("Chromosome Arms:", model.ChromosomeArms)
	fmt.Println("Death Risk:", model.DeathRisk)
	fmt.Println("Cumulative Probability:", model.CumulativeProb)
}
