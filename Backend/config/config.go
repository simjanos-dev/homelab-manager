package config

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Config struct {
	Drives []string
}

func (conf *Config) Load() error {
	filePath, err := os.Getwd()

	if err != nil {
		return errors.New("working directory cannot be retrieved")
	}

	file, err := os.Open(filePath + "/config.conf")

	if err != nil {
		return errors.New(err.Error())
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		conf.LoadConfigFileLine(fileScanner.Text())
	}

	file.Close()

	return nil
}

func (conf *Config) LoadConfigFileLine(line string) {
	if line[0] == '#' || len(strings.TrimSpace(line)) == 0 {
		return
	}

	lineData := strings.Split(line, "=")
	if len(lineData) != 2 {
		return
	}

	switch lineData[0] {
	case "drive":
		conf.Drives = append(conf.Drives, lineData[1])
	}
}
