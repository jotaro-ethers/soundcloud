-- name: GetFollowers :many
SELECT * FROM Follow WHERE following_id = $1;

-- name: GetFollowing :many
SELECT * FROM Follow WHERE follower_id = $1;

-- name: FollowUser :exec
INSERT INTO Follow (follower_id, following_id)
VALUES ($1, $2);

-- name: UnfollowUser :exec
DELETE FROM Follow WHERE follower_id = $1 AND following_id = $2;
