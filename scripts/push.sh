#!/bin/bash

echo "logging in to docker"
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 998684527758.dkr.ecr.ap-northeast-1.amazonaws.com

echo "building image"
docker build -t snek-stg-app .

echo "tagging image for ecr"
docker tag snek-stg-app:latest 998684527758.dkr.ecr.ap-northeast-1.amazonaws.com/snek-stg-app:latest

echo "pushing image to ecr"
docker push 998684527758.dkr.ecr.ap-northeast-1.amazonaws.com/snek-stg-app:latest