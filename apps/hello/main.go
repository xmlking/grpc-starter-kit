package main

import (
    "log"
    "net/http"

    "github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
    // Components routing:
    app.Route("/", &hello{})
    app.Route("/hello", &hello{})
    app.RunWhenOnBrowser()

    // HTTP routing:
    http.Handle("/", &app.Handler{
        Name:        "Hello",
        Description: "An Hello World! example",
    })

    if err := http.ListenAndServe(":8000", nil); err != nil {
        log.Fatal(err)
    }
}
