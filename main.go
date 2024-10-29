package main

import (
	"encoding/json"
	"fmt"
	"gobuild/internal"
	"gobuild/luna"
	"io"
	"net/http"
	"strconv"
)

type Product struct {
	Name  string  `json:"title"`
	Price float64 `json:"price"`
	ID    int     `json:"id"`
}

func FetchTodos(_ ...map[string]string) map[string]interface{} {
	// fetch data from here https://fakestoreapi.com/products
	url := "https://fakestoreapi.com/products"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	c := []Product{}

	err = json.Unmarshal(body, &c)
	if err != nil {
		panic(err)
	}

	//

	return map[string]interface{}{
		"products": c,
	}
}
func FetchProduct(id int) map[string]interface{} {
	url := fmt.Sprintf("https://fakestoreapi.com/products/%d", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	c := Product{}

	err = json.Unmarshal(body, &c)
	if err != nil {
		panic(err)
	}

	return map[string]interface{}{
		"product": c,
	}
}

func WrappedFetchTodos(params ...map[string]string) map[string]interface{} {
	if len(params) > 0 {
		fmt.Println("params", len(params))
		// Get the first map from params
		paramMap := params[0]

		// Try to get the "id" value from the map
		if id, ok := paramMap["id"]; ok {
			// convert id to int
			//
			num, err := strconv.Atoi(id)
			if err != nil {
				panic(err)
			}
			return FetchProduct(num)
		}
	}

	// No params or "id" not found, return empty products
	return map[string]interface{}{
		"products": []Product{},
	}
}
func returnName(_ ...map[string]string) map[string]interface{} {
	return map[string]interface{}{
		"name": "Luna",
	}
}

func main() {

	engine, err := luna.New(
		luna.Config{
			ENV:         "development",
			FrontendDir: "frontend",
			TailwindCSS: true,
			Routes: []internal.ReactRoute{
				{
					Path:  "/",
					Props: returnName,
				},
				{
					Path:  "/about",
					Props: FetchTodos,
				},
				{
					Path:  "/about/:id",
					Props: WrappedFetchTodos,
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}

	engine.FrontEnd()
	engine.Server.Start(":8080")
}
