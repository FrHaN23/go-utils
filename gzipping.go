package main

import (
  "net/http"
  "compress/gzip"
)

func gzipping(w http.ResponseWriter,r *http.request) error {
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		if _, err := gz.Write(byteReader); err != nil {
			return errors.New("failed to write data"))
		}
		return nil
	}
}
