package service

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Config struct {
	HexoPath string
}

func ReadConfigFile(fileName string) *Config {
	//file exists
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	var conf Config

	//read configuration
	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		result := strings.Split(s, "=")
		if result[0] == "hexoPath" {
			conf.HexoPath = result[1]
		}
	}

	return &conf
}
