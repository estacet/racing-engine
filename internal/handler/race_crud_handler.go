package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/greeflas/racing-engine-backend/pkg/server"

	"github.com/google/uuid"

	"github.com/greeflas/racing-engine-backend/internal/service"
)

type RaceCRUDHandler struct {
	raceService *service.RaceService
	validate    *validator.Validate
}

func NewRaceCRUDHandler(
	raceService *service.RaceService,
	validate *validator.Validate,
) *RaceCRUDHandler {
	return &RaceCRUDHandler{
		raceService: raceService,
		validate:    validate,
	}
}

func (h *RaceCRUDHandler) RegisterRoutes(mux *http.ServeMux) {
	const basePath = "/race"
	const resourcePath = basePath + "/{id}"

	mux.HandleFunc(`GET `+resourcePath, h.get)
	mux.HandleFunc(`GET `+basePath, h.getList)
	mux.HandleFunc(`POST `+basePath, h.create)
	mux.HandleFunc(`PATCH `+resourcePath, h.update)
}

func (h *RaceCRUDHandler) get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		server.HandleError(w, err)

		return
	}

	race, err := h.raceService.GetById(r.Context(), parsedId)
	if err != nil {
		server.HandleError(w, err)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(race); err != nil {
		server.HandleError(w, err)

		return
	}
}

func (h *RaceCRUDHandler) getList(w http.ResponseWriter, r *http.Request) {
	racesList, err := h.raceService.GetList(r.Context())
	if err != nil {
		server.HandleError(w, err)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(racesList); err != nil {
		server.HandleError(w, err)

		return
	}
}

func (h *RaceCRUDHandler) create(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		server.HandleError(w, err)

		return
	}
	defer r.Body.Close()

	args := new(service.RaceArgs)

	if err := json.Unmarshal(bodyBytes, args); err != nil {
		server.HandleError(w, err)

		return
	}

	err = h.validate.Struct(args)
	if err != nil {
		server.HandleError(w, err)

		return
	}

	err = h.raceService.Create(r.Context(), args)
	if err != nil {
		server.HandleError(w, err)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

//nolint:dupl
func (h *RaceCRUDHandler) update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		server.HandleError(w, err)

		return
	}
	defer r.Body.Close()

	args := new(service.RaceArgs)

	if err := json.Unmarshal(bodyBytes, args); err != nil {
		server.HandleError(w, err)

		return
	}

	err = h.validate.Struct(args)
	if err != nil {
		server.HandleError(w, err)

		return
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		server.HandleError(w, err)

		return
	}

	err = h.raceService.Update(r.Context(), parsedId, args)

	if err != nil {
		server.HandleError(w, err)

		return
	}

	w.WriteHeader(http.StatusOK)
}
