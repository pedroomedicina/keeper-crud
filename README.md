# Bootstrapping the project with docker

### create a network so server and postgres containers can communicate:

docker network create --driver bridge keeper-network

### run postgres database container:

docker run --name keeper-db --network keeper-network -e POSTGRES_PASSWORD=0909 -e POSTGRES_DB=keeper -p 5432:5432 -d postgres

### build server docker image:

docker build -t keeper .

### run server container:

docker run --env-file .env -p 8888:8888 --name keeper-server --network keeper-network keeper