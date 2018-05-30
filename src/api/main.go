package main

import (
    "sync"
    "fmt"
    "time"
    "math/rand"
    "github.com/emikohmann/accelerator-api/src/api/app"
)

func main() {
    app.StartApp()
}


func EjemploChannels() {

    var wg sync.WaitGroup

    ch := make(chan int)

    wg.Add(1)
    go cargarDato(ch, &wg)
    fmt.Println(<- ch)

    wg.Wait()
}

func cargarDato(ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    milliseconds := time.Duration(100 + rand.Intn(100))
    time.Sleep(milliseconds * time.Millisecond)

    ch <- 5
}












func EjemploGoRoutines() {
    n := 100

    var waitGroup sync.WaitGroup
    waitGroup.Add(n)

    rateLimiter := make(chan bool, 100)

    for i := 0; i < n; i++ {
        rateLimiter <- true
        go show(rateLimiter, fmt.Sprintf("Mensaje %d", i), &waitGroup)
    }

    waitGroup.Wait()
}

func show(rl chan bool, msg string, waitGroup *sync.WaitGroup) {
    defer waitGroup.Done()
    milliseconds := time.Duration(100 + rand.Intn(100))
    time.Sleep(milliseconds * time.Millisecond)
    fmt.Println(msg)

    <- rl
}