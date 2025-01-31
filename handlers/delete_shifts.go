package handlers

import (
	"errors"
	"net/http"

	"github.com/Bulut-Bilisimciler/go-shift-service/models"
	"github.com/gin-gonic/gin"
)

// HandleGetShifts godoc
// @Summary get shifts by dto
// @Schemes
// @Description get shifts by dto
// @Tags shifts
// @Accept json
// @Produce json
// @Param pagination query models.Pagination true "pagination"
// @Security BearerAuth
// @Success 200 {object} handlers.RespondJson "get shifts by success"
// @Failure 400 {object} handlers.RespondJson "invalid pagination query"
// @Failure 422 {object} handlers.RespondJson "shifts not found"
// @Failure 500 {object} handlers.RespondJson "internal server error"
// @Router /shifts [get]
func (ss *ShiftService) HandleDeleteShift(c *gin.Context) (int, interface{}, error) {
	var params models.Pagination
	if err := c.ShouldBindQuery(&params); !errors.Is(err, nil) {
		return http.StatusBadRequest, nil, errors.New("invalid page or limit query for pagination")
	}

	limit := params.Limit
	offset := (params.Page - 1) * params.Limit

	// get shifts
	var shifts []models.Shift
	if err := ss.db.Where("deleted_at is NULL").Limit(limit).Offset(offset).Order("created_at DESC").Find(&shifts).Error; !errors.Is(err, nil) {
		return http.StatusNotFound, nil, errors.New("threads not found")
	}

	return http.StatusOK, shifts, nil

}
