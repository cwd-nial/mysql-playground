WITH cte AS (SELECT MAX(id) AS max_id
             FROM playground.message_history_log
             GROUP BY message_type, space_type, receiver_id, log_data
             HAVING COUNT(*) > 1)
SELECT m.*
FROM playground.message_history_log m
         INNER JOIN cte ON m.id = cte.max_id;
