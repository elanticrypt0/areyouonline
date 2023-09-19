package src

import (
	"fmt"
	"net"
)

type Server struct {
	Domain        string
	DomainAlt     string
	StatusCurrent uint8
	StatusPrev    uint8
}

func NewServer(d1, d2 string) *Server {
	return &Server{
		Domain:        d1,
		DomainAlt:     d2,
		StatusCurrent: 0,
		StatusPrev:    1,
	}
}

func (s *Server) StatusHasChanged() bool {
	return s.HasDifferentStatus()
}

func (s *Server) UpdateStatus(status uint8) {
	s.StatusPrev = s.StatusCurrent
	s.StatusCurrent = status
}

func (s *Server) HasDifferentStatus() bool {
	return s.StatusCurrent != s.StatusPrev
}

func (s *Server) PrintStatus() {
	msg := fmt.Sprintf("Estado actual: %d / Estado previo: %d", s.StatusCurrent, s.StatusPrev)
	fmt.Println(msg)
}

// TCP

func (s *Server) CheckServerStatus(n *Notify) {
	// cheque que el servidor esté funcionando por su dominio principal o alternativo
	var status uint8 = 1
	// cheque el dominio principal
	conn1, err := net.Dial("tcp", s.Domain+":http")
	if err != nil {
		status = 2
		// chequea el dominio alternativo
		conn2, err2 := net.Dial("tcp", s.DomainAlt+":http")
		if err2 != nil {
			// el servidor está caído
			status = 0
		} else {
			conn2.Close()
		}
	} else {
		conn1.Close()
	}

	// actualiza el estado
	s.UpdateStatus(status)

	// Si el estado cambia entonces notifico que cambió
	if s.StatusHasChanged() {
		fmt.Println("Estado ha cambiado")
		// guardo el log
		// Envio emails
		s.PrintCurrentStatus(n)
	}
}

func (s *Server) PrintCurrentStatus(n *Notify) {
	// imprime mensajes y notificaciones
	switch s.StatusCurrent {
	case 0:
		n.notify("El servidor está caído")
		break
	case 1:
		n.notify("El servidor funciona OK (" + s.Domain + ")")
		break
	case 2:
		n.notify("El servidor funciona por su enlace alternativo (" + s.DomainAlt + ")")
		break
	}
}
