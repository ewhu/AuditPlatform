// cmd/auditplatform/main.go
package main

import (
"flag"
"log"
"os"

"auditplatform/internal/auditplatform"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := auditplatform.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
