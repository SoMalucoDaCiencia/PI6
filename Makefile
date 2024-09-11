
# Pega as mudanças do repositório remoto e as aplica
update:
	git fetch --prune && git pull;

# Inicia o docker-compose.yml
run:
	docker-compose up -d

recompile:
	clear; \
	docker rm -f /go-app-pi6; docker rmi $$(docker images | grep 'go-app-pi6'); \
	make run;
