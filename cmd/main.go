package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocli-demo",
	Short: "A simple CLI application for fetching data from a REST API",
	Run: func(cmd *cobra.Command, args []string) {
		// Handle the command logic here
		fmt.Println("Welcome to CLI App!")
	},
}

var fetchDataCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explain word meaning from Meriam Webster dictionary REST API",
	Run: func(cmd *cobra.Command, args []string) {
		// Call the function to fetch data from the REST API
		// fmt.Println("command args: ", args)
		// fmt.Println("cmd: ", cmd)
		wordForm, _ := cmd.Flags().GetString("wordForm")
		fmt.Println(wordForm)
		for _, word := range args {
			err := fetchData(word)
			if err != nil {
				log.Fatal(err)
			}

		}
	},
}

var (
	wordForm string
)

func init() {
	fetchDataCmd.Flags().StringVar(&wordForm, "wordForm", "", "Specify 'noun/verb/adjective' form for the word meaning")
}

func fetchData(word string) error {
	api_url := fmt.Sprintf("https://www.dictionaryapi.com/api/v3/references/collegiate/json/%s?key=%s", word, os.Getenv("API_KEY"))

	response, err := http.Get(api_url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code: %d", response.StatusCode)
	}

	// // Read and print the response body
	// // Modify this part based on the structure of your API response
	// // For simplicity, this example assumes a plain text response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var jsonData []map[string]interface{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return err
	}

	// fmt.Println("API Response:", jsonData)
	pronun := jsonData[0]["hwi"].(map[string]interface{})["prs"].([]interface{})[0].(map[string]interface{})["mw"]
	word_form := jsonData[0]["fl"]
	meaning := jsonData[0]["def"].([]interface{})[0].(map[string]interface{})["sseq"].([]interface{})[0].([]interface{})[0].([]interface{})[1].(map[string]interface{})["dt"].([]interface{})[0].([]interface{})[1]
	fmt.Println(pronun, word_form, meaning)
	return nil
}

func main() {
	// fmt.Println("Hello World")
	// Add the "fetch" command to the root command
	api_key := os.Getenv("API_KEY")

	if api_key == "" {
		err := fmt.Errorf("environment variable API_KEY is not set")
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd.AddCommand(fetchDataCmd)

	// Execute the CLI application
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
