package main

import (
    "fmt"
    "math/rand"
    "time"
)

const (
    DIST = 100
)

func tortoise(end chan<- string, bite chan<- bool, chHare <-chan int) {
    pos := 0
    lastHarePos := 0
    posHare := 0
    for pos < DIST {
        pos++
        select {
        case posHare = <- chHare:
        default:
            posHare = lastHarePos
        }
        if posHare == pos {
            select {
            case bite <- true:
                fmt.Println("I'm the migthy tortoise *CHOMP CHOMP*!")
            default:
            }
        }
        lastHarePos = posHare
        time.Sleep(time.Millisecond * 10)
        fmt.Printf("Migthy tortoise at %d\n", pos)
    }
    select {
    case end <- "I'm the migthy tortoise and I won!":
    default:
        fmt.Println("I'm the migthy tortoise anyways!")
    }
}

func hare(end chan<- string, bite <-chan bool, chHare chan<- int) {
    pos := 0
    duration := time.Duration(0)
    action := ""
    for pos < DIST {
        if rand.Intn(100) > 50{
            action = "sleeping"
            duration = 50
        } else {
            action = "is"
            duration = 10
        }
        fmt.Printf("Hare %s at %d\n", action, pos)
        select {
        case <-bite:
            pos -= 3
            fmt.Printf("Hare: AW SHHHHUT!, back to %d\n", pos)
            action = "is"
        case <-time.After(time.Millisecond * duration):
            pos += 5
            if pos >= DIST {
                break
            }
            chHare <- pos
        }

    }
    select {
    case end <- "I'm the Hare, I won!":
    default:
        fmt.Println("Hare: :(")
    }
}

func main() {
    end := make(chan string)
    bite := make(chan bool)
    chHare := make(chan int)
    go tortoise(end, bite, chHare)
    go hare(end, bite, chHare)
    fmt.Println(<-end)
}