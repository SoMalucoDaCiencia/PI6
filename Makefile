
# Pega as mudanças do repositório remoto e as aplica
update:
	git fetch --prune && git pull;

# Inicia o docker-compose.yml
# Uma vez que o docker esteja rodando, inicia o go
run:
	docker-compose up -d && \
	go run src/core/main.go;
