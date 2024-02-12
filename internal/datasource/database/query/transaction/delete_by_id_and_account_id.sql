DELETE FROM
    account_transaction
WHERE
    id = $1
    AND account_id = $2;
