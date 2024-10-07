-- name: GetSubscriptionsByAccountId :many
SELECT * FROM Subscription WHERE account_id = $1;

-- name: CreateSubscription :exec
INSERT INTO Subscription (account_id, plan_type, start_date, end_date)
VALUES ($1, $2, $3, $4);

-- name: CancelSubscription :exec
DELETE FROM Subscription WHERE subscription_id = $1;
