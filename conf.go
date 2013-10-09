package tattoo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

const CONFIG_NAME = "settings.json"

type Config struct {
	// sys config
	Port        int
	Certificate string
	Path        string
	SiteBase    string
	SiteURL     string
	// site config
	SiteTitle     string
	SiteSubTitle  string
	AuthorName    string
	TimelineCount int
	ThemeName     string
}

var config *Config
var sessionToken string

func init() {
	config = new(Config)
	// default config
	config.Port = 8888
	config.Certificate = SHA256Sum("42")
	config.SiteBase = "localhost"
	config.SiteURL = "http://localhost:8888"
	config.SiteTitle = "TATTOO!"
	config.Path = "/"
	config.AuthorName = "root"
	config.TimelineCount = 3
	config.ThemeName = "sealscript"
	sessionToken = GenerateSessionToken()
}

// Get the session token string
func GetSessionToken() string {
	return sessionToken
}

// Generate a session token
func GenerateSessionToken() string {
	sessionToken = SHA256Sum(strconv.Itoa(time.Now().Nanosecond()))
	return sessionToken
}

// Regerenerate the session token
func RevokeSessionToken() {
	sessionToken = GenerateSessionToken()
}

// Get the configuration struct
func GetConfig() *Config {
	return config
}

// Ensures that a configuration file is present by attempting to load
// it. Writes a new configuration file if it can not be loaded. Returns
// an error if it the file can not be read or written.
func (config *Config) EnsurePresent() error {
	if err := config.Load(); err != nil {
		fmt.Println("Failed to load configuration file")
		if err = config.Save(); err != nil {
			fmt.Println("Failed to save configuration file")
			return err
		}
	}
	return nil
}

// Loads the configuration by reading it from the "CONFIG_NAME" configuration
// file and unmarshaling it from JSON. Returns an error if it fails.
func (config *Config) Load() error {
	buff, err := ioutil.ReadFile(CONFIG_NAME)
	if err != nil {
		fmt.Println("Read file failed:", err)
		return err
	}
	if err := json.Unmarshal(buff, config); err != nil {
		fmt.Println("Unmarshal json failed:", err)
		return err
	}
	return nil
}

// Save the configuration by marshaling it to JSON and writing it to the
// "CONFIG_NAME" configuration file. Returns an error if it fails.
func (config *Config) Save() error {
	jsobj, err := json.Marshal(config)
	if err != nil {
		fmt.Println("Marshal json failed:", err)
		return err
	}
	if err = ioutil.WriteFile(CONFIG_NAME, jsobj, 0644); err != nil {
		fmt.Println("Write file failed:", err)
		return err
	}
	return nil
}

// Use another configuration struct
func (config *Config) Update(newcfg *Config) {
	*config = *newcfg
}

// Get a string representation of the configuration
func (config *Config) String() string {
	return fmt.Sprintf("{ Port: %v, Certificate: %v }", config.Port, config.Certificate)
}
