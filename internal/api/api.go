package api

import (
	"net/http"

	store "github.com/R-Renan/semana-tech-go-react-server/internal/store/pgstore"
	"github.com/go-chi/chi"
)

type apiHandler struct {
    q *store.Queries
    r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    h.r.ServeHTTP(w, r)
}

func NewHandler(q *store.Queries) http.Handler {
    a := apiHandler{
        q: q,
    }
		r := chi.NewRouter()

		a.r = r
    return a
}
