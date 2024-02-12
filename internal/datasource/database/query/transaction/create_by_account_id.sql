INSERT INTO account_transaction (
    name,
    account_id,
    value,
    created_at
)
VALUES (
    $1,
    $2,
    $3,
    CURRENT_TIMESTAMP
)
RETURNING id;
