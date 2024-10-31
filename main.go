package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	obsws "github.com/christopher-dG/go-obs-websocket"
)

// Config holds the application configuration
type Config struct {
	host    string
	port    int
	scene1  string
	scene2  string
	timeout time.Duration
	verbose bool
}

var osExit = os.Exit
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	// Initialize logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Parse configuration
	config := parseFlags()

	if config.verbose {
		log.Printf("Connecting to OBS at %s:%d\n", config.host, config.port)
		log.Printf("Switching between scenes: %s and %s\n", config.scene1, config.scene2)
	}

	// Create and configure client
	client := createClient(config)
	defer client.Disconnect()

	// Connect with timeout
	if err := connectWithTimeout(&client, config.timeout); err != nil {
		log.Fatalf("Failed to connect to OBS: %v", err)
	}

	if err := switchScene(client, config); err != nil {
		log.Fatalf("Failed to switch scene: %v", err)
	}
}

// parseFlags handles command-line argument parsing and validation
func parseFlags() Config {
	config := Config{}

	// Add version flag
	showVersion := flag.Bool("version", false, "Show version information")

	flag.StringVar(&config.host, "host", "localhost", "OBS WebSocket host")
	flag.IntVar(&config.port, "port", 4444, "OBS WebSocket port")
	flag.DurationVar(&config.timeout, "timeout", 5*time.Second, "Connection timeout")
	flag.BoolVar(&config.verbose, "verbose", false, "Enable verbose logging")

	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <scene1> <scene2>\n\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	// Check if version flag was set
	if *showVersion {
		fmt.Printf("obs_switchscene %s (commit: %s, built at: %s)\n", version, commit, date)
		osExit(0)
	}
	// Validate positional arguments
	args := flag.Args()
	if len(args) != 2 {
		flag.Usage()
		osExit(1)
	}

	config.scene1 = args[0]
	config.scene2 = args[1]

	return config
}

// createClient creates and configures an OBS WebSocket client
func createClient(config Config) obsws.Client {
	return obsws.Client{
		Host: config.host,
		Port: config.port,
	}
}

// connectWithTimeout attempts to connect to OBS with a timeout
func connectWithTimeout(client *obsws.Client, timeout time.Duration) error {
	connectionChan := make(chan error, 1)

	go func() {
		connectionChan <- client.Connect()
	}()

	select {
	case <-time.After(timeout):
		return fmt.Errorf("connection timeout after %v", timeout)
	case err := <-connectionChan:
		return err
	}
}

// switchScene handles the scene switching logic
func switchScene(client obsws.Client, config Config) error {
	// Get current scene
	sceneList, err := obsws.NewGetSceneListRequest().SendReceive(client)
	if err != nil {
		return fmt.Errorf("failed to get scene list: %w", err)
	}

	// Validate that both scenes exist
	scenes := make(map[string]bool)
	for _, scene := range sceneList.Scenes {
		scenes[scene.Name] = true
	}

	if !scenes[config.scene1] || !scenes[config.scene2] {
		return fmt.Errorf("one or both scenes do not exist: %s, %s", config.scene1, config.scene2)
	}

	// Determine which scene to switch to
	targetScene := config.scene1
	if sceneList.CurrentScene == config.scene1 {
		targetScene = config.scene2
	}

	// Switch scene
	if config.verbose {
		log.Printf("Switching to scene: %s\n", targetScene)
	}

	_, err = obsws.NewSetCurrentSceneRequest(targetScene).SendReceive(client)
	if err != nil {
		return fmt.Errorf("failed to switch scene: %w", err)
	}

	return nil
}
