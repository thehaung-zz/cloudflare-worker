package utils

import "github.com/go-chi/chi"

func SetGlobalPrefix(prefix string, r *chi.Mux) {
	r.Mount(prefix, r)
}
