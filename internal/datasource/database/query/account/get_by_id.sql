SELECT
    id,
    name,
    initial_value,
    current_value,
    created_at,
    updated_at
FROM
    account
WHERE
    id = $1;
