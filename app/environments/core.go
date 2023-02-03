package environments

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/joho/godotenv"
)

// env vars used in app
var (
	Initialized  bool
	Environment  string
	DevMode      bool
	BaseURL      string
	ServiceName  string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	AppVersionNo string
)

// Init does env variables initialization
func Init(useProductionEnv bool) {
	log.Println("Init env: is_production:", useProductionEnv, "is_initialized:", Initialized)
	if useProductionEnv {
		if err := godotenv.Load(); err != nil {
			log.Println("Can't load .env file on the root directory.")
		}
	} else {
		if err := godotenv.Load(".env.local"); err != nil {
			log.Println("Can't load .env.local file on the root directory.")
		}
	}
	ServiceName = requireEnv("SERVICE_NAME")
	Environment = requireEnv("ENVIRONMENT")
	BaseURL = requireEnv("BASE_URL")
	DevMode = strings.ToLower(os.Getenv("DEV_MODE")) == "true"
	Initialized = true
}

func requireEnv(envName string) string {
	value, found := syscall.Getenv(envName)
	if !found {
		panic(fmt.Sprintf("%s env is required", envName))
	}
	return value
}
