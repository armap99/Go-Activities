package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"./claseschat"
)

func conexion() net.Conn {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	return c
}

func CrearUsuario() *claseschat.Usuario {
	usuario := new(claseschat.Usuario)
	var nombre string
	fmt.Println("Antes de comenzar ingresa tu Nombre: ")
	fmt.Scan(&nombre)
	usuario.Nombre = nombre
	usuario.Conectado = 1

	return usuario
}

func regresarUsuario(c net.Conn, usuario *claseschat.Usuario) {
	msg := usuario
	err := gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	conexion := conexion()
	usuario := CrearUsuario()
	regresarUsuario(conexion, usuario)

	var input string
	fmt.Scanln(&input)

}
