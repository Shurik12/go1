package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go1/src/menu"
	"go1/src/wildberies"

	"github.com/go-yaml/yaml"
)

// Config struct for webapp config
type Config struct {
	Awesome struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		Host string `yaml:"host"`

		// Port is the local machine TCP Port to bind the HTTP Server to
		Port string `yaml:"port"`

		// Authentication token to access wildberies api
		Auth_token string `yaml:"auth_token"`

		Timeout struct {
			// Server is the general server timeout to use
			// for graceful shutdowns
			Server time.Duration `yaml:"server"`

			// Write is the amount of time to wait until an HTTP server
			// write opperation is cancelled
			Write time.Duration `yaml:"write"`

			// Read is the amount of time to wait until an HTTP server
			// read operation is cancelled
			Read time.Duration `yaml:"read"`

			// Read is the amount of time to wait
			// until an IDLE HTTP session is closed
			Idle time.Duration `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"awesome"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	// Return the configuration path
	return configPath, nil
}

func f1(args ...string) error {
	println(1)
	return nil
}

func f2(args ...string) error {
	println(2)
	return nil
}

type (
	// Doer is an interface that can do http request
	Doer interface {
		Do(req *http.Request) (*http.Response, error)
	}
	// A Client manages communication with the Wildberies API.

)

func main() {
	fmt.Println("Hello, World!")

	cfgPath, err := ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	// Request to Wildberies=========================================================================
	key := cfg.Awesome.Auth_token
	client := wildberies.NewClient(
		wildberies.AccessToken(key),
	)
	// client.Income().SellerInfo(context.Background())
	incomes, _, _ := client.Supplier().Incomes(context.Background())
	fmt.Println(incomes)
	// ==============================================================================================

	commands := []menu.CommandOption{
		{Command: "1", Description: "Print playlists", Function: f1},
		{Command: "2", Description: "Runs command2", Function: f2},
	}

	menuOptions := menu.NewMenuOptions("'menu' for help > ", 0, "menu")

	menu := menu.NewMenu(commands, menuOptions)
	menu.Start()
}
