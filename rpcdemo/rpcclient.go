package main

import (
    "log"
    "net/rpc"
    "fmt"
)

const serverAddress string = "localhost"

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

func main() {
    client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }

    args := &Args{7, 8}
    var reply int
    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}
