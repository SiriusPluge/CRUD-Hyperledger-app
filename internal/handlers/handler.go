package handlers

import (
	"app-hyperledger/internal/service"
	"app-hyperledger/pkg/models"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) СreateService(ctx *gin.Context) {
	requestBody := ctx.Request.Body
	defer requestBody.Close()

	// Get body
	buf, err := io.ReadAll(requestBody)
	if err != nil {
		respError(ctx, errors.Wrap(err, "error read body"))
		return
	}

	// Unmarshal
	body := &models.ReqBody{}
	err = json.Unmarshal(buf, body)
	if err != nil {
		respError(ctx, errors.Wrap(err, "unmarshal error"))
		return
	}

	err = h.services.СreateService(body)
	if err != nil {
		respError(ctx, errors.Wrap(err, "create services error"))
		return
	}

	respOk(ctx, "Successful create services")
}

func (h *Handler) GetServiceByUUID(ctx *gin.Context) {
	requestBody := ctx.Request.Body
	defer requestBody.Close()

	// Get body
	buf, err := io.ReadAll(requestBody)
	if err != nil {
		respError(ctx, errors.Wrap(err, "error read body"))
		return
	}

	// Unmarshal
	body := &models.ReqBody{}
	err = json.Unmarshal(buf, body)
	if err != nil {
		respError(ctx, errors.Wrap(err, "unmarshal error"))
		return
	}

	service, err := h.services.GetServiceByUUID(body)
	if err != nil {
		respError(ctx, errors.Wrap(err, "create services error"))
		return
	}

	respOk(ctx, fmt.Sprintf("Get service succsessfull. Service: \n%v", service))
}

func (h *Handler) SetServiceStatus(ctx *gin.Context) {
	requestBody := ctx.Request.Body
	defer requestBody.Close()

	// Get body
	buf, err := io.ReadAll(requestBody)
	if err != nil {
		respError(ctx, errors.Wrap(err, "error read body"))
		return
	}

	// Unmarshal
	body := &models.ReqBody{}
	err = json.Unmarshal(buf, body)
	if err != nil {
		respError(ctx, errors.Wrap(err, "unmarshal error"))
		return
	}

	err = h.services.SetServiceStatus(body)
	if err != nil {
		respError(ctx, errors.Wrap(err, "create services error"))
		return
	}

	respOk(ctx, "Set service succsessfull")
}

func (h *Handler) WithDrawService(ctx *gin.Context) {
	requestBody := ctx.Request.Body
	defer requestBody.Close()

	// Get body
	buf, err := io.ReadAll(requestBody)
	if err != nil {
		respError(ctx, errors.Wrap(err, "error read body"))
		return
	}

	// Unmarshal
	body := &models.ReqBody{}
	err = json.Unmarshal(buf, body)
	if err != nil {
		respError(ctx, errors.Wrap(err, "unmarshal error"))
		return
	}

	err = h.services.WithDrawService(body)
	if err != nil {
		respError(ctx, errors.Wrap(err, "create services error"))
		return
	}

	respOk(ctx, "Withdraw service succsessfull")
}

func (h *Handler) DeleteService(ctx *gin.Context) {
	requestBody := ctx.Request.Body
	defer requestBody.Close()

	// Get body
	buf, err := io.ReadAll(requestBody)
	if err != nil {
		respError(ctx, errors.Wrap(err, "error read body"))
		return
	}

	// Unmarshal
	body := &models.ReqBody{}
	err = json.Unmarshal(buf, body)
	if err != nil {
		respError(ctx, errors.Wrap(err, "unmarshal error"))
		return
	}

	err = h.services.DeleteService(body)
	if err != nil {
		respError(ctx, errors.Wrap(err, "create services error"))
		return
	}

	respOk(ctx, "Delete service succsessfull")
}
