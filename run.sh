#!/bin/bash

IMAGE_NAME="toxic-rest-api"
IMAGE_TAG="latest"
DOCKERFILE="./Dockerfile"
CONTAINER_NAME="toxic-rest-api"

./mvnw clean package

if [ "$(docker ps -a -q -f name=$CONTAINER_NAME)" ]; then
  echo "Stop and remove existing Toxic REST API Docker container..."
  docker stop $CONTAINER_NAME
  docker rm $CONTAINER_NAME
fi

if [ "$(docker images -q $IMAGE_NAME:$IMAGE_TAG)" ]; then
  echo "Remove existing Toxic REST API Docker image..."
  docker rmi $IMAGE_NAME:$IMAGE_TAG
fi

docker build -t "$IMAGE_NAME:$IMAGE_TAG" -f "$DOCKERFILE" .

if [ $? -eq 0 ]; then
  echo "Toxic REST API Docker image has been built: $IMAGE_NAME:$IMAGE_TAG."
  docker run -d --name $CONTAINER_NAME -p 8080:8080 "$IMAGE_NAME:$IMAGE_TAG"

  if [ "$(docker ps -q -f name=$CONTAINER_NAME)" ]; then
    echo "Toxic REST API Docker container has been started."
  else
    echo "Failed to start Toxic REST API Docker container."
  fi

else
  echo "Failed to build Toxic REST API Docker image."
fi
