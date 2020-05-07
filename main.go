package main

/*
#cgo darwin LDFLAGS: -L/usr/local/lib/ -lexample
#cgo darwin LDFLAGS: -L/usr/local/lib/ -lexample
#cgo linux LDFLAGS: -L/usr/local/lib/libexample.so -lexample
#include <example.h>
#include <stdlib.h>
*/
import "C"
import (
    "fmt"
    "unsafe"
)

func main() {
    for {
        t := C.example_create_thing()

        buf := encode(t)

        decode(t, buf)

        C.example_destroy_thing(t)
    }
}

func encode(t unsafe.Pointer) []byte {
    buf := make([]byte, 718)

    sz := C.example_encode(t, unsafe.Pointer(&buf[0]), C.int(len(buf)))
    if sz == -1 {
        panic("provided buffer was not large enough")
    }

    return buf[:sz]
}

func decode(t unsafe.Pointer, buf []byte) {
    fmt.Println("decoding")
    fmt.Println(string(buf))

    C.example_decode(t, unsafe.Pointer(&buf[0]), C.int(len(buf)))
}
