
## Status
[![Build Status](https://travis-ci.com/bishy999/go-hello-world.svg?branch=master)](https://travis-ci.com/bishy999/go-hello-world)
[![Go Report Card](https://goreportcard.com/badge/github.com/bishy999/go-hello-world)](https://goreportcard.com/report/github.com/bishy999/go-hello-world)
![GitHub Repo size](https://img.shields.io/github/repo-size/bishy999/go-hello-world)
[![GitHub Tag](https://img.shields.io/github/tag/bishy999/go-hello-world.svg)](https://github.com/bishy999/go-hello-world/releases/latest)
[![GitHub Activity](https://img.shields.io/github/commit-activity/m/bishy999/go-hello-world)](https://github.com/bishy999/go-hello-world)
[![GitHub Contributors](https://img.shields.io/github/contributors/bishy999/go-hello-world)](https://github.com/bishy999/go-hello-world)


# pull and run image
```
sudo docker run --name=my-running-webapp -d -p 8080:80 bishy999/golang:1.5-my-webapp
```


# check app is working via browser/cli
 
```
http://localhost:8080
curl http://localhost:8080
```


##########################################################################################
#                  Steps to build and deploy simple golang webapp with docker            #
##########################################################################################


Note: this is automatically done via travis (see .travis.yml) but manual steps are listed here for reference


# build image (don't use cache)

```
sudo docker build --no-cache -t my-webapp .
```


# list images

```
sudo docker images
```
 

# create a container from your image and run it
 
```
sudo docker run --name=my-running-webapp -d -p 8080:80 my-webapp
```


# tag image

```
sudo docker tag <image ID>  <docker hub username>/<repository><image name>:<version label or tag>

e.g.

sudo docker tag my-webapp bishy999/golang:1.5-my-webapp
```


# give terminal your docker hub credentials

```
sudo docker login
```


# push image to docker hub

```
docker push <docker hub username>/<repository><image name>

e.g. on Mac

sudo docker push bishy999/golang:1.5-my-webapp
```


# check docker hub

```
image used in example here is stored on docker hub e.g https://hub.docker.com/r/bishy999/golang
```
