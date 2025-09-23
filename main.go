package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "i am imjamam hossain. i am junior software engineer.")

}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" { //r.method = post, put, patch, delete
		http.Error(w, "please give me GET request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func main() {
	mux := http.NewServeMux() // router

	mux.HandleFunc("/hello", helloHandler) //route

	mux.HandleFunc("/about", aboutHandler) //route

	mux.HandleFunc("/products", getProducts)

	fmt.Println("Server running on: 3000")

	err := http.ListenAndServe(":3000", mux) //"Failed to start the server"

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is yellow, i love to eat orange",
		Price:       300,
		ImgUrl:      "https://cdn.britannica.com/24/174524-050-A851D3F2/Oranges.jpg",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green, i love to eat Apple",
		Price:       360,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRJJfODaTyBw4581VyPy5wQHvq4yfAIzGRHVA&s",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is green, i love to eat Banana",
		Price:       120,
		ImgUrl:      "https://5.imimg.com/data5/SELLER/Default/2024/8/446445408/ZN/UB/XS/225575346/banan-500x500.png",
	}

	prd4 := Product{
		ID:          4,
		Title:       "Mango",
		Description: "mango is green, i love to eat Mango",
		Price:       120,
		ImgUrl:      "https://www.mangovaiya.com/wp-content/uploads/2024/02/Mango.webp",
	}

	prd5 := Product{
		ID:          5,
		Title:       "Graps",
		Description: "Graps is green, i love to eat Graps",
		Price:       120,
		ImgUrl:      "https://t4.ftcdn.net/jpg/03/01/98/69/360_F_301986993_SYvMrcYECPje0HK6qRQQcm6uC7d3tpVC.jpg",
	}

	prd6 := Product{
		ID:          6,
		Title:       "Lychee",
		Description: "Lychee is white, i love to eat Lychee",
		Price:       120,
		ImgUrl:      "https://images-prod.healthline.com/hlcmsresource/images/AN_images/lychees-1296x728-feature.jpg",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
	productList = append(productList, prd6)
}

/*
   [ ] => list
	{ } => object

	JSON => javascript object notation
*/
