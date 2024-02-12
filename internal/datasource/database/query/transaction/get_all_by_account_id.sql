SELECT
    id,
    name,
    account_id,
    value,
    created_at
FROM
    account_transaction
WHERE
account_id = $1;
