package main

import "fmt"

func bridge(chanStream <-chan <-chan int) <-chan int {
    out := make(chan int)

    go func() {
        defer close(out)

        for stream := range chanStream {
            for val := range stream {
                out <- val
            }
        }
    }()

    return out
}

func main() {
    chanStream := make(chan (<-chan int))

    go func() {
        defer close(chanStream)

        for i := 1; i <= 3; i++ {
            stream := make(chan int, 2)

            stream <- i*10 + 1
            stream <- i*10 + 2

            close(stream)

            chanStream <- stream
        }
    }()

    for v := range bridge(chanStream) {
        fmt.Println(v)
    }
}