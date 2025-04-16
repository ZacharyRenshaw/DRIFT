package maploader

import (
	"os"
	"strconv"
	"drift/types"
	"encoding/csv"
)

func LoadMap(model *types.Model){

	if m.Map == nil {
		m.Map = make(map[int]map[int]int)
	}

	file, err := os.Open("C:/Go/Programs/Drift/modules/", model.Parameters["map"], ".csv")
	if err != nil {}
	defer file.Close()

    // land = 1
    // coastal water = 2
    // open water = 3
    // high mountain = 4
    // desert = 5
    // ice = 6

	reader := csv.NewReader(file)

   // Read header row
	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV header: %v", err)
	}

   latIdx, lonIdx, terrainTypeIdx := -1, -1, -1
	for i, column := range header {
		switch column {
		case "lat":
			latIdx = i
		case "lon":
			lonIdx = i
		case "terrain_type":
			terrainTypeIdx = i
		}
	}

	// Ensure all required columns were found
	if latIdx == -1 || lonIdx == -1 || terrainTypeIdx == -1 {
		return fmt.Errorf("CSV is missing required columns")
	}

	// Read data rows
	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file or error
	}

}