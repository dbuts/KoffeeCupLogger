package main

import (
	"fmt"
	"bufio"
	"os"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Config struct{
	Test string `yaml:"Test"`
	Rentals []Rental `yaml:"checkedOut"`
} 


type Rental struct{
	Tag string `yaml:"tag,omitempty"`
	Drinker string `yaml:"drinker,omitempty"`
	Location string `yaml:"location,omitempty"`
	Date string `yaml:"date,omitempty"`
	Card string `yaml:"card,omitempty"`
}

func main() {
	fmt.Print("Koffee Member Name: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println("What's good, " + text + "?")

	filename, _ := filepath.Abs("/home/david/Development/KoffeeCupLogger/tracking/testprof.yaml")
	fmt.Println("Directory: " + filename)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR 1")
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("ERROR 2")
		panic(err)
	}

	fmt.Printf("\n%#v\n", config)
}
