package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

// Methods eligible for overriding
var eligible = map[string]struct{}{
	http.MethodDelete: {},
	http.MethodPut:    {},
	http.MethodPatch:  {},
}

const formType = "application/x-www-form-urlencoded"

// New allows HTML <form method="post">'s to dispatch PATCH, PUT and
// DELETE requests by overriding the request method using a hidden "_method"
// field in the form body.
func MethodOverride() *methodOverride {
	return &methodOverride{}
}

type methodOverride struct {
}

var _ Middleware = &methodOverride{}

func (m *methodOverride) Middleware(next http.Handler) http.Handler {
	fmt.Println("HELLO....")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only override POST requests
		if r.Method != http.MethodPost {
			next.ServeHTTP(w, r)
			return
		}
		// Must have a request body and set the content-type to
		// application/x-www-form-urlencoded.
		if r.Body == nil || r.Header.Get("Content-Type") != formType {
			next.ServeHTTP(w, r)
			return
		}
		// Try parsing the request form
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Check if the _method form value is set
		override := strings.ToUpper(r.Form.Get("_method"))
		// Ensure the method is eligible for overriding
		if _, ok := eligible[override]; !ok {
			next.ServeHTTP(w, r)
			return
		}
		// Override the request method
		r.Method = override
		next.ServeHTTP(w, r)
	})
}
