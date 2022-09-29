package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Enter []string
	Leave []string
}

const ondir = ".ondir"

func LoadConfig(path string) Config {
	var config Config
	source, err := ioutil.ReadFile(filepath.Join(filepath.Clean(path)))
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

func FindOndirConfig(path string) (string, error) {
	config := filepath.Join(path, ondir)
	if _, err := os.Stat(config); err == nil {
		return config, nil
	} else if path == "/" {
		return "", errors.New("not found")
	}
	if path, err := filepath.Abs(filepath.Join(path, "..")); err == nil {
		return FindOndirConfig(path)
	} else {
		return "", errors.New("not found")
	}
}

func ChangeDir(from string, to string) {
	leave, leave_err := FindOndirConfig(from)
	enter, enter_err := FindOndirConfig(to)
	if leave != enter {
		if leave_err == nil {
			fmt.Println("# LEAVE :", leave)
			config := LoadConfig(leave)
			fmt.Println(strings.Join(config.Leave, "\n"))
		}
		if enter_err == nil {
			fmt.Println("# ENTER :", enter)
			config := LoadConfig(enter)
			fmt.Println(strings.Join(config.Enter, "\n"))
		}
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
