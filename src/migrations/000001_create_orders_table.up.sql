CREATE TABLE IF NOT EXISTS orders (
    id bigserial PRIMARY KEY,
    uid CHAR(19) UNIQUE NOT NULL,
    data jsonb NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
