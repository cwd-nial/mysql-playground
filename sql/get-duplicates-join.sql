SELECT DISTINCTROW m1.*
FROM playground.message_history_log m1
         INNER JOIN playground.message_history_log m2
WHERE m1.id > m2.id
  AND m1.message_type = m2.message_type
  AND m1.space_type = m2.space_type
  AND m1.receiver_id = m2.receiver_id
  AND ((m1.log_data IS NULL AND m2.log_data IS NULL) OR m1.log_data = m2.log_data);
