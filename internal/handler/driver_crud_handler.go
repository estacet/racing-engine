package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/greeflas/racing-engine-backend/internal/service"
	"github.com/greeflas/racing-engine-backend/pkg/server"
)

type DriverCRUDHandler struct {
	driverService *service.DriverService
	validate      *validator.Validate
}

func NewDriverCRUDHandler(
	driverService *service.DriverService,
	validate *validator.Validate,
) *DriverCRUDHandler {
	return &DriverCRUDHandler{
		driverService: driverService,
		validate:      validate,
	}
}

func (h *DriverCRUDHandler) RegisterRoutes(mux *http.ServeMux) {
	const basePath = "/driver"
	const resourcePath = basePath + "/{id}"

	mux.HandleFunc(`POST `+basePath, h.create)
	mux.HandleFunc("GET "+basePath, h.getList)
	mux.HandleFunc(`GET `+resourcePath, h.get)
	mux.HandleFunc(`PATCH `+resourcePath, h.update)
}

func (h *DriverCRUDHandler) create(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		server.HandleError(w, err)

		return
	}
	defer r.Body.Close()

	args := new(service.CreateDriverArgs)

	if err := json.Unmarshal(bodyBytes, args); err != nil {
		server.HandleError(w, err)

		return
	}

	err = h.validate.Struct(args)
	if err != nil {
		server.HandleError(w, err)

		return
	}

	err = h.driverService.Create(r.Context(), args)
	if err != nil {
		server.HandleError(w, err)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

//nolint:dupl
func (h *DriverCRUDHandler) update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		server.HandleError(w, err)

		return
	}
	defer r.Body.Close()

	args := new(service.UpdateDriverArgs)

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

	err = h.driverService.Update(r.Context(), parsedId, args)

	if err != nil {
		server.HandleError(w, err)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *DriverCRUDHandler) getList(w http.ResponseWriter, r *http.Request) {
	driversList, err := h.driverService.GetList(r.Context())
	if err != nil {
		server.HandleError(w, err)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(driversList); err != nil {
		server.HandleError(w, err)

		return
	}
}

func (h *DriverCRUDHandler) get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		server.HandleError(w, err)

		return
	}

	driver, err := h.driverService.GetById(r.Context(), parsedId)
	if err != nil {
		server.HandleError(w, err)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(driver); err != nil {
		server.HandleError(w, err)

		return
	}
}
