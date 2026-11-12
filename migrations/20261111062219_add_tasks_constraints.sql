-- +goose Up
ALTER TABLE public.tasks
	ADD CONSTRAINT tasks_name_not_blank_check
		CHECK (btrim(name) <> ''),
	
	ADD CONSTRAINT tasks_completion_check
		CHECK (is_done = (finished_time IS NOT NULL));

-- +goose Down
ALTER TABLE public.tasks
	DROP CONSTRAINT tasks_name_not_blank_check,
	DROP CONSTRAINT tasks_completion_check;
