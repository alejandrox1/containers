# containers

Everything one could wish to know about containers

## Docker 101

### Contents

* Networks
* Port Mappings
* Exec'ing Into Containers

* Volumes

* Image Building
* CMD v. ENTRYPOINT


**SPecifying container Network**
```
docker build -t server . && docker run --rm --network host server
```

curl as normal

```
curl localhost:8000/
```

**Exec'ing Inside a Container**

```
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
13de88e52858        server              "./app"             6 seconds ago       Up 4 seconds                            vigilant_cray
```

```
$ docker exec -it 13de88e52858 bash
root@N501VW:/app# ls
app  file.log  go.mod  main.go
root@N501VW:/app# cat file.log
```


**Port Mapping**
```
$ docker build -t server . && docker run --rm -p 3000:8000 server
```

```
$ curl localhost:3000/
```

Mapping multiple ports and using UDP

```
$ docker build -t server . && docker run --rm -p 3000:8000 -p 8000:9000/udp server
```

```
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS                                            NAMES
5c4386fe21ce        server              "./app"             4 seconds ago       Up 3 seconds        0.0.0.0:3000->8000/tcp, 0.0.0.0:8000->9000/udp   infallible_mayer
```

**Volumes**
```
$ docker build -t server . && docker run --rm -it -v ${PWD}:/app -p 8000:8000 server
```
