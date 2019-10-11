package main

import (
	"fmt"
//	"bufio"
//	"os"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"time"
)

type Config struct{
	Test string `yaml:"Test"`
	Rentals []Rental `yaml:"checkedOut"`
	Returned []Rental `yaml:"returned"`
}


type Rental struct{
	Tag string `yaml:"tag,omitempty"`
	Drinker string `yaml:"drinker,omitempty"`
	Location string `yaml:"location,omitempty"`
	Date string `yaml:"date,omitempty"`
	Card string `yaml:"card,omitempty"`
}

func main() {
	currentTime := time.Now()

	fmt.Println("Today's Date is: "+ currentTime.Format("2006.01.02 15:04:05"))

	filename, _ := filepath.Abs("./tracking/testprof.yaml")
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

	getCheckedOut(&config)
	getReturned(&config)
}

func getCheckedOut(config *Config){
	fmt.Println("Checked Out Containers: ")
	for i:=0; i<len((*config).Rentals); i++{
		fmt.Println((*config).Rentals[i].Tag)
	}
}

func getReturned(config *Config){
	fmt.Println("Returned Containers: ")
	for i:=0; i<len((*config).Returned); i++{
		fmt.Println((*config).Returned[i].Tag)
	}
}

func checkOut(config *Config, rental Rental){
	//TODO add Rental struct to checkedOut in yaml	
}

func checkIn(config *Config, tag string){
	for i:=len((*config).Rentals); i>=0; i--{
		if (*config).Rentals[i].Tag == tag{
			//TODO move Rental struct from checkedOut to checkedIn in the yaml	
		}
	}

}





















