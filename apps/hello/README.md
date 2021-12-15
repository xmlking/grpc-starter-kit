# go-app

WebApp built with [go-app](https://go-app.dev/)

BUILD THE CLIENT

```bash
GOARCH=wasm GOOS=js go build -o web/app.wasm
```

BUILD THE SERVER

```bash
go build
```

RUN THE APP

> always rebuild the client 

```bash
go run .
# or
./hello

open http://localhost:8000/
```


## Reference 
- https://go-app.dev/getting-started
- passwordless [login](https://github.com/teamhanko/apple-wwdc21-webauthn-example)
