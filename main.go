package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Address *Address `json:"address"`

}
//EverPerson has an address
type Address struct {
	City string `json:"city"`
	State string `json:"state"`

}
//Datastructures to open the file

func main() {
	csvFile, _ := os.Open("stats.csv")
	//creating a reader to read the csv file
	reader :=csv.NewReader(csvFile)
	var people []Person
	//iterating through the end of the file
	for{
		line, error := reader.Read()
		//Error handling
		if error == io.EOF{
			break
			//if its a real error
		}else if error != nil {
			log.Fatal(error)
		}
		//taking each column and adding it to each person struct
		people = append(people, Person{
			Firstname: line[0],
			Lastname: line[1],
			Address : &Address{
				City:line[2],
				State: line[3],
			},
		})
	}
	//adding Json
	peopleJson, _ := json.Marshal(people)
	fmt.Println(string(peopleJson))

}
