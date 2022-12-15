package repository

import (
	"cipactonal/internal/model"
	"context"
	"time"
)

func (r *Repository) RatingScores(ctx context.Context) ([]model.RatingScore, error) {
	sql := `select rc.name, r.created_at, round(rating * rc.weight * 10) as score
from ratings r
         join rating_categories rc on r.rating_category_id = rc.id
         join tickets t on r.ticket_id = t.id;`

	ratingScores := make([]model.RatingScore, 0)
	err := r.connection.SelectContext(ctx, &ratingScores, sql)
	if err != nil {
		return nil, err
	}
	return ratingScores, nil
}

func (r *Repository) CategoryScores(ctx context.Context, start time.Time, end time.Time) ([]model.CategoryScore, error) {
	dailySql := `
select rc.name, r.created_at, round(r.rating * rc.weight * 10) as score
from ratings r
         join rating_categories rc on rc.id = r.rating_category_id
where r.created_at > ?
  and r.created_at < ?
group by strftime('%Y-%m-%d', r.created_at), rc.id;
`
	weeklySql := `
select rc.name as rc_name, r.created_at, round(r.rating * rc.weight * 10) as score
from ratings r
         join rating_categories rc on rc.id = r.rating_category_id
where r.created_at > ?
  and r.created_at < ?
group by strftime('%Y-%m', r.created_at), rc.id;
`

	categoryScores := make([]model.CategoryScore, 0)

	// For periods longer than one month weekly aggregates should be returned instead of daily values.
	if end.Sub(start) > 30*24*time.Hour {
		err := r.connection.SelectContext(ctx, &categoryScores, weeklySql, start, end)
		if err != nil {
			return nil, err
		}
	} else {
		err := r.connection.SelectContext(ctx, &categoryScores, dailySql, start, end)
		if err != nil {
			return nil, err
		}
	}

	return categoryScores, nil
}

func (r *Repository) TicketScores(ctx context.Context, start time.Time, end time.Time) ([]model.TicketScore, error) {
	sql := `select t.id as ticket_id, rc.name, r.created_at, round(r.rating * rc.weight * 10) as score
from ratings r
         join rating_categories rc on rc.id = r.rating_category_id
         join tickets t on t.id = r.ticket_id
where r.created_at > ?
  and r.created_at < ?
group by t.id, strftime('%Y-%m', r.created_at), rc.id;
`

	ticketScores := make([]model.TicketScore, 0)
	err := r.connection.SelectContext(ctx, &ticketScores, sql, start, end)
	if err != nil {
		return nil, err
	}
	return ticketScores, nil
}

func (r *Repository) QualityScore(ctx context.Context, start time.Time, end time.Time) (int, error) {
	sql := `select round(sum(r.rating * rc.weight * 10) / count(r.id)) as avg_score
from ratings r
         join rating_categories rc on rc.id = r.rating_category_id
where r.created_at > ? 
  and r.created_at < ?; 
`
	var qualityScore int
	err := r.connection.GetContext(ctx, &qualityScore, sql, start, end)
	if err != nil {
		return 0, err
	}
	return qualityScore, nil
}

func (r *Repository) ScoreChange(ctx context.Context,
	firstStart time.Time, firstEnd time.Time,
	secondStart time.Time, secondEnd time.Time,
) (int, error) {
	sql := `select round((sum(r.rating * rc.weight * 10) / count(r.id)) - (sum(r2.rating * rc2.weight * 10) / count(r2.id))) as score_change
from ratings r
         join rating_categories rc on rc.id = r.rating_category_id
         join ratings r2 on r2.ticket_id = r.ticket_id
         join rating_categories rc2 on rc2.id = r2.rating_category_id
where r.created_at > ?
  and r.created_at < ?
  and r2.created_at > ?
  and r2.created_at < ?
      `

	var scoreChange int
	err := r.connection.GetContext(ctx, &scoreChange, sql, firstStart, firstEnd, secondStart, secondEnd)
	if err != nil {
		return 0, err
	}
	return scoreChange, nil
}
