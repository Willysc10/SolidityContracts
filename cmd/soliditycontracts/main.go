// cmd/soliditycontracts/main.go
package main

import (
    "flag"
    "log"
    "os"

    "soliditycontracts/internal/soliditycontracts"
)

func main() {
    verbose := flag.Bool("verbose", false, "Enable verbose logging")
    input := flag.String("input", "", "Input file path")
    output := flag.String("output", "", "Output file path")
    flag.Parse()

    app := soliditycontracts.NewApp(*verbose)
    if err := app.Run(*input, *output); err != nil {
        log.Printf("Error: %v", err)
        os.Exit(1)
    }
}
