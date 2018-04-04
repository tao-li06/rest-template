package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	"os"
)

type Configuration struct {
	DB struct {
		Name string `required:"true"`
		User string `required:"true"`
		Passowrd string `required:"true"`
		Type string `required:"true"`
		Address string `required:"true"`
		Port string `required:"true"`
	}
}

var Config = Configuration{}

func (this *Configuration) load() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	configor.Load(&Config, fmt.Sprintf("config-%s.yml", env))
}