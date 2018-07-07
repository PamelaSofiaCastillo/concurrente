package main
import (
	"fmt"
	"time"
)

const C = 10

var status chan int    //1-load 2-run 3-unload // este es el status del carro
var pasajeros chan int //1 abordando 2 desabordando
var end chan int
var flag bool = true

func Carro() {

	cantidadPersonas := 0
	status <- 1 //carga
	for cantidadPersonas <= C {
		select {
		case persona := <-pasajeros:
			time.Sleep(time.Second)
			fmt.Printf("La persona %d subio al carro\n", persona)
			cantidadPersonas += 1
			status <- 1

		default:
		}
	}
	status <- 2
	time.Sleep(time.Second)
	fmt.Printf("El carro esta corriendo\n")
	time.Sleep(time.Second * 5)
	fmt.Printf("El carro llego\n")
	status <- 3
	for i := 0; i < cantidadPersonas; i++ {
		pasajeros <- i
		status <- 3
	}
	end <- 1
}

func Pasajero() {
	pasajero := -1
	for flag {
		select {
		case <-end:
			time.Sleep(time.Second)
			fmt.Printf("Se bajaron todos\n")
			flag = false
		case busStatus := <-status:
			if busStatus == 1 {
				pasajero += 1
				if pasajero <= C {
					pasajeros <- pasajero
				}
			} else if busStatus == 2 {
				pasajero = 0
				fmt.Printf("El carro esta lleno\n")
			} else if busStatus == 3 {
				select {
				case pasajero := <-pasajeros:
					time.Sleep(time.Second)
					fmt.Printf("el pasajero %d esta bajando \n", pasajero)
				default:
				}
			}

		default:

		}

	}
}

func main() {
	status = make(chan int, 1)
	pasajeros = make(chan int, 1)
	end = make(chan int, 1)
	go Carro()
	go Pasajero()
	time.Sleep(time.Second * 45)
}
