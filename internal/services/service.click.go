package services

import (
	"click-counter/internal/errs"
	"click-counter/internal/models"
	"click-counter/internal/repositories"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type ClickService struct {
	clickRepository repositories.ClickRepository
}

func NewClickService(clickRepository repositories.ClickRepository) *ClickService {
	return &ClickService{
		clickRepository: clickRepository,
	}
}

func (s *ClickService) LogClick(bannerID string) (int, error) {
	err := s.clickRepository.Save(bannerID)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("ClickService.LogClick - s.clickRepository.Save: %w", err)
	}

	return http.StatusCreated, nil
}

func (s *ClickService) GetStats(bannerID string, startTimestamp, endTimestamp time.Time) ([]models.Click, int, error) {
	clicks, err := s.clickRepository.FindByIdInRange(bannerID, startTimestamp, endTimestamp)
	if err != nil {
		if errors.Is(err, errs.ErrBannerNotFound) {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("ClickService.GetStats - s.clickRepository.FindByIdInRange: %w", err)
	}

	return clicks, http.StatusOK, nil
}
