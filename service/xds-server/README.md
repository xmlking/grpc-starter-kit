# Envoy XDS Server

This is a sample repo which demonstrates how to spin up an xDS Server for Envoy Proxy. 

## Sample Apps

Run some sample apps in docker to give some endpoints to route to:
```
docker run -d --rm --name=echo9100 -p 9100:8080 stevesloka/echo-server echo-server --echotext=Sample-Endpoint!
docker run -d --rm --name=echo9101 -p 9101:8080 stevesloka/echo-server echo-server --echotext=Sample-Endpoint!
docker run -d --rm --name=echo9102 -p 9102:8080 stevesloka/echo-server echo-server --echotext=Sample-Endpoint!
docker run -d --rm --name=echo9103 -p 9103:8080 stevesloka/echo-server echo-server --echotext=Sample-Endpoint!
docker run -d --rm --name=echo9104 -p 9104:8080 stevesloka/echo-server echo-server --echotext=Sample-Endpoint!
```

## Stop All Sample Apps

Stop all the sample endpoints created in the previous step:
```
docker stop echo9100
docker stop echo9101
docker stop echo9102
docker stop echo9103
docker stop echo9104
```


go run service/xds-server/cmd/server/main.go


xds-server

keep k8s resource files here and keep cmd/main.go  in grpc-starter-kit
https://github.com/stevesloka/envoy-xds-server
https://www.youtube.com/watch?v=qAuq4cKEG_E


https://github.com/tak2siva/Envoy-Pilot

kube:
https://github.com/pradeepmvn/xds-controller
https://github.com/tommy351/kubenvoy/blob/master/pkg/kds/server.go

logger wrapper
https://github.com/tommy351/kubenvoy/blob/master/pkg/envoy/logger.go
