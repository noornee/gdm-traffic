package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/noornee/gdm-traffic/internal/handler"
	"github.com/noornee/gdm-traffic/internal/utils"
)

var (
	origin            string
	destination       string
	durationInTraffic bool
)

func init() {
	flag.StringVar(&origin, "origin", "", "accepts longitude and latitude (comma-separated). e.g.: --origin='12.345,-98.765'")
	flag.StringVar(&destination, "dest", "", "accepts longitude and latitude (comma-separated). e.g.: --dest='45.678,-123.456'")
	flag.BoolVar(&durationInTraffic, "dit", false, "prints duration in traffic")
}

func main() {
	// Parse flags
	flag.Parse()

	switch {
	case flag.NFlag() == 0:
		// If no flags are provided, handle appropriately
		handleNoFlags()
	case origin == "" || destination == "":
		// If either origin or destination is empty, handle appropriately
		handleEmptyFlags()
	case true:
		// Check and process the validity of provided flags
		handleInvalidFlagValue()
	}

	// get api key
	apiKey := os.Getenv("MATRIX_API_KEY")
	if apiKey == "" {
		utils.ErrorLog.Fatal("MATRIX_API_KEY not set in environment var")
	}

	// fetch traffic data
	body, err := handler.FetchDistanceMatrixData(origin, destination, apiKey)
	if err != nil {
		utils.ErrorLog.Fatal(err.Error())
	}

	var matrixData handler.MatrixAPIResponse
	err = json.Unmarshal(body, &matrixData)
	if err != nil {
		utils.ErrorLog.Fatal("error unmarshaling data", err.Error())
	}

	// FINAL

	// Duratiom in traffic text
	dit := matrixData.Rows[0].Elements[0].DurationInTraffic.Text

	// if tbe travel time flag is passed and the value of `dit` isnt an empty string
	if durationInTraffic == true && dit != "" {
		utils.InfoLog.Printf("Travel time for origin '%s' and destination '%s' is %s\n", origin, destination, dit)
		return
	}
	// else
	// just print the full response
	utils.InfoLog.Println(string(body))
}

// handles when no flag is passed
func handleNoFlags() {
	utils.ErrorLog.Println("No flags provided")
	fmt.Printf("Run %s --help to see the available flags\n", os.Args[0])
	// flag.PrintDefaults()
	os.Exit(1)
}

// handles when the flag passed is an empty string
func handleEmptyFlags() {
	utils.ErrorLog.Println("Both origin (--origin) and destination (--dest) flags are required")
	fmt.Println("Usage:")
	flag.PrintDefaults()
	os.Exit(1)
}

// This handles a situation when the value passed to the flags are invalid
func handleInvalidFlagValue() {
	// Split origin and destination values by comma
	originParts := strings.Split(origin, ",")
	destinationParts := strings.Split(destination, ",")

	// Check if the number of parts is not 2 for both origin and destination
	if len(originParts) != 2 {
		utils.ErrorLog.Fatal("Invalid origin value: Please provide longitude and latitude separated by a comma")
	}
	if len(destinationParts) != 2 {
		utils.ErrorLog.Fatal("Invalid destination value: Please provide longitude and latitude separated by a comma")
	}

	// Longitude and latitude values are typically floating point numbers so
	// we need to check if the passed values are valid floating point numbers

	// Origin
	// Checks if the passed longitude and latitude values for the Origin are valid floats
	if _, err := strconv.ParseFloat(originParts[0], 64); err != nil {
		utils.ErrorLog.Fatal("Invalid longitude value for origin: Please provide a valid coordinate")
	}
	if _, err := strconv.ParseFloat(originParts[1], 64); err != nil {
		utils.ErrorLog.Fatal("Invalid latitude value for origin: Please provide a valid coordinate")
	}

	// Destination
	// Checks if the passed longitude and latitude values for the Destination are valid floats
	if _, err := strconv.ParseFloat(destinationParts[0], 64); err != nil {
		utils.ErrorLog.Fatal("Invalid longitude value for destination: Please provide a valid coordinate")
	}
	if _, err := strconv.ParseFloat(destinationParts[1], 64); err != nil {
		utils.ErrorLog.Fatal("Invalid latitude value for destination: Please provide a valid coordinate")
	}
}
