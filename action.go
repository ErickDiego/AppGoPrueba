package main

import (
	"fmt"
	"net/http"
)

type Producto struct {
	Nombre string
	Tipo   string
}

var listProducto = []Producto{
	{"Lechuga", "verdura"},
	{"Tomate", "verdura"}}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo desde mi servidor Web con GO")
}
