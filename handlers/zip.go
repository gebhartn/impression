// Package handlers contains a collection of route handlers and middlewares
// to handle the uploading and compression of file (read: image) uploads.
package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

const (
	acceptEncoding  = "Accept-Encoding"
	contentEncoding = "Content-Encoding"
)

// GzipHandler is just a method receiver for the actual handler implementation
type GzipHandler struct{}

// GzipMiddleware creates a gzipped response if the headers accept gzip encoding
func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get(acceptEncoding), "gzip") {
			wrw := NewWrappedResponseWriter(rw)
			wrw.Header().Set(contentEncoding, "gzip")

			next.ServeHTTP(wrw, r)
			defer wrw.Flush()

			return
		}

		next.ServeHTTP(rw, r)
	})
}

// WrappedResponseWriter wraps a normal RW with a gzip writer
type WrappedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

// NewWrappedResponseWriter constructs a new instance of a WrappedResponseWriter
func NewWrappedResponseWriter(rw http.ResponseWriter) *WrappedResponseWriter {
	gw := gzip.NewWriter(rw)

	return &WrappedResponseWriter{rw: rw, gw: gw}
}

// Header implements the Header method for our writer
func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.rw.Header()
}

// Write implements the Write method for our writer
func (wr *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

// WriteHeader implements the WriteHeader method for our writer
func (wr *WrappedResponseWriter) WriteHeader(status int) {
	wr.rw.WriteHeader(status)
}

// Flush implements the Flush method for our writer
func (wr *WrappedResponseWriter) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}
