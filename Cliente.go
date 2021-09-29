package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"

	"./claseschat"
)

func obtenercadenaespacios() string {
	s := ""
	for s == "" {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s = scanner.Text()
	}
	return s
}

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
	//msg := usuario
	err := gob.NewEncoder(c).Encode(usuario)
	if err != nil {
		fmt.Println(err)
	}

}

func MenuUsuario(c net.Conn, usuario *claseschat.Usuario) {
	var opcliente int
	for {
		fmt.Println("Opciones cliente: ")
		fmt.Println("1-Enviar mensaje de texto ")
		fmt.Println("2-Enviar archivo ")
		fmt.Println("3-Mostrar mensajes recibidos ")
		fmt.Scan(&opcliente)
		if opcliente == 1 {
			mensaje := new(claseschat.Mensaje)
			mensaje.Enviador = usuario.Nombre
			var destinatario string
			var contendio string
			fmt.Println("A quien quieres enviar el mensaje: ")
			fmt.Scan(&destinatario)
			mensaje.Destinatario = destinatario
			fmt.Println("Texto: ")
			contendio = obtenercadenaespacios()
			mensaje.Contenido = contendio
			MandarMensaje(c, mensaje)

		} else if opcliente == 2 {

		} else if opcliente == 3 {
			usuario.MostarMensajesRecibidos()

		} else if opcliente == 4 {
			return
		}
	}

}

func MandarMensaje(conexion net.Conn, mensaje *claseschat.Mensaje) {
	err := gob.NewEncoder(conexion).Encode(mensaje)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	conexion := conexion()
	usuario := CrearUsuario()
	regresarUsuario(conexion, usuario)
	MenuUsuario(conexion, usuario)

	var input string
	fmt.Scanln(&input)

}
