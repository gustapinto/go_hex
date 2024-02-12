INSERT INTO account (
    name,
    initial_value,
    current_value,
    created_at,
    updated_at
)
VALUES (
    $1,
    $2,
    $3,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
RETURNING id;
