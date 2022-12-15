package router

import (
	"cipactonal/internal/model"
	proto "cipactonal/pkg/api"
	"context"
	"time"
)

type UseCase interface {
	CategoryScores(ctx context.Context, start time.Time, end time.Time) ([]model.CategoryScore, error)
	TicketScores(ctx context.Context, start time.Time, end time.Time) ([]model.TicketScore, error)
	QualityScore(ctx context.Context, start time.Time, end time.Time) (int, error)
	ScoreChange(ctx context.Context,
		firstStart time.Time, firstEnd time.Time,
		secondStart time.Time, secondEnd time.Time,
	) (int, error)
}

type Handler struct {
	proto.UnimplementedCipactonalServer
	u UseCase
}

func NewHandler(u UseCase) *Handler {
	return &Handler{u: u}
}

// CategoryScores Aggregated category scores over a period of time
func (h *Handler) CategoryScores(ctx context.Context, req *proto.CategoryScoresRequest) (*proto.CategoryScoresResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, err
	}

	scores, err := h.u.CategoryScores(ctx, req.GetTimePeriod().Start.AsTime(), req.GetTimePeriod().End.AsTime())
	if err != nil {
		return nil, err
	}

	var respScores = make([]*proto.Scores, len(scores))
	for i, score := range respScores {
		respScores[i] = &proto.Scores{
			Name:      score.Name,
			CreatedAt: score.CreatedAt,
			Score:     score.Score,
		}
	}
	return &proto.CategoryScoresResponse{
		Scores: respScores,
	}, nil
}

// TicketScores Aggregate scores for categories within defined period by ticket.
func (h *Handler) TicketScores(ctx context.Context, req *proto.TicketScoresRequest) (*proto.TicketScoresResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, err
	}

	scores, err := h.u.TicketScores(ctx, req.GetTimePeriod().GetStart().AsTime(), req.GetTimePeriod().GetStart().AsTime())
	if err != nil {
		return nil, err
	}

	var respScores = make([]*proto.Scores, len(scores))
	for i, score := range respScores {
		respScores[i] = &proto.Scores{
			Name:      score.Name,
			CreatedAt: score.CreatedAt,
			Score:     score.Score,
		}
	}
	return &proto.TicketScoresResponse{
		Scores: respScores,
	}, nil
}

// QualityScore Overal quality score
func (h *Handler) QualityScore(ctx context.Context, req *proto.QualityScoreRequest) (*proto.QualityScoreResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, err
	}

	score, err := h.u.QualityScore(ctx, req.GetTimePeriod().GetStart().AsTime(), req.GetTimePeriod().GetStart().AsTime())
	if err != nil {
		return nil, err
	}
	return &proto.QualityScoreResponse{
		Score: int32(score),
	}, nil
}

// ScoreChange Period over Period score change
func (h *Handler) ScoreChange(ctx context.Context, req *proto.ScoreChangeRequest) (*proto.ScoreChangeResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, err
	}

	score, err := h.u.ScoreChange(
		ctx, req.GetFirstTimePeriod().GetStart().AsTime(), req.GetFirstTimePeriod().GetStart().AsTime(),
		req.GetSecondTimePeriod().GetStart().AsTime(), req.GetSecondTimePeriod().GetStart().AsTime(),
	)
	if err != nil {
		return nil, err
	}
	return &proto.ScoreChangeResponse{
		Score: int32(score),
	}, nil
}
