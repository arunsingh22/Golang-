## DOCKER TERMS
--------------------------------------------------------------------------------------------------------------------------
- A container is a normal operating system process except that this process is isolated and has its own file system, its own networking, and its own isolated process tree separate from the host.

export GO111MODULE=on;export GOPROXY=direct;export GOSUMDB=off

## WORKSPACE CLEANUP
sudo docker container rm $(sudo docker container ls -aq) -f
sudo docker image rm $(sudo docker images -q) -f

## DOCKER CMDS
-------------------------------------------------------------------------------------------------------------------------
- sudo docker build -t <ImageName:Tag> <-f <fileName.dockerfile> >

### RUNNING CONTAINER:
- sudo docker run -d -p [host_port]:[container_port] --name <containerName> <imageName> 
example:  docker run -d -p 8080:8080 --name <containerName> <imageName>


### START/STOP/REMOVING CONTAINER
- sudo docker container start <containerName>
- sudo docker container stop <containerName>
- sudo docker container rm <containerName> <containerName2> <...>


Use this for multi-stage dockerfile for deployment.
**https://github.com/GoogleContainerTools/distroless** 

sudo docker ps // list only running containers
sudo docker ps -a // list all containers 

 
sudo docker system prune 
sudo docker system prune -a


docker images  // list all images in local system.
docker images rm <ImageID>



-------------------------------------------------------------------------------------------

CODE Quality
1. modularize code ALWAYS if you have more than 15 lines 
2. Don't use pairs use class/struct
3. use preIndex = i-1, instead of i-1 everywhere 
4. Do not use long comparision rather break it down into if claues.
5. Give comments to complex intricate things.
--------------------------------------------------------------------------------------------------------------------

//FixMe: won't passing a ref be better than copy.
//FixMe (db.go) item struct doesn't accomated newer values 




4. functional testing.

-------------------------------------------------
services have config files always and not libarary 
homogenity of code is important. 


