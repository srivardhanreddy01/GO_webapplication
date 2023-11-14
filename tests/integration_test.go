package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/srivardhanreddy01/webapplication_go/api/handlers"
)

func TestHealthzHandlerIntegration(t *testing.T) {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory: %v", err)
	}

	// Construct the path to the .env file
	projectDir := filepath.Dir(currentDir)
	envFilePath := filepath.Join(projectDir, ".env")

	// Print the path for verification
	fmt.Printf("Using .env file at: %s\n", envFilePath)

	// Load environment variables from .env file
	err = godotenv.Load(envFilePath)
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(handlers.HealthzHandler))
	defer server.Close()

	// Log the server URL
	t.Logf("Test server URL: %s", server.URL)

	// Perform a GET request to the /healthz endpoint
	response, err := http.Get(server.URL + "/healthz")
	if err != nil {
		t.Fatalf("Error sending GET request: %v", err)
	}
	defer response.Body.Close()

	// Log the response status code
	t.Logf("Response status code: %d", response.StatusCode)

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}
}
