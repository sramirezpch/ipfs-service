package config

import (
	"os"
	"reflect"
	"strings"
)

type Config struct {
	PinataApiKey    string `default:"2959abc1a254e029f3a6" env:"PINATA_API_KEY"`
	PinataSecretKey string `default:"3be0a06bc92285d6452eeb49e7f3978e3111496b46e008d36c3f24a48d4103f6" env:"PINATA_SECRET_KEY"`
	PinataJWT       string `default:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySW5mb3JtYXRpb24iOnsiaWQiOiJhZTQxMTY4ZS1mZDA4LTRjZDctYjM4Ni1jNjFlMDY2Y2QxYzAiLCJlbWFpbCI6InNyYW1pcmV6cGNoQGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJwaW5fcG9saWN5Ijp7InJlZ2lvbnMiOlt7ImlkIjoiRlJBMSIsImRlc2lyZWRSZXBsaWNhdGlvbkNvdW50IjoxfSx7ImlkIjoiTllDMSIsImRlc2lyZWRSZXBsaWNhdGlvbkNvdW50IjoxfV0sInZlcnNpb24iOjF9LCJtZmFfZW5hYmxlZCI6ZmFsc2UsInN0YXR1cyI6IkFDVElWRSJ9LCJhdXRoZW50aWNhdGlvblR5cGUiOiJzY29wZWRLZXkiLCJzY29wZWRLZXlLZXkiOiIyOTU5YWJjMWEyNTRlMDI5ZjNhNiIsInNjb3BlZEtleVNlY3JldCI6IjNiZTBhMDZiYzkyMjg1ZDY0NTJlZWI0OWU3ZjM5NzhlMzExMTQ5NmI0NmUwMDhkMzZjM2YyNGE0OGQ0MTAzZjYiLCJpYXQiOjE2Nzc4MTUyNzR9.zrhpTjaQj8nDwTxFAC4zceipucxnMcAYuG8ZP10MBY0" env:"PINATA_JWT"`
	ImageServiceUrl string `default:"localhost:8081" env:"IMAGE_SERVICE_URL"`
}

func NewConfig() *Config {
	config := &Config{}
	configValue := reflect.ValueOf(config).Elem()
	configType := reflect.TypeOf(config).Elem()

	for i := 0; i < configValue.NumField(); i++ {
		fieldValue := configValue.Field(i)
		fieldType := configType.Field(i)

		// Check if there is an environment variable with the same env tag from the field
		envTag := fieldType.Tag.Get("env")

		if envTag == "" {
			continue
		}

		envValue, exists := os.LookupEnv(strings.ToUpper(envTag))

		if exists && envValue != "" {
			// Env variable does exsist and has a value
			switch fieldValue.Kind() {
			case reflect.String:
				fieldValue.SetString(envValue)
			}
		} else {
			// Either env does not exist or does not have a value
			switch fieldValue.Kind() {
			case reflect.String:
				defaultValue := fieldType.Tag.Get("default")
				fieldValue.SetString(defaultValue)
			}
		}
	}

	return config
}
