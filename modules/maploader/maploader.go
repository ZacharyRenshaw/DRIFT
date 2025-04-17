package maploader

import (
	"drift/types"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func LoadMap(model *types.Model) {
	minLat, minLon, maxLat, maxLon := 0, 0, 0, 0

	if model.Map == nil {
		model.Map = make(map[int]map[int]int)
	}

	filename := fmt.Sprintf("modules/maploader/%s_map.csv", model.MapName)
	file, err := os.Open(filename)
	if err != nil {
		print("Cannot open ", filename)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		print("Error reading map file")
		return
	}

	// land = 1
	// coastal water = 2
	// open water = 3
	// high mountain = 4
	// desert = 5
	// ice = 6

	for _, record := range records[1:] {

		lat, err := strconv.Atoi(record[0])
		if err != nil {
			print("error reading map data at ", record)
		}
		if lat < minLat {
			minLat = lat
		}
		if lat > maxLat {
			maxLat = lat
		}

		lon, err := strconv.Atoi(record[1])
		if err != nil {
			print("error reading map data at ", record)
		}
		if lon < minLon {
			minLon = lon
		}
		if lon > maxLon {
			maxLon = lon
		}

		terrain, err := strconv.Atoi(record[2])
		if err != nil {
			print("error reading map data at ", record)
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

}
