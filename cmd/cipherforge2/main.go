// cmd/cipherforge2/main.go
package main

import (
"flag"
"log"
"os"

"cipherforge2/internal/cipherforge2"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := cipherforge2.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
