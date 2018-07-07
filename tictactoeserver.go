package main

import (
	"bufio"
	"fmt"
	"net"
)

func dibujaTablero(tab []byte) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			idx := i*3 + j
			if tab[idx] != 0 {
				fmt.Printf(" %c ", rune(tab[idx]))
			} else {
				fmt.Printf(" . ")
			}
		}
		fmt.Println("")
	}
}

func scanJugada(tab []byte, p byte) {
	var i, j byte
	valid := false
	for !valid {
		fmt.Printf("Jugada para %c: ", rune(p))
		fmt.Scanf("%d %d\n", &i, &j)
		idx := i*3 + j
		if tab[idx] == 0 {
			tab[idx] = p
			valid = true
		}
	}
}

func findWinner(tab []byte) byte {
	var winner byte
	for i := 0; i < 3; i++ {
		if (tab[i*3] == tab[i*3+1] && tab[i*3+1] == tab[i*3+2]) ||
			(tab[i] == tab[i+3] && tab[i+3] == tab[i+6]) {
			winner = tab[i*4]
			break
		}
	}
	if winner == 0 &&
		((tab[0] == tab[4] && tab[4] == tab[8]) ||
			(tab[2] == tab[4] && tab[4] == tab[6])) {
		winner = tab[4]
	}
	return winner
}

func chooseToken(tab []byte) byte {
	var token byte
	for i := 0; i < 9; i++ {
		if tab[i] != 0 {
			token = tab[i]
			break
		}
	}
	if token == 111 {
		return 120
	} else {
		return 111
	}
}

func servicio(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	tab := make([]byte, 10)
	var ganador byte
	first := true
	var p byte
	for ganador == 0 {
		r.Read(tab)
		if first {
			first = false
			p = chooseToken(tab)
		}
		fmt.Println("Jugada de cliente")
		dibujaTablero(tab)
		ganador = findWinner(tab)
		if ganador != 0 {
			fmt.Printf(" -- Ganador: %c --\n", ganador)
		} else {
			scanJugada(tab, p)
			dibujaTablero(tab)
			w.Write(tab)
			w.Flush()
			ganador = findWinner(tab)
		}
	}
	fmt.Println("-- Conexión finalizada --")
}

func main() {
	lstnr, _ := net.Listen("tcp", "localhost:8080")
	defer lstnr.Close()
	for {
		conn, _ := lstnr.Accept()
		go servicio(conn)
	}
}
