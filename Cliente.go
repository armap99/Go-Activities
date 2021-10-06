package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"time"

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
	msg := usuario
	err := gob.NewEncoder(c).Encode(msg)
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
		fmt.Println("4-Enviar mensaje de texto general ")
		fmt.Println("5-Enviar archivo general ")
		fmt.Println("6-Salir ")
		fmt.Scan(&opcliente)
		if opcliente == 1 {
			mensaje := new(claseschat.Mensaje)
			mensaje.Enviador = usuario.Nombre
			mensaje.DiaEnvio = time.Now()
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
			mensajeArchivo := new(claseschat.Mensaje)
			mensajeArchivo.Enviador = usuario.Nombre
			mensajeArchivo.DiaEnvio = time.Now()
			var destinatarioArchivo string
			var nombreArchivo string
			fmt.Println("A quien quieres enviar el mensaje: ")
			fmt.Scan(&destinatarioArchivo)
			mensajeArchivo.Destinatario = destinatarioArchivo
			fmt.Println("Nombre del archivo: ")
			nombreArchivo = obtenercadenaespacios()
			mensajeArchivo.ArchivoE.NombreArchivo = nombreArchivo
			MandarArchivo(c, mensajeArchivo)

		} else if opcliente == 3 {
			usuario.MostrarConChat()

		} else if opcliente == 4 {
			mensaje := new(claseschat.Mensaje)
			mensaje.Enviador = usuario.Nombre
			mensaje.DiaEnvio = time.Now()
			var contendio string
			mensaje.Destinatario = "Todos"
			fmt.Println("Texto: ")
			contendio = obtenercadenaespacios()
			mensaje.Contenido = contendio
			MandarMensaje(c, mensaje)

		} else if opcliente == 5 {
			mensajeArchivo := new(claseschat.Mensaje)
			mensajeArchivo.Enviador = usuario.Nombre
			mensajeArchivo.DiaEnvio = time.Now()
			var nombreArchivo string
			mensajeArchivo.Destinatario = "Todos"
			fmt.Println("Nombre del archivo: ")
			nombreArchivo = obtenercadenaespacios()
			mensajeArchivo.ArchivoE.NombreArchivo = nombreArchivo
			MandarArchivo(c, mensajeArchivo)

		} else if opcliente == 6 {
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

func MandarArchivo(conexion net.Conn, mensaje *claseschat.Mensaje) { //sacar toda la info del archivo antes de mandar
	f, err := os.Open(mensaje.ArchivoE.NombreArchivo)
	if err != nil {
		fmt.Println("No se encontro el archivo, intenta de nuevo")
		return
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	total := stat.Size()
	bs := make([]byte, total) //byts
	count, err := f.Read(bs)  //tamano
	if err != nil {
		fmt.Println(err.Error())
	}
	mensaje.ArchivoE.Longitud = count
	mensaje.ArchivoE.Bytes = bs
	err = gob.NewEncoder(conexion).Encode(mensaje)
	if err != nil {
		fmt.Println(err)
	}

}

func EsperandoMensajes(c net.Conn, usuario *claseschat.Usuario) {
	for {
		var msgs claseschat.Mensaje
		err := gob.NewDecoder(c).Decode(&msgs)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			usuario.MensajesRecibidos = append(usuario.MensajesRecibidos, msgs)
			if msgs.ArchivoE.Longitud != 0 {
				decofificarArchivo(msgs, usuario)
			}
			msgs.MostrarMensajeRecibidos()
		}

	}
}

func decofificarArchivo(mensaje claseschat.Mensaje, usuario *claseschat.Usuario) {
	if mensaje.Destinatario == "Todos" {
		f, err := os.Create("Para_" + usuario.Nombre + "-" + "De_" + mensaje.Enviador + "_" + mensaje.ArchivoE.NombreArchivo)
		if err != nil {
			fmt.Println("Algo salio mal: Defodificador de archivo")
			return
		}
		defer f.Close()

		_, err = f.Write(mensaje.ArchivoE.Bytes)
		if err != nil {
			fmt.Println("Algo salio mal", err.Error())
			return
		}
	} else {
		f, err := os.Create("Para_" + mensaje.Destinatario + "-" + "De_" + mensaje.Enviador + "_" + mensaje.ArchivoE.NombreArchivo)
		if err != nil {
			fmt.Println("Algo salio mal: Defodificador de archivo")
			return
		}
		defer f.Close()

		_, err = f.Write(mensaje.ArchivoE.Bytes)
		if err != nil {
			fmt.Println("Algo salio mal", err.Error())
			return
		}
	}

}

func main() {
	conexion := conexion()
	usuario := CrearUsuario()
	regresarUsuario(conexion, usuario)
	go EsperandoMensajes(conexion, usuario)
	MenuUsuario(conexion, usuario)

	var input string
	fmt.Scanln(&input)

}
