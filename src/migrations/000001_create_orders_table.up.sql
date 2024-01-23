CREATE TABLE IF NOT EXISTS orders (
    id bigserial PRIMARY KEY,
    data jsonb NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
