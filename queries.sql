-- rating score is rating * weight
select rc.weight, rating, rc.name, r.created_at, round(rating * rc.weight * 10) as score
from ratings r
         join rating_categories rc on r.rating_category_id = rc.id
         join tickets t on r.ticket_id = t.id;

-- Aggregated category scores over a period of time
-- E.g. what have the daily ticket scores been for a past week or what were the scores between 1st and 31st of January.
select rc.name as rc_name, r.created_at, round(r.rating * rc.weight * 10) as score
from ratings r
         join rating_categories rc on rc.id = r.rating_category_id
where r.created_at > '2019-03-01'
  and r.created_at < '2019-04-01'
group by strftime('%Y-%m-%d', r.created_at), rc.id;

-- For periods longer than one month weekly aggregates should be returned instead of daily values.
select rc.name as rc_name, r.created_at, round(r.rating * rc.weight * 10) as score
from ratings r
         join rating_categories rc on rc.id = r.rating_category_id
where r.created_at > '2019-03-01'
  and r.created_at < '2019-03-25'
group by strftime('%Y-%m', r.created_at), rc.id;

-- Aggregate scores for categories within defined period by ticket.
-- E.g. what were the scores for rating category for all tickets in a given period.
-- E.g. what aggregate category scores tickets have within defined rating time range have.
select t.id, rc.name as rc_name, r.created_at, round(r.rating * rc.weight * 10) as score
from ratings r
         join rating_categories rc on rc.id = r.rating_category_id
         join tickets t on t.id = r.ticket_id
where r.created_at > '2019-06-03'
  and r.created_at < '2019-06-04'
group by t.id, strftime('%Y-%m', r.created_at), rc.id;

--Overal quality score
--What is the overall aggregate score for a period.
--E.g. the overall score over past week has been 96%.
select round(sum(r.rating * rc.weight * 10) / count(r.id)) as avg_score
from ratings r
         join rating_categories rc on rc.id = r.rating_category_id
where r.created_at > '2019-04-01'
  and r.created_at < '2019-06-01';

-- Period over Period score change
-- What has been the change from selected period over previous period.
-- E.g. current week vs. previous week or December vs. January change in percentages.
select round((sum(r.rating * rc.weight * 10) / count(r.id)) - (sum(r2.rating * rc2.weight * 10) / count(r2.id))) as score_change
from ratings r
         join rating_categories rc on rc.id = r.rating_category_id
         join ratings r2 on r2.ticket_id = r.ticket_id
         join rating_categories rc2 on rc2.id = r2.rating_category_id
where r.created_at > '2019-03-01'
  and r.created_at < '2019-04-01'
  and r2.created_at > '2019-02-01'
  and r2.created_at < '2019-03-01'