package api

import (
	"net/http"

	"github.com/R-Renan/semana-tech-go-react-server/internal/store/pgstore"
	"github.com/go-chi/chi"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	r := chi.NewRouter()

	// Inicializar apiHandler com todos os campos preenchidos
	a := apiHandler{
		q: q,
		r: r,
	}

	return a
}
