test:
	go run main.go \
		--credentials ../../secrets/credentials.json \
		--config ../../secrets/config.json \
		--base-url http://localhost:8000/api/

prod:
	go run main.go \
		--credentials ../../secrets/credentials2.json \
		--config ../../secrets/config2.json

.PHONY: test prod
