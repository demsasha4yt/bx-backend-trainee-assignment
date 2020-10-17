CREATE TABLE avito_mockapi (
    id BIGSERIAL PRIMARY KEY,
    avito_id BIGINT NOT NULL,
    price INT NOT NULL DEFAULT 100 CHECK (price > 0),
    deleted BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    changed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    count_changes INT NOT NULL DEFAULT 0
)