-- +goose Up
CREATE TABLE public.tasks (
	id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	name VARCHAR(128) NOT NULL,
	description TEXT NOT NULL,
	created_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	is_done BOOL NOT NULL DEFAULT FALSE,
	finished_time TIMESTAMPTZ
);

-- +goose Down
DROP TABLE public.tasks;
