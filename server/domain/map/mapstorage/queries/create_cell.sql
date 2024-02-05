INSERT INTO cells (x, y, user_id, building, score) 
VALUES (?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE user_id = VALUES(y)
