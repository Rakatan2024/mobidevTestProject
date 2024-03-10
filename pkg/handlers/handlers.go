package handlers

import (
	"mobidevtestProject/pkg/handlers/router"
	"mobidevtestProject/pkg/middleware"
	service "mobidevtestProject/pkg/service"
	"log"
	"net/http"
)

type Handler struct {
	svc   service.SvcInterface
	route *router.Router
	l     *log.Logger
}

func CreateHandler(svc service.SvcInterface, route *router.Router, l *log.Logger) Handler {
	return Handler{svc: svc, route: route, l: l}
}

func (h *Handler) HTTPHandle() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", middleware.Middleware(h.route))
	return mux
}

func (h *Handler) Routers() {
	h.route.Post("/signup", h.SignUp)
	h.route.Post("/login", h.LogIn)
	h.route.Get("/auth/confirmUserAccount", h.ConfirmAccount)
	h.route.Post("/updateProfile", h.ProfileEdit)
	h.route.Get("/", h.TempHome)
	//h.route.Post("/login", h.LogIn)

}

func (h *Handler) WriteHTTPStatusBadRequest(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}

func (h *Handler) WriteHTTPStatusUnauthorized(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(err.Error()))

}

func (h *Handler) WriteHTTPStatusInternalServerError(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
