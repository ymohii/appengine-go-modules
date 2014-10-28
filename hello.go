package hello

import (
    "appengine"
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/hello-d", handlerD)
    http.HandleFunc("/hello-1", handler1)
    http.HandleFunc("/hello-2", handler2)
}

func handlerD(w http.ResponseWriter, r *http.Request) {
    context := appengine.NewContext(r)
    fmt.Fprint(w, "Hello, world! " + appengine.ModuleName(context))
}

func handler1(w http.ResponseWriter, r *http.Request) {
    context := appengine.NewContext(r)
    fmt.Fprint(w, "Hello, world! " + appengine.ModuleName(context))
}

func handler2(w http.ResponseWriter, r *http.Request) {
    context := appengine.NewContext(r)
    fmt.Fprint(w, "Hello, world! " + appengine.ModuleName(context))
}
