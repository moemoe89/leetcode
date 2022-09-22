SELECT DISTINCT author_id AS id
FROM Views
WHERE author_id = viewer_id
ORDER BY id

SELECT author_id AS id
FROM Views
WHERE author_id = viewer_id
GROUP BY id
ORDER BY id
