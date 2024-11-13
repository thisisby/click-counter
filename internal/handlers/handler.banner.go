package handlers

import (
	"click-counter/internal/services"
	"click-counter/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type ClickHandler struct {
	clickService *services.ClickService
}

func NewClickHandler(clickService *services.ClickService) *ClickHandler {
	return &ClickHandler{
		clickService: clickService,
	}
}

func (h *ClickHandler) Generate(ctx echo.Context) error {
	bannerID := ctx.Param("banner_id")
	if bannerID == "" {
		logger.ZeroLogger.Error().Msg("banner_id is required")
		return NewErrorResponse(ctx, http.StatusBadRequest, "banner_id is required")
	}

	statusCode, err := h.clickService.LogClick(bannerID)
	if err != nil {
		logger.ZeroLogger.Error().Msg(err.Error())
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Click successfully generated", nil)

}

func (h *ClickHandler) GetStats(ctx echo.Context) error {
	bannerID := ctx.Param("banner_id")
	if bannerID == "" {
		logger.ZeroLogger.Error().Msg("BannerID is required")
		return NewErrorResponse(ctx, http.StatusBadRequest, "BannerID is required")
	}

	var request struct {
		Start time.Time `json:"tsFrom"`
		End   time.Time `json:"tsTo"`
	}

	if err := ctx.Bind(&request); err != nil {
		logger.ZeroLogger.Error().Msg(err.Error())
		return NewErrorResponse(ctx, http.StatusBadRequest, "invalid request payload")
	}

	statsPayload, statusCode, err := h.clickService.GetStats(bannerID, request.Start, request.End)
	if err != nil {
		logger.ZeroLogger.Error().Msg(err.Error())
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	out := map[string]interface{}{
		"data":  statsPayload,
		"count": len(statsPayload),
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Click successfully", out)
}
