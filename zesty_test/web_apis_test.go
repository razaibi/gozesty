package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"gopkg.in/yaml.v2"
)

func loadConfig(filename string) (TestConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return TestConfig{}, err
	}

	var config TestConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return TestConfig{}, err
	}

	return config, nil
}

// Checks for API calls
func TestZ(t *testing.T) {

	folderPath := "calls"      // Replace with the path to your folder
	targetExtension := ".yaml" // Replace with the extension you want to filter

	files, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading folder:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == targetExtension {
			config, err := loadConfig(
				filepath.Join(
					folderPath,
					file.Name(),
				),
			)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			runConfigTests(t, config)
		}
	}

	if err != nil {
		t.Fatalf("Failed to load test configuration: %v", err)
	}

}

func runConfigTests(t *testing.T, config TestConfig) {
	for _, endpoint := range config.Endpoints {

		req, err := http.NewRequest(endpoint.RequestMethod, endpoint.EndpointURL, nil)
		if err != nil {
			t.Fatalf("Failed to create HTTP request: %v", err)
		}

		// Set request headers
		for key, value := range endpoint.RequestHeaders {
			req.Header.Set(key, value)
		}

		client := &http.Client{}
		start := time.Now()
		resp, err := client.Do(req)
		t.Log("Request time:", time.Since(start))
		if err != nil {
			t.Fatalf("Failed to make HTTP request: %v", err)
		}
		defer resp.Body.Close()

		for _, condition := range endpoint.ExpectedConditions {

			t.Run(
				fmt.Sprintf(
					"%s:Status-Code=%d",
					endpoint.Name,
					condition.ExpectedStatus),
				func(t *testing.T) {
					if resp.StatusCode != condition.ExpectedStatus {
						t.Fail()
					}
				})

			expected_data, err := os.ReadFile(
				filepath.Join(
					"payloads",
					condition.ExpectedResponse,
				),
			)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			t.Run(
				fmt.Sprintf(
					"%s:Response=%s",
					endpoint.Name,
					condition.ExpectedResponse,
				),
				func(t *testing.T) {
					respBody, _ := io.ReadAll(resp.Body)
					isSame := CompareJSON(string(respBody), string(expected_data))

					if !isSame {
						t.Fail()
					}
				})
		}

		fmt.Println("")
	}
}
