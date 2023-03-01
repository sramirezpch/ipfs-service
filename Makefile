run-docker:
	docker build -f ./infra/dockerfile -t ipfs-writer . --no-cache
	docker run --name ipfs-writer -p 8080:8080 --rm -it -d ipfs-writer