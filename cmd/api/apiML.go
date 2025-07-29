package api

import (
	"endcoding/json"
	"log"
	"net/http"

	
	"github.com/go-chi/chi/v5"
)
type ProdutoPromoMl struct{
	Nome string `json:"nome"`
	Produto float64 `json:"preço"`
}

func main(){

	r := chi.NemRoute()

	r.Get("/Produtos", func (w http.ResponseWriter, r *http.Request) {

		produtos := []ProdutoPromoMl{
			{}
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(produtos)
	})
	log.Println("Usando porta: 3001")
	http.ListenAndServe(":3001", r)
}