SELECT
    id,
    name,
    account_id,
    value,
    created_at
FROM
    account_transaction
WHERE
    id = $1
    AND account_id = $2;
