package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var respExcludeHeader = map[string]bool{
	"Content-Length":    true,
	"Transfer-Encoding": true,
	"Trailer":           true,
}

func DumpResponse(r *http.Response) ([]byte, error) {
	var w bytes.Buffer

	// Status line
	text := r.Status
	if text == "" {
		text = http.StatusText(r.StatusCode)
	} else {
		// Just to reduce stutter, if user set r.Status to "200 OK" and StatusCode to 200.
		// Not important.
		text = strings.TrimPrefix(text, strconv.Itoa(r.StatusCode)+" ")
	}

	if _, err := fmt.Fprintf(&w, "HTTP/%d.%d %03d %s\r\n", r.ProtoMajor, r.ProtoMinor, r.StatusCode, text); err != nil {
		return nil, err
	}

	// Rest of header
	err := r.Header.WriteSubset(&w, respExcludeHeader)
	if err != nil {
		return nil, err
	}

	// End-of-header
	if _, err := io.WriteString(&w, "\r\n"); err != nil {
		return nil, err
	}

	// Success
	return w.Bytes(), nil
}
