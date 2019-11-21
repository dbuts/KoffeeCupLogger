package main

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"time"
	"math/rand"
	"strconv"
	"log"
)

type Config struct{
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

	filename, _ := filepath.Abs("./tracking/presentation.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("---Checked Out Mugs Before Transaction---")
	getCheckedOut(&config)
	fmt.Println("---Checking in Cup K000002---")
	checkIn(&config, "K000002")
	fmt.Println("---Checking out New Cup---")
	var temp Rental = getRandRental()
	checkOut(&config, temp)
	getCheckedOut(&config)


	//Create backup of yaml and create updated version
	backup_name := filename[:len(filename)-4] + "_backup.yaml"
	err = os.Rename(filename, backup_name)
	if err != nil {
		log.Fatal(err)
	}
	//Marshal config back into the .yaml
	retTemp := marshalConfig(&config, filename)

	fmt.Println("\n---File after Operations---")
	fmt.Println("\n\n" + retTemp)

}

func marshalConfig (conf *Config, filename string) string{
	d, err1 := yaml.Marshal(*conf)
	if err1 !=nil {
		log.Fatal(err1)
	} else {
		os.Remove(filename)
		file, err2 := os.Create(filename)
		if err2 !=nil {
			log.Fatal(err2)
		}
		text, err5 := file.Write(d)
		if err5 != nil {
			fmt.Println(text)
			log.Fatal(err5)
		}
	}
	retString := string(d)
	return retString
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
	(*config).Rentals = append((*config).Rentals, rental)
}

func checkIn(config *Config, tag string){
	for i:=len((*config).Rentals)-1; i>=0; i--{
		if (*config).Rentals[i].Tag == tag{
			//Store rental in temporary variable
			temp := (*config).Rentals[i]
			//Copy last Rental over Rental at ith position
			length := len((*config).Rentals)
			(*config).Rentals[i] = (*config).Rentals[length-1]
			//Truncate last element
			(*config).Rentals =(*config).Rentals[:length-1]

			//Move into returned
			(*config).Returned = append((*config).Returned, temp)
		}
	}
}

func getRandRental() Rental{
	first_names := []string{"David","John","Jeff","Tom","Michael","Bill","Edward"}
	last_names  := []string{"Smith","Butler","Jackson","Walters","Arrigoni","Johnson"}
	rand.Seed(time.Now().Unix())

	var temp Rental

	temp.Tag = "K" + strconv.Itoa(rand.Intn(599999))
	temp.Drinker = first_names[rand.Intn(len(first_names))] + " " +last_names[rand.Intn(len(last_names))]
	temp.Location = "Chicago"
	currentTime := time.Now()
	temp.Date = currentTime.Format("2006.01.02 15:04:05")
	temp.Card = "0000 0000 00000"

	return temp
}
