package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
)

type Fruit = struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Size  string `json:"size"`
}

type FruitInput = struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Size  string `json:"size"`
}

var fruits = []Fruit{
	{Id: 1, Name: "Apple", Color: "red", Size: "medium"},
	{Id: 2, Name: "Banana", Color: "yellow", Size: "medium"},
	{Id: 3, Name: "Kiwi", Color: "brown", Size: "small"},
}

func main() {
	fmt.Println("Server is up and running on Port 3000...")
	handleRequests()
}

func handleRequests() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homePage)
	mux.HandleFunc("GET /fruits", returnFruits)
	mux.HandleFunc("POST /fruits", addFruit)
	mux.HandleFunc("GET /fruits/{id}", returnSingleFruit)

	log.Fatal(http.ListenAndServe(":3000", mux))
}

func addFruit(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var newFruit FruitInput

	err := decoder.Decode(&newFruit)
	if err != nil {
		fmt.Printf("error parsing Fruit: %v\n", err)
		return
	}

	fruitWithId := Fruit{
		Id:    slices.MaxFunc(fruits, func(a, b Fruit) int { return cmp.Compare(a.Id, b.Id) }).Id + 1,
		Name:  newFruit.Name,
		Color: newFruit.Color,
		Size:  newFruit.Size,
	}

	fruits = append(fruits, fruitWithId)

	json.NewEncoder(w).Encode(fruitWithId)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello Home Page")
}

func returnSingleFruit(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		fmt.Printf("[returnSingleFruit] Error: %v\n", err)
		return
	}
	idx := slices.IndexFunc(fruits, func(f Fruit) bool { return f.Id == id })

	if idx == -1 {
		fmt.Printf("fruit with ID %v not found", id)
		return
	}

	json.NewEncoder(w).Encode(fruits[idx])
}

func returnFruits(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(fruits)
}
