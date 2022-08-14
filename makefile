DEV:
docker-compose -f compose/local/docker-compose.yml --env-file .env.dev up -d

PROD:
docker-compose -f compose/local/docker-compose.yml --env-file .env.prod up -d