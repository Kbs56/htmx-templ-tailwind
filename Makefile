run: 
	templ generate
	@go run cmd/main.go

## css: build tailwindcss
.PHONY: css
css:
	pnpx tailwindcss -i css/input.css -o css/output.css --watch
