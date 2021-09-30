package claseschat

import (
	"fmt"
	"net"
	"os"
	"sort"
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
	ArchivoE     Archivo
	DiaEnvio     time.Time
}

func (m *Mensaje) MostrarMensajeRecibidos() {
	fmt.Println(m.DiaEnvio.Format("06-Jan-02"), " De: ", m.Enviador, " | ", m.Contenido, m.ArchivoE.NombreArchivo)
}

func (m *Mensaje) MostrarMensajeEnviados() {
	fmt.Println("            ", m.DiaEnvio.Format("06-Jan-02"), " De: ", m.Enviador, " | ", m.Contenido)
}

func (m *Mensaje) MostrarMensajeNormal() {
	fmt.Println(m.DiaEnvio.Format("06-Jan-02"), " De: ", m.Enviador, " Para : ", m.Destinatario, " Contenido: ", m.Contenido, m.ArchivoE.NombreArchivo)
}

func (m *Mensaje) MensajeConFormato() string {
	aux := m.DiaEnvio.Format("06-Jan-02") + " Enviado: " + m.Enviador + " Destinatario: " + m.Destinatario + "Contenido: " + m.Contenido + m.ArchivoE.NombreArchivo
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

type sortByEnviador []Mensaje

func (a sortByEnviador) Len() int           { return len(a) }
func (a sortByEnviador) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByEnviador) Less(i, j int) bool { return a[i].Enviador < a[j].Enviador }

func (u *Usuario) MostarMensajesRecibidos() { //poner bonito
	for _, f := range u.MensajesRecibidos {
		f.MostrarMensajeRecibidos()
	}
}

func (u *Usuario) MostrarConChat() {
	auxl := u.MensajesRecibidos
	sort.Sort(sortByEnviador(auxl))
	var Enviador string
	Enviador = auxl[0].Enviador
	fmt.Println("Mensajes enviados por " + Enviador)
	for j := 0; j < len(auxl); j++ {
		if auxl[j].Enviador == Enviador {
			auxl[j].MostrarMensajeRecibidos()
		} else {
			Enviador = auxl[j].Enviador
			fmt.Println("Mensajes enviados por " + Enviador)
			auxl[j].MostrarMensajeRecibidos()
		}
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
	for _, f := range s.Usuarios {
		for _, w := range f.MensajesRecibidos {
			w.MostrarMensajeNormal()
		}
	}
}

func (s *Servidor) RespaldarMensajes() {
	mensajes := []string{}
	for _, m := range s.Usuarios {
		for _, w := range m.MensajesRecibidos {
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
