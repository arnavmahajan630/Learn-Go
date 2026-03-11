package main

import (
    "fmt"
    "sync"
)

var pool = sync.Pool{
    New: func() any {
        fmt.Println("Allocating new object")
        return make([]byte, 1024)
    },
}

func main() {
    b1 := pool.Get().([]byte)
    pool.Put(b1)

    b2 := pool.Get().([]byte)
    pool.Put(b2)
}