package gosl

import (
	"errors"
	"fmt"
	"github.com/knadh/koanf/parsers/hcl"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// ParseFileWithEnvToStruct parses the given file from path to struct *T using
// "knadh/koanf" lib with an (optional) environment variables for a secret data.
//
// You can use any of the supported file formats (JSON, YAML, TOML, or HCL). The
// structured file can be placed both locally (by system path) and accessible via
// HTTP (by URL).
//
// If err != nil, returns zero-value for a struct and error.
//
// Example:
//
//	package main
//
//	import (
//		"fmt"
//		"log"
//
//		"github.com/koddr/gosl"
//	)
//
//	type config struct {
//		ServerURL string `koanf:"server_url"`
//		AuthType  string `koanf:"auth_type"`
//		Token     string `koanf:"token"`
//	}
//
//	func main() {
//		pathToFile := "https://github.com/user/repo/config.yaml"
//		envPrefix := "MY_CONFIG" // for ex., MY_CONFIG_TOKEN=my-secret-1234567
//		modelToParse := &config{}
//
//		cfg, err := gosl.ParseFileWithEnvToStruct(pathToFile, envPrefix, modelToParse)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Println(cfg)
//	}
func ParseFileWithEnvToStruct[T any](path, envPrefix string, model *T) (*T, error) {
	// Check, if path is not empty.
	if path == "" {
		return nil, errors.New("error: given path of the structured file is empty")
	}

	// Create a new koanf instance.
	k := koanf.New(".")

	// Create a new variable with structured file extension.
	parserFormat := filepath.Ext(path)

	// Check the format of the structured file.
	switch parserFormat {
	case ".json", ".yaml", ".yml", ".toml", ".tf":
		// Create a new variable for the koanf parser.
		var parser koanf.Parser

		// Check the format of the structured file for get right koanf parser.
		switch parserFormat {
		case ".json":
			parser = json.Parser() // JSON format parser
		case ".yaml", ".yml":
			parser = yaml.Parser() // YAML format parser
		case ".toml":
			parser = toml.Parser() // TOML format parser
		case ".tf":
			parser = hcl.Parser(true) // HCL (Terraform) format parser
		}

		// Check, if path of the structured file is URL.
		u, err := url.Parse(path)
		if err != nil {
			return nil, err
		}

		// Check the schema of the given URL.
		switch u.Scheme {
		case "", "file":
			// Get the structured file from system path.
			fileInfo, err := os.Stat(path)

			// Check, if file is not dir.
			if fileInfo.IsDir() {
				return nil, fmt.Errorf("error: path of the structured file (%s) is dir", path)
			}

			// Check, if file exists.
			if err == nil || !os.IsNotExist(err) {
				// Load structured file from path (with parser of the file format).
				if err = k.Load(file.Provider(path), parser); err != nil {
					return nil, fmt.Errorf("error: structured file is not found in the given path (%s)", path)
				}
			}
		case "http", "https":
			// Get the given file from URL.
			resp, err := http.Get(path)
			if err != nil {
				return nil, fmt.Errorf("error: structured file is not found in the given URL (%s)", path)
			}
			defer resp.Body.Close()

			// Read the structured file from URL.
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, errors.New("error: raw body from the URL is not valid")
			}

			// Load structured file from URL (with parser of the file format).
			if err = k.Load(rawbytes.Provider(body), parser); err != nil {
				return nil, fmt.Errorf(
					"error: not valid structure of the %s file from the given URL (%s)",
					strings.ToUpper(strings.TrimPrefix(parserFormat, ".")), path,
				)
			}
		default:
			// If the path's schema is unknown, default action is error.
			return nil, errors.New("error: unknown path of structured file, use system path or http(s) URL")
		}
	default:
		// If the format of the structured file is unknown, default action is error.
		return nil, errors.New("error: unknown format of structured file, see: https://github.com/knadh/koanf")
	}

	// Check, if environment variables prefix was given.
	if envPrefix != "" {
		// Load environment variables.
		if err := k.Load(env.Provider(envPrefix, ".", func(s string) string {
			// Return cleared value of the environment variables.
			return strings.Replace(
				strings.ToLower(strings.TrimPrefix(s, fmt.Sprintf("%s_", envPrefix))),
				"_", ".", -1,
			)
		}), nil); err != nil {
			return nil, fmt.Errorf("error parsing environment variables, %w", err)
		}

		// Merge environment variables into the structured file data.
		if err := k.Merge(k); err != nil {
			return nil, fmt.Errorf("error merging environment variables into the structured file data, %w", err)
		}
	}

	// Unmarshal structured data to the given struct.
	if err := k.Unmarshal("", &model); err != nil {
		return nil, fmt.Errorf("error unmarshalling data from structured file to struct, %w", err)
	}

	return model, nil
}
