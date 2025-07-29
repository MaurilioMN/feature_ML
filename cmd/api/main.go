package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// struct ///////////////////
type ProdutoPromoMl struct {
	Nome  string  `json:"nome"`
	Preco float32 `json:"preco"`
	Link  string  `json:"link"`
}

// funcões /////////////////
func buscarProdutos(w http.ResponseWriter, r *http.Request) {

	token := "SEU_ACCESS_TOKEN_AQUI" // Token para as consultas

	termo := r.URL.Query().Get("q")
	if termo == "" {
		http.Error(w, "Parâmetro 'q' é obrigatório", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://api.mercadolibre.com/sites/MLB/search?q=%s&sort=price_asc", termo)

	// Fazer a requisição HTTP
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Estrutura da busca esperada
	var resultado struct {
		Results []struct {
			Title     string  `json:"title"`
			Price     float32 `json:"price"`
			Permalink string  `json:"permalink"`
		} `json:"results"`
	}

	err = json.NewDecoder(resp.Body).Decode(&resultado)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusInternalServerError)
		return
	}

	// Preparar a resposta com os campos que você quer
	var produtos []ProdutoPromoMl
	for _, item := range resultado.Results {
		produtos = append(produtos, ProdutoPromoMl{
			Nome:  item.Title,
			Preco: item.Price,
			Link:  item.Permalink,
		})
	}

	// Responder com JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)
}

func main() {
	r := chi.NewRouter()

	r.Get("/buscar", buscarProdutos)

	log.Println("API rodando em http://localhost:3001")
	http.ListenAndServe(":3001", r)
}
