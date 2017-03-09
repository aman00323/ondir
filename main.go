package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Enter []string
	Leave []string
}

func LoadConfig(path string) Config {
	var config Config
	source, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func usage() {
	fmt.Println("Usage...")
}

func ChangeDir(from string, to string) {
	from_config := filepath.Join(from, ".ondir")
	if _, err := os.Stat(from_config); err == nil {
		config := LoadConfig(from_config)
		fmt.Println(strings.Join(config.Leave, "\n"))
	}

	to_config := filepath.Join(to, ".ondir")
	if _, err := os.Stat(to_config); err == nil {
		config := LoadConfig(to_config)
		fmt.Println(strings.Join(config.Enter, "\n"))
	}
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		usage()
		os.Exit(1)
	} else {
		if args[0] != args[1] {
			ChangeDir(args[0], args[1])
		}
	}
	os.Exit(0)
}
