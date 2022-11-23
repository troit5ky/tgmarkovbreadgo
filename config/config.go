package config

import (
	"bytes"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type TOML struct {
	Token         string
	Cooldown      int64
	DatabaseLimit int64
}

var Config TOML

func Init() error {
	data, err := os.ReadFile("config.toml")
	if err != nil {
		createExample()
		log.Fatal(err)

		return err
	}

	_, err = toml.Decode(string(data), &Config)
	if err != nil {
		createExample()
		log.Fatal(err)

		return err
	}

	return nil
}

func createExample() {
	defer log.Println("Config initialized")

	data := &bytes.Buffer{}
	toml.NewEncoder(data).Encode(TOML{
		Token:         "Bot token",
		Cooldown:      7,
		DatabaseLimit: 7000,
	})

	bytes := data.Bytes()
	os.WriteFile("example.toml", bytes, 0644)

	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		os.WriteFile("config.toml", bytes, 0644)
	}

	log.Println("Please, check example.toml")
}
