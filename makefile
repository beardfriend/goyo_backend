dev:
	docker-compose -f compose/local/docker-compose.yml --env-file .env.dev up -d
	
prod:
	docker-compose -f compose/local/docker-compose.yml --env-file .env.prod up -d