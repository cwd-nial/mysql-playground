DELETE FROM playground.message_history_log WHERE id IN (
WITH cte AS (SELECT id,
                    ROW_NUMBER() OVER (PARTITION BY message_type, space_type, receiver_id, log_data) AS DuplicateCount
             FROM playground.message_history_log)
SELECT m.id FROM playground.message_history_log m INNER JOIN cte ON m.id = cte.id
WHERE cte.DuplicateCount > 1);
