#!/bin/bash
docker-compose -p microservice -f docker/docker-compose-migrations.yml up -d db
docker-compose -p microservice -f docker/docker-compose-migrations.yml logs
# wait for a bit so it gets provisioned
sleep 30
# run migrations
docker-compose -p microservice -f docker/docker-compose-migrations.yml run --rm migrations
# run our service
docker-compose -p microservice -f docker/docker-compose-migrations.yml up -d
sleep 2
docker-compose ps
