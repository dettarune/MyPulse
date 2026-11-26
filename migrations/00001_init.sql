-- +goose Up
CREATE TABLE employees (
    id            BIGSERIAL PRIMARY KEY,
    name          TEXT NOT NULL,
    username      TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role          TEXT NOT NULL CHECK (role IN ('sales', 'supervisor', 'admin')),
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE attendances (
    id            BIGSERIAL PRIMARY KEY,
    employee_id   BIGINT NOT NULL REFERENCES employees(id),
    checked_in_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    latitude      DOUBLE PRECISION NOT NULL,
    longitude     DOUBLE PRECISION NOT NULL
);

CREATE INDEX idx_attendances_employee_time
    ON attendances (employee_id, checked_in_at DESC);

-- +goose Down
DROP TABLE attendances;
DROP TABLE employees;
