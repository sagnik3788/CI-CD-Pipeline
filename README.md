# CI/CD Pipeline

## Step 1: Create a Basic Application

Create a basic application that you want to deploy using the pipeline. This could be a simple web server, such as a Go server (`main.go` and `go.mod`).

## Step 2: Dockerize the Application

Create a `Dockerfile` to build the Docker image for your application and test it locally.

## Step 3: Setup Jenkins Environment

Create a new Jenkins image that includes Docker, so it can build and manage Docker images.

```Dockerfile
FROM jenkins/jenkins:lts
USER root
RUN apt-get update
RUN curl -sSL https://get.docker.com/ | sh
