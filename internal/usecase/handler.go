package usecase

import (
	"cipactonal/internal/model"
	"context"
	"time"
)

type Repository interface {
	RatingScores(ctx context.Context) ([]model.RatingScore, error)
	CategoryScores(ctx context.Context, start time.Time, end time.Time) ([]model.CategoryScore, error)
	TicketScores(ctx context.Context, start time.Time, end time.Time) ([]model.TicketScore, error)
	QualityScore(ctx context.Context, start time.Time, end time.Time) (int, error)
	ScoreChange(ctx context.Context,
		firstStart time.Time, firstEnd time.Time,
		secondStart time.Time, secondEnd time.Time,
	) (int, error)
}
type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) RatingScores(ctx context.Context) ([]model.RatingScore, error) {
	return h.repo.RatingScores(ctx)
}

func (h *Handler) CategoryScores(ctx context.Context, start time.Time, end time.Time) ([]model.CategoryScore, error) {
	return h.repo.CategoryScores(ctx, start, end)
}

func (h *Handler) TicketScores(ctx context.Context, start time.Time, end time.Time) ([]model.TicketScore, error) {
	return h.repo.TicketScores(ctx, start, end)
}

func (h *Handler) QualityScore(ctx context.Context, start time.Time, end time.Time) (int, error) {
	return h.repo.QualityScore(ctx, start, end)
}

func (h *Handler) ScoreChange(
	ctx context.Context,
	firstStart time.Time, firstEnd time.Time,
	secondStart time.Time, secondEnd time.Time,
) (int, error) {
	return h.repo.ScoreChange(ctx, firstStart, firstEnd, secondStart, secondEnd)
}
