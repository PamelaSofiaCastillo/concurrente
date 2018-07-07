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
	winner := byte(0)
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

func main() {
	conn, _ := net.Dial("tcp", "localhost:8080")
	defer conn.Close()
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	tab := make([]byte, 10)
	ganador := byte(0)
	var c rune
	fmt.Printf("Selecciona ficha: ")
	fmt.Scanf("%c\n", &c)
	p := byte(c)
	for ganador == 0 {
		ganador = findWinner(tab)
		if ganador == 0 {
			scanJugada(tab, p)
			dibujaTablero(tab)
			w.Write(tab)
			w.Flush()
			r.Read(tab)
			fmt.Println("Jugada de servidor")
			dibujaTablero(tab)
		}
	}
	fmt.Printf(" -- Ganador: %c --\n", ganador)
	fmt.Println("-- Conexión finalizada --")
}
