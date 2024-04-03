SET timezone = 'America/Sao_Paulo';
CREATE SCHEMA IF NOT EXISTS pixel_pay;

CREATE TABLE pixel_pay.users
(
    id              serial primary key,
    name            varchar(150) not null,
    document        varchar(50) unique,
    documentType    varchar(20) not null,
    email           varchar(50) unique,
    -- in cents
    balance         pg_catalog.float4 default 0,
    password        varchar(50) not null
);

CREATE TABLE pixel_pay.transactions
(
    id          SERIAL PRIMARY KEY,
    payeer_id   INTEGER     NOT NULL,
    payee_id    INTEGER     NOT NULL,
    value       INTEGER     NOT NULL,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON pixel_pay.transactions (payeer_id);
CREATE INDEX ON pixel_pay.transactions (payee_id);

CREATE OR REPLACE FUNCTION pixel_pay.transfer(
    payeer_id_tx INTEGER,
    payee_id_tx INTEGER,
    value_tx INTEGER
) RETURNS TABLE (status BOOL, message TEXT) AS $$
DECLARE
    payeer_balance INTEGER;
BEGIN
    PERFORM pg_advisory_xact_lock(payeer_id_tx);
    SELECT balance INTO payeer_balance FROM pixel_pay.users WHERE id = payeer_id_tx;
    IF payeer_balance < value_tx THEN
        return query SELECT FALSE AS status, 'insufficient balance' AS message;
        RETURN;
    END IF;

    UPDATE pixel_pay.users SET balance = balance - value_tx WHERE id = payeer_id_tx;
    UPDATE pixel_pay.users SET balance = balance + value_tx WHERE id = payee_id_tx;

    INSERT INTO pixel_pay.transactions (payeer_id, payee_id, value)
    VALUES (payeer_id_tx, payee_id_tx, value_tx);

    return query SELECT true AS status, 'success' AS message;
END;
$$ LANGUAGE plpgsql;

DO
$do$
    DECLARE
        counter INTEGER := 1;
    BEGIN
        FOR counter IN 1..50 LOOP
                INSERT INTO pixel_pay.users (name, document, documentType, email, password, balance)
                VALUES (
                           'User ' || counter,
                           '00000000' || counter,
                           CASE
                               WHEN counter % 2 = 0 THEN 'cnpj'
                               ELSE 'cpf'
                           END,
                           'user' || counter || '@example.com',
                           'Password' || counter,
                           counter * 100

                       );
            END LOOP;
    END
$do$;