// Demonstrates that using the vendored dependencies on Google Appengine doesn't
// always work as desired
package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Gophers!")
		ctx := appengine.NewContext(r)
		log.Debugf(ctx, "Because we're importing `golang.org/x/net/context`, we can't compile this")
	})
}
