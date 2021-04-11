package urlshort

import (
	"net/http"
)

// First step, create MapHandler() fx.
// This will map any paths (stored as keys in the map)
// to their corresponding URL (values in the map).
// If no path, then fallback http.Handler will be called.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}
