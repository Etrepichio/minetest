# Docker Dev Setup

## Requirements
  - Docker
  - Docker Compose (If you have docker for mac this is auto included)
  - Go environment properly setup
  
## Setup
  - You will need to create a `local-docker-dev-network`
    - This is so multiple docker-compose projects can communicate with each other
 
 **Create Docker Network:**
 
 `$ docker network create local-docker-dev-network`
## How to run
 cd into this directory then run the following command
 
 `$ ./compile_run.sh`
  
  This does the following: 

   1. Builds the go project on your local machine
   1. Adds compiled binary to docker image
   1. Runs minesweepersvc binary inside of docker
   
Next go to `http://localhost:48080` to verify all is running as expected