package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"./claseschat"
)

func servidor(lservidor *claseschat.Servidor) {
	servidor, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		cliente, err := servidor.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//msg := lprocesos.Primero()
		//err = gob.NewEncoder(cliente).Encode(msg)//mando al cliente

		go handleClient(cliente, lservidor)
	}
}

func handleClient(c net.Conn, servidor *claseschat.Servidor) { // se agrega al usuario al servidor
	var msg claseschat.Usuario
	err := gob.NewDecoder(c).Decode(&msg)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		msg.Conexion = c
		servidor.Usuarios = append(servidor.Usuarios, msg)
	}
	for {
		mensaje := new(claseschat.Mensaje)
		err := gob.NewDecoder(c).Decode(&mensaje)
		if err != nil {
			fmt.Println(err)
			return
		} else {

			fmt.Println("Mensaje enviado: [ ", mensaje.Enviador, " | ", mensaje.Destinatario, " ] ")
			if mensaje.ArchivoE.Longitud != 0 {
				fmt.Println("Archivo: ", mensaje.ArchivoE.NombreArchivo)
			}
			EnviarMensaje(mensaje, servidor)
		}
	}

}

func EnviarMensaje(msg *claseschat.Mensaje, servidor *claseschat.Servidor) {
	var aux net.Conn
	for _, f := range servidor.Usuarios {
		if f.Nombre == msg.Destinatario {
			aux = f.Conexion
		}
	}
	err := gob.NewEncoder(aux).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	lservidor := new(claseschat.Servidor)
	go servidor(lservidor) //inicio de serviodr
	var op int
	Detener := false
	for !Detener {
		fmt.Println("Opciones de servidor: ")
		fmt.Println("1-Mostrar todos los mensajes ")
		fmt.Println("2-Respaldar todos los mensajes ")
		fmt.Println("3-Terminar servidor ")
		fmt.Scan(&op)
		if op == 1 {
			lservidor.MostrarTodosLosMensajes()
		} else if op == 2 {
			lservidor.RespaldarMensajes()
		} else if op == 3 {
			Detener = true
		}
	}

}
