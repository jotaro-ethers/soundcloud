-- name: GetReportsBySongId :many
SELECT * FROM Report WHERE song_id = $1;

-- name: CreateReport :exec
INSERT INTO Report (account_id, song_id, reason)
VALUES ($1, $2, $3);

-- name: DeleteReport :exec
DELETE FROM Report WHERE report_id = $1;
