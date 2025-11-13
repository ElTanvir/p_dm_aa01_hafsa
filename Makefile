sqlc:
	docker run --rm -v "${CURDIR}:/src" -w /src sqlc/sqlc generate
templfmt:
	@templ fmt .
templgen:
	@templ generate
twc:
	tailwindcss -i internal/modules/root/css/input.css -o static/styles.css --minify --watch

.PHONY: postgres new_migration
