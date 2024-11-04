package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Bold = "\033[1m"

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "year",
		Short: "Select the year to find upcoming movies",
		Long:  `Select a year between 2024 and 2029 to find upcoming movies.`,
		Run:   movieCommandHandler("year", "Select a "+Bold+Magenta+"YEAR"+Reset+" to look at upcoming movies ("+Green+"2024 - 2029"+Reset+"):\n"),
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "genre",
		Short: "Select your favourite genre to find upcoming movies",
		Long:  `Select a genre, such as Drama, Thriller, Horror to see upcoming movies.`,
		Run:   movieCommandHandler("genre", "Select a "+Bold+Cyan+"GENRE"+Reset+" to look at upcoming movies (e.g. "+Green+"Drama, Thriller, Horror"+Reset+"):\n"),
		// Action, Adult, Adventure, Animation, Biography, Comedy, Crime, Documentary, Drama, Family,
		// Fantasy, Film-Noir, History, Horror, Musical, Mystery, Romance
		// Sci-Fi, Short, Sport, Thriller, War, Western
	})
}

// Define the structure of the JSON response
type Response struct {
	Results []Movie `json:"results"`
}

type Movie struct {
	TitleText struct {
		Text string `json:"text"`
	} `json:"titleText"`
	ReleaseYear struct {
		Year int `json:"year"`
	} `json:"releaseyear"`
}

type Config struct {
	ApiKey      string `json:"api_key"`
	ApiHost     string `json:"api_host"`
	LogFilePath string `json:"log_file_path"`
}

type MovieLog struct {
	SearchQuery string
	Title       string
	Year        int
}

func loadConfig() (Config, error) {
	var config Config

	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return config, err
	}

	if err := json.Unmarshal(file, &config); err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return config, err
	}
	return config, nil
}

// Pointer to the reader variable for user input
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	// Read everything the user inputs up until they hit enter - new line key
	input, error := r.ReadString('\n')

	// Capitalise first letter of Input
	caser := cases.Title(language.English)
	title := caser.String(input)

	// Get rid of white space
	return strings.TrimSpace(title), error
}

func movieCommandHandler(param, prompt string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		input, err := getInput(prompt, reader)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		fmt.Printf("Getting movies for the "+Bold+Yellow+"%s"+Reset+" "+Blue+"%v\n"+Reset, param, input)

		config, err := loadConfig()
		if err != nil {
			log.Fatal("Error loading config:", err)
		}

		// Debugging line to check config file is being called correctly
		// fmt.Printf("API Key: %s\nAPI Host: %s\nLog File Path: %s\n", config.ApiKey, config.ApiHost, config.LogFilePath)

		fetchMovies(param, input, config.LogFilePath)
	}
}

// func fetchMovies(param, value string) {
func fetchMovies(param, value, fileName string) {
	config, err := loadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
		return
	}

	// Search movies from 2024 - 2029, ten at a time, scroll through pages for other 10 for that year
	// 2024 and 2025 have a lot of movies so may need to specify
	// 2025 = 80
	url := fmt.Sprintf("https://moviesdatabase.p.rapidapi.com/titles/x/upcoming?%s=%s", param, value)

	// Create a new HTTP client with a timeout
	client := &http.Client{Timeout: 20 * time.Second}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating a new HTTP request:", err)
		return
	}

	// Set headers
	req.Header.Add("x-rapidapi-key", config.ApiKey)
	req.Header.Add("x-rapidapi-host", config.ApiHost)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "MyApp/1.0")

	// Making the HTTP request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer res.Body.Close()

	// Check if the request was successful
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Received non-200 response: %d\n", res.StatusCode)
		fmt.Printf("Making request to URL: %s\n", url)
		fmt.Printf("Request headers: %v\n", req.Header)
		return
	}

	// Read the body of the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading the body:", err)
		return
	}

	// response doesnt hold data yet but will, It is setting a variable to be linked to the information it receives from movies struct
	// after decoded from json
	var response Response
	// &response points to the empty variable to store the JSON data that decoded into a GO struct
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Could not unmarshal response:", err)
	}

	if len(response.Results) == 0 {
		fmt.Println("No upcoming movies found.")
		return
	}

	// Print the movie titles in a list format
	for _, movie := range response.Results {
		fmt.Printf("- %s\n", movie.TitleText.Text)

		movieLog := MovieLog{
			SearchQuery: value,
			Title:       movie.TitleText.Text,
			Year:        movie.ReleaseYear.Year,
		}

		// Save Movie data to file
		saveToFile(movieLog, fileName)
	}
}

// Each title is passed to the file
// func saveToFile(data interface{}, filePath string) {
func saveToFile(logEntry MovieLog, filePath string) {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Printf("Failed to create directory: %v", err)
		return
	}
	// Opening a file called movie.txt (os.OpenFile), Append - adds info to the file without overwriting it
	// Create makes the file if it doesnt exist
	// Wronly - Can only write to the file and not read
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Other file log examples
	// logEntry := fmt.Sprintf("[%s] %s", data, time.Now().Format("2006-01-02 15:04:05"))
	// logLine := fmt.Sprintf("- \"%s\", Year: \"%d\", Search Query: \"%s\", [%s]",	logEntry.Title, logEntry.Year, logEntry.SearchQuery, time.Now().Format("2006-01-02 15:04:05"))

	// Title width 40 chars, year width 4 chars, search width 10 chars
	logLine := fmt.Sprintf("| %-55s | %-4d | %-10s |", logEntry.Title, logEntry.Year, logEntry.SearchQuery)

	_, err = fmt.Fprintln(writer, logLine)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	writer.Flush()
}
