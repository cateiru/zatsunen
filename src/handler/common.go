package handler

import (
	"fmt"
	"net/http"
)

func (h Handler) RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello, %s", h.Config.Mode)))
}
