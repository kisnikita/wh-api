package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
	"github.com/kisnikita/wh-api/internal/model"
	"net/http"
)

func (h *Handler) getAvailableProducts(c *gin.Context) {
	warehouseID, err := castId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	products, err := h.services.GetAvailableProducts(warehouseID)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Writer.Header().Set("Content-Type", jsonapi.MediaType)
	c.Writer.WriteHeader(http.StatusOK)

	if err = jsonapi.MarshalPayload(c.Writer, products); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) reserveProducts(c *gin.Context) {
	var reservations model.ReserveProductsRequest

	if err := jsonapi.UnmarshalPayload(c.Request.Body, &reservations); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.services.ReserveProducts(reservations)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Writer.Header().Set("Content-Type", jsonapi.MediaType)
	c.Writer.WriteHeader(http.StatusOK)

	if err = jsonapi.MarshalPayload(c.Writer, response); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) releaseProducts(c *gin.Context) {
	var releases model.ReleaseProductsRequest

	if err := jsonapi.UnmarshalPayload(c.Request.Body, &releases); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.services.ReleaseProducts(releases)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Writer.Header().Set("Content-Type", jsonapi.MediaType)
	c.Writer.WriteHeader(http.StatusOK)

	if err = jsonapi.MarshalPayload(c.Writer, response); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}
