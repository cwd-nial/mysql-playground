WITH cte AS (SELECT id,
                    ROW_NUMBER() OVER (PARTITION BY message_type, space_type, receiver_id, log_data) AS duplicate_count
             FROM playground.message_history_log)
SELECT m.* FROM playground.message_history_log m INNER JOIN cte ON m.id = cte.id
WHERE cte.duplicate_count > 1;
