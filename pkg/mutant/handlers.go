package mutant

import (
	"encoding/json"
	"net/http"

	"github.com/coding-kiko/MutantCheckingApp/pkg/errors"
)

type handlers struct {
	service MutantService
}

func NewHandlers(s MutantService) Handlers {
	return &handlers{
		service: s,
	}
}

type Handlers interface {
	CheckMutationHandler(w http.ResponseWriter, r *http.Request)
	StatsHandler(w http.ResponseWriter, r *http.Request)
}

func (h *handlers) CheckMutationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		resp := errors.CreateResponse(errors.NewMethodNotAllowed())
		w.WriteHeader(resp.Error.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	var req DnaRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := errors.CreateResponse(errors.NewBadRequest())
		w.WriteHeader(resp.Error.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}
	if len(req.Dna) == 0 {
		resp := errors.CreateResponse(errors.NewBadRequest("bad request: missing dna"))
		w.WriteHeader(resp.Error.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	err = h.service.AnalyzeDna(req)
	if err != nil {
		resp := errors.CreateResponse(err)
		w.WriteHeader(resp.Error.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	json.NewEncoder(w).Encode(SuccessResponse{Data: MutantDnaResponse{Code: 200, Message: "mutation detected"}})
}

func (h *handlers) StatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		resp := errors.CreateResponse(errors.NewMethodNotAllowed())
		w.WriteHeader(resp.Error.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	stats, err := h.service.GetStats()
	if err != nil {
		resp := errors.CreateResponse(err)
		w.WriteHeader(resp.Error.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	json.NewEncoder(w).Encode(SuccessResponse{Data: stats})
}
