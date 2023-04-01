package faas

import (
	"github.com/favetelinguis/henke-larsson-ab-monorepo/handlers/rest"
	"net/http"
)

// Handler for GCP Cloud Functions
func Translate(w http.ResponseWriter, r *http.Request) {
	rest.TranslateHandler(w, r)
}
