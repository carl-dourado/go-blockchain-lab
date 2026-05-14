package server

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/carl-dourado/go-blockchain-lab/internal/blockchain"
)

type API struct {
	store blockchain.Store
	mu    sync.Mutex
}

type recordRequest struct {
	Author string `json:"author"`
	Data   string `json:"data"`
}

func NewHandler(store blockchain.Store) http.Handler {
	api := &API{store: store}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", api.explorer)
	mux.HandleFunc("GET /explorer", api.explorer)
	mux.HandleFunc("GET /health", api.health)
	mux.HandleFunc("GET /chain", api.chain)
	mux.HandleFunc("GET /pending", api.pending)
	mux.HandleFunc("POST /records", api.createRecord)
	mux.HandleFunc("POST /mine", api.mine)
	mux.HandleFunc("GET /validate", api.validate)
	return mux
}

func (api *API) explorer(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(explorerHTML))
}

func (api *API) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (api *API) chain(w http.ResponseWriter, _ *http.Request) {
	api.withState(w, func(state *blockchain.ChainState) (any, int, error) {
		return state.Chain, http.StatusOK, nil
	})
}

func (api *API) pending(w http.ResponseWriter, _ *http.Request) {
	api.withState(w, func(state *blockchain.ChainState) (any, int, error) {
		return state.Pending, http.StatusOK, nil
	})
}

func (api *API) createRecord(w http.ResponseWriter, r *http.Request) {
	var body recordRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	api.withState(w, func(state *blockchain.ChainState) (any, int, error) {
		record, err := state.AddRecord(body.Author, body.Data)
		if err != nil {
			return nil, http.StatusBadRequest, err
		}

		if err := api.store.Save(state); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return record, http.StatusCreated, nil
	})
}

func (api *API) mine(w http.ResponseWriter, _ *http.Request) {
	api.withState(w, func(state *blockchain.ChainState) (any, int, error) {
		block, err := state.MinePending()
		if err != nil {
			return nil, http.StatusBadRequest, err
		}

		if err := api.store.Save(state); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return block, http.StatusCreated, nil
	})
}

func (api *API) validate(w http.ResponseWriter, _ *http.Request) {
	api.withState(w, func(state *blockchain.ChainState) (any, int, error) {
		if err := state.Validate(); err != nil {
			return map[string]any{"valid": false, "error": err.Error()}, http.StatusOK, nil
		}

		return map[string]any{"valid": true}, http.StatusOK, nil
	})
}

func (api *API) withState(w http.ResponseWriter, fn func(*blockchain.ChainState) (any, int, error)) {
	api.mu.Lock()
	defer api.mu.Unlock()

	state, err := api.store.Load()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	payload, status, err := fn(state)
	if err != nil {
		writeError(w, status, err.Error())
		return
	}

	writeJSON(w, status, payload)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
