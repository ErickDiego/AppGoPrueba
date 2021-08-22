package main

//Creacion de la estructura
type Producto struct {
	Nombre string
	Tipo   string
}

//estructura para almacenar N cantidad de productos
type Productos struct {
	Productos []Producto `json: "producto`
}
