UPDATE
    account
SET
    name = $1,
    current_value = $2
WHERE
    id = $3;
