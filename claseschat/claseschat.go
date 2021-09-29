package claseschat

import (
	"fmt"
	"net"
	"os"
	"time"
)

/////////////////////////////////////////////////////////////////////////////////////
//Archivo
/////////////////////////////////////////////////////////////////////////////////////
type Archivo struct {
	Bytes         []byte
	Longitud      int
	NombreArchivo string
}

/////////////////////////////////////////////////////////////////////////////////////
//Mensaje
/////////////////////////////////////////////////////////////////////////////////////
type Mensaje struct {
	Enviador     string
	Destinatario string
	Contenido    string
	Archivo      Archivo
	DiaEnvio     time.Time
}

func (m *Mensaje) MostrarMensajeRecibidos() {
	fmt.Println(m.DiaEnvio, " De: ", m.Enviador, " | ", m.Contenido, m.Archivo.NombreArchivo)
}

func (m *Mensaje) MostrarMensajeEnviados() {
	fmt.Println("            ", m.DiaEnvio, " De: ", m.Enviador, " | ", m.Contenido)
}

func (m *Mensaje) MostrarMensajeNormal() {
	fmt.Println(m.DiaEnvio, " De: ", m.Enviador, " Para : ", m.Destinatario, " Contenido: ", m.Contenido, m.Archivo.NombreArchivo)
}

func (m *Mensaje) MensajeConFormato() string {
	aux := m.DiaEnvio.Format("18-May-99") + "Enviado : " + m.Enviador + " Destinatario: " + m.Destinatario + "Contenido: " + m.Contenido + m.Archivo.NombreArchivo
	return aux
}

/////////////////////////////////////////////////////////////////////////////////////
//Ususario
/////////////////////////////////////////////////////////////////////////////////////
type Usuario struct {
	Nombre            string
	Conectado         int
	Conexion          net.Conn
	MensajesRecibidos []Mensaje
}

func (u *Usuario) MostarMensajesRecibidos() {
	for _, f := range u.MensajesRecibidos {
		f.MostrarMensajeRecibidos()
	}
}

/////////////////////////////////////////////////////////////////////////////////////
//Chat
/////////////////////////////////////////////////////////////////////////////////////
type Chat struct {
	Propietarios [2]Usuario
	Mensajes     []Mensaje
}

func (c *Chat) MostrarConversacion(id string) {
	for _, f := range c.Mensajes {
		if f.Enviador == id {
			f.MostrarMensajeEnviados()
		} else {
			f.MostrarMensajeRecibidos()
		}

	}

}

/////////////////////////////////////////////////////////////////////////////////////
//Servidor
/////////////////////////////////////////////////////////////////////////////////////
type Servidor struct {
	Usuarios []Usuario
	Chats    []Chat
}

func (s *Servidor) NuevoUsuario(c Usuario) {
	s.Usuarios = append(s.Usuarios, c)
}

func (s *Servidor) Conectados() {
	fmt.Println("Conectados: ")
	for _, f := range s.Usuarios {
		if f.Conectado == 1 {
			fmt.Println(f.Nombre, ",", f.Conexion)
		}
	}
}

func (s *Servidor) MostrarTodosLosMensajes() {
	for _, f := range s.Chats {
		for _, w := range f.Mensajes {
			w.MostrarMensajeNormal()
		}
	}
}

func (s *Servidor) RespaldarMensajes() {
	mensajes := []string{}
	for _, m := range s.Chats {
		for _, w := range m.Mensajes {
			mensajes = append(mensajes, w.MensajeConFormato())
		}

	}
	file, err := os.Create("RespaldoMensajes.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for _, s := range mensajes {
		file.WriteString(s + "\n")
	}
	fmt.Println("Mensajes guardados")
}
