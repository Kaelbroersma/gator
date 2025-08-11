-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES($1, $2, $3, $4, $5)
  RETURNING *
)
SELECT inserted_feed_follow.*, users.name AS user_name, feeds.name AS feed_name
FROM inserted_feed_follow
JOIN users ON inserted_feed_follow.user_id = users.id
JOIN feeds ON inserted_feed_follow.feed_id = feeds.id; 

-- name: GetFeedFollowsForUser :many
SELECT f.name
FROM feed_follows ff
JOIN feeds f ON ff.feed_id = f.id 
WHERE ff.user_id = $1;

-- name: UnfollowFeed :exec
DELETE FROM feed_follows
WHERE user_id = $1 AND feed_id = $2;
