-- name: GetFollowers :many
SELECT * FROM Follow WHERE following_id = $1;

-- name: GetFollowing :many
SELECT * FROM Follow WHERE follower_id = $1;

-- name: FollowAccount :exec
INSERT INTO Follow (follower_id, following_id)
VALUES ($1, $2);

-- name: UnfollowAccount :exec
DELETE FROM Follow WHERE follower_id = $1 AND following_id = $2;
