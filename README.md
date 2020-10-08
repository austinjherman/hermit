# Hermit
Going to be ... a command line utility filled with useful everyday stuff.


## Set Up
I set this up in docker / docker-compose as to have a fresh OS for building and compiling. This is pretty convenient because it means you don't have to install golang on your system or worry about complex folder structures or annoying evironmental conflicts. 

This means, of course, that if you want to work on this repo locally you need to have docker and docker-compose installed on your system. 


## Commands so I remember

`docker-compose up -d --build`
  - will set up your container and volumes

`docker-compose run hermit go run app`
  - will run your app in the container

`docker-compose run hermit go build -o ./bin/hermit`
  - builds a linux binary

`docker-compose run -e GOOS=darwin hermit go build -o ./bin/hermit` 
  - builds a Mac binary