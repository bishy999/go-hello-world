# pull and run image
```
sudo docker run --name=my-running-webapp -d -p 8080:80 bishy999/golang:1.4-my-webapp
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

sudo docker tag my-webapp bishy999/golang:1.4-my-webapp
```


# give terminal your docker hub credentials

```
sudo docker login
```


# push image to docker hub

```
docker push <docker hub username>/<repository><image name>

e.g. on Mac

sudo docker push bishy999/golang:1.4-my-webapp
```


# check docker hub

```
image used in example here is stored on docker hub e.g https://hub.docker.com/r/bishy999/golang
```
