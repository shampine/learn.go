package main

import "fmt"
import "os"
import "io/ioutil"

func main() {
    fmt.Println("Hello, World.")

    bytes, err := ioutil.ReadAll(os.Stdin)

    if err == nil {
        fmt.Println(string(bytes))
    }
}
