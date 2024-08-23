
update:
	git fetch --prune && git pull; # Pega as mudanças do repositório remoto e as aplica

run:
	docker-compose up -d && \ # Inicia o docker-compose.yml
	go run src/core/main.go;  # Uma vez que o docker esteja rodando, inicia o go
