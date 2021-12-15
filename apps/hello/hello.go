package main

import (
    "fmt"

    "github.com/maxence-charriere/go-app/v9/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
    app.Compo
    name string
}

func (f *hello) OnPreRender(ctx app.Context) {
    fmt.Println("component prerendered")
}

func (f *hello) OnMount(ctx app.Context) {
    fmt.Println("component mounted")
}

func (f *hello) OnNav(ctx app.Context) {
    fmt.Println("component navigated:")
}

func (f *hello) OnDismount() {
    fmt.Println("component dismounted")
}

// The Render method is where the component appearance is defined. Here, a
func (h *hello) Render() app.UI {
    return app.Div().Body(
        app.H1().Body(
            app.Text("Hello, "),
            app.If(h.name != "",
                app.Text(h.name),
            ).Else(
                app.Text("World!"),
            ),
        ),
        app.P().Body(
            app.Input().
                Type("text").
                Value(h.name).
                Placeholder("What is your name?").
                AutoFocus(true).
                OnChange(h.ValueTo(&h.name)),
        ),
    )
}
