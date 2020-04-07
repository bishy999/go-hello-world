#############################################################################################
#                       Build and deploy simple golang webapp with docker                   #
#############################################################################################



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

sudo docker tag 9131a73eae3a bishy999/golang:1.0-my-webapp
```


# give terminal your docker hub credentials

```
sudo docker login
```


# push image to docker hub

```
docker push <docker hub username>/<repository><image name>

e.g. on Mac

sudo docker push bishy999/golang:1.0-my-webapp
```

# check docker hub

```
Your image should be there e.g https://hub.docker.com/
```

# pull and run image
```
sudo docker run --name=my-running-webapp -d -p 8080:80 bishy999/golang:1.0-my-webapp
```


# check app is working via browser
 
```
http://localhost:8080
```