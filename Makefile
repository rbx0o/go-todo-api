include .env
export

migrations-up-all:
	goose up

migrations-up-one:
	goose up-by-one

migrations-down-all:
	goose down