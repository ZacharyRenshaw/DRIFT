package maploader

import (
	"drift/modules/csvutils"
	"drift/types"
	"fmt"
)

// A few healthy guardrails
const minValidLat = -90
const maxValidLat = 90
const minValidLon = -180
const maxValidLon = 180

// Terrain types
// TODO you might want to move this somewhere else
type Terrain int

const (
	InvalidTerrainLow Terrain = iota // 0
	Land                             // 1
	CoastalWater                     // 2
	OpenWater                        // 3
	HighMountain                     // 4
	Desert                           // 5
	Ice                              // 6
	InvalidTerrainHigh
)

// Load the map from a CSV file and populate the model's Map map.
func LoadMap(model *types.Model, mapRoot string) error {
	minLat, minLon, maxLat, maxLon := 0, 0, 0, 0
	// Initialize the map if it's nil
	if model.Map == nil {
		model.Map = make(map[int]map[int]int)
	}

	// Derive the filename from the model and load the CSV file
	filename := fmt.Sprintf("%s_map.csv", model.MapName)
	csvLoader := csvutils.CSVLoader{
		FileName:   filename,
		Dir:        mapRoot,
		MinRecords: 2,
	}
	records, err := csvLoader.LoadCSV()
	if err != nil {
		return err
	}

	// Skip the header row and process each record
	for _, record := range records[1:] {
		// Ensure the record has at least 3 fields
		err := csvLoader.CheckRecord(record, 3)
		if err != nil {
			return err
		}

		// Each record should have three fields: latitude, longitude, and terrain type
		lat, err := csvLoader.Atoi(record, 0)
		if err != nil {
			return err
		}
		if lat < minValidLat || lat > maxValidLat {
			return csvutils.ErrInvalidField{
				CSVLoader: csvLoader,
				Record:    record,
				Field:     0,
				Message:   "Latitude out of range",
			}
		}
		if lat < minLat {
			minLat = lat
		}
		if lat > maxLat {
			maxLat = lat
		}

		lon, err := csvLoader.Atoi(record, 1)
		if err != nil {
			return err
		}
		if lon < minValidLon || lon > maxValidLon {
			return csvutils.ErrInvalidField{
				CSVLoader: csvLoader,
				Record:    record,
				Field:     1,
				Message:   "Longitude out of range",
			}
		}
		if lon < minLon {
			minLon = lon
		}
		if lon > maxLon {
			maxLon = lon
		}

		terrain, err := csvLoader.Atoi(record, 2)
		if err != nil {
			return err
		}
		if terrain <= int(InvalidTerrainLow) || terrain >= int(InvalidTerrainHigh) {
			return csvutils.ErrInvalidField{
				CSVLoader: csvLoader,
				Record:    record,
				Field:     2,
				Message: fmt.Sprintf("Terrain type out of range [%d, %d]",
					InvalidTerrainLow+1,
					InvalidTerrainHigh-1),
			}
		}

		if model.Map[lat] == nil {
			model.Map[lat] = make(map[int]int)
		}
		model.Map[lat][lon] = terrain
	}
	model.FreeParameters["minLat"] = minLat
	model.FreeParameters["minLon"] = minLon
	model.FreeParameters["maxLat"] = maxLat
	model.FreeParameters["maxLon"] = maxLon

	latRange := maxLat - minLat + 1
	lonRange := maxLon - minLon + 1
	width := model.FreeParameters["map_width"]
	height := model.FreeParameters["map_height"]
	model.FreeParameters["tileWidth"] = width / lonRange
	model.FreeParameters["tileHeight"] = height / latRange

	return nil
}
