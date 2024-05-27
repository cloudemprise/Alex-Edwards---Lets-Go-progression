package main

/*
Alex Edwards : Let's Go

Chapter 6 Middleware

Section 6.2 Setting Security Headers
*/

import (
	"net/http"
)

/*
Flow of Control:
secureHeaders → servemux → application handler → servemux → secureHeaders
*/

// secureHeaders automatically adds HTTP security headers. (wraps servermux)
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")
		// Any code here will execute on the way down the chain.
		next.ServeHTTP(w, r)
		// Any code here will execute on the way back up the chain.
	})
}
