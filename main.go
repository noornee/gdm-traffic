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
	default:
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

	// if the dit flag is passed
	if durationInTraffic {
		var matrixData handler.MatrixAPIResponse
		err = json.Unmarshal(body, &matrixData)
		if err != nil {
			utils.ErrorLog.Fatal("error unmarshaling data", err.Error())
		}

		// Duratiom in traffic text
		dit := matrixData.Rows[0].Elements[0].DurationInTraffic.Text
		if dit != "" {
			utils.InfoLog.Printf("Travel time for origin '%s' and destination '%s' is %s\n", origin, destination, dit)
			return
		}
	}

	// else
	// just print the full response
	utils.InfoLog.Println(string(body))
}

// handleNoFlags handles when no flag is passed
func handleNoFlags() {
	utils.ErrorLog.Println("No flags provided")
	fmt.Printf("Run %s --help to see the available flags\n", os.Args[0])
	// flag.PrintDefaults()
	os.Exit(1)
}

// handleEmptyFlags handles when the flag passed is an empty string
func handleEmptyFlags() {
	utils.ErrorLog.Println("Both origin (--origin) and destination (--dest) flags are required")
	fmt.Println("Usage:")
	flag.PrintDefaults()
	os.Exit(1)
}

// handleInvalidFlagValue This handles a situation when the value passed to the flags are invalid
func handleInvalidFlagValue() {
	// Split origin and destination values by comma
	originCords := strings.Split(origin, ",")
	destinationCords := strings.Split(destination, ",")

	// Check if the number of coordinates is not 2 for both origin and destination
	if len(originCords) != 2 {
		utils.ErrorLog.Fatal("Invalid origin value: Please provide longitude and latitude separated by a comma")
	}
	if len(destinationCords) != 2 {
		utils.ErrorLog.Fatal("Invalid destination value: Please provide longitude and latitude separated by a comma")
	}

	// Longitude and latitude values are typically floating point numbers so
	// we need to check if the passed values are valid floating point numbers

	// Origin
	// Checks if the passed longitude and latitude values for the Origin are valid floats
	if _, err := strconv.ParseFloat(originCords[0], 64); err != nil {
		utils.ErrorLog.Fatal("Invalid longitude value for origin: Please provide a valid coordinate")
	}
	if _, err := strconv.ParseFloat(originCords[1], 64); err != nil {
		utils.ErrorLog.Fatal("Invalid latitude value for origin: Please provide a valid coordinate")
	}

	// Destination
	// Checks if the passed longitude and latitude values for the Destination are valid floats
	if _, err := strconv.ParseFloat(destinationCords[0], 64); err != nil {
		utils.ErrorLog.Fatal("Invalid longitude value for destination: Please provide a valid coordinate")
	}
	if _, err := strconv.ParseFloat(destinationCords[1], 64); err != nil {
		utils.ErrorLog.Fatal("Invalid latitude value for destination: Please provide a valid coordinate")
	}
}
