package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Get_Config() *Data_Config {
	yamlFile, err := ioutil.ReadFile("config.yml")
	data := &Data_Config{}
	err2 := yaml.Unmarshal(yamlFile, data)

	fmt.Print(data)
	if err != nil {
		fmt.Print("have error, find or can't open the config.yml\n")
		fmt.Println(err)
	}
	if err2 != nil {
		fmt.Print("have error, find or can't open the config.yml\n")
		fmt.Println(err2)
	}
	return data

}
