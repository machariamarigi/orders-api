package handler

import (
	"fmt"
	"net/http"
)

type Order struct {}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing orders")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting an order by ID")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating an order by ID")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting an order by ID")
}