# First stage: build the executable.
FROM golang:1.16.0-alpine AS builder

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group && \
    apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init && \
    rm -rf /var/cache/apk/* /tmp/*  && \
    GRPC_HEALTH_PROBE_VERSION=v0.3.2 && \
    wget -q -O /bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

ENV GO111MODULE=on
ENV GOPROXY="https://proxy.golang.org,direct"

# Get dependancies - will also be cached if we won't change mod/sum
WORKDIR /
COPY .go.mod go.sum ./
RUN go mod download -x && \
    go get github.com/ahmetb/govvv && \
    go get github.com/markbates/pkger/cmd/pkger && \
    rm go.mod go.sum

# First stage: build the executable.
FROM ghcr.io/xmlking/grpc-starter-kit/base:${BASE_VERSION} AS builder

# Set the environment variables for the go command:
# * CGO_ENABLED=0 to build a statically-linked executable
# * GOFLAGS=-mod=vendor to force `go build` to look into the `/vendor` folder.
#ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor
ENV CGO_ENABLED=1 GOOS=linux
ENV GOPROXY="https://proxy.golang.org,direct"

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./
# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download -x

# COPY the source code as the last step
COPY ./ ./

# Build the executable to `/app`. Mark the build as statically linked.
ARG VERSION=0.0.1
ARG TYPE=service
ARG TARGET=account

RUN pkger -o $TYPE/$TARGET -include /config/config.yml -include /config/config.production.yml -include /config/certs
RUN go build -a \
    -ldflags="-w -s -linkmode external -extldflags '-static' $(govvv -flags -version ${VERSION} -pkg $(go list ./internal/config) )" \
    -o /app ./$TYPE/$TARGET/main.go

# Final stage: the running container.
FROM scratch AS final

# copy 1 MiB busybox executable
COPY --from=busybox:1.31.1 /bin/busybox /bin/busybox

# copy grpc-health-probe to use with readiness and liveness probes
COPY --from=builder /bin/grpc_health_probe /bin/grpc_health_probe

# copy dumb-ini from base
COPY --from=builder /usr/bin/dumb-init /usr/bin/dumb-init

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled executable from the second stage.
ARG VERSION=0.0.1
ARG TYPE=service
ARG TARGET=account
COPY --from=builder /app /app
COPY --from=builder src/config /config

# Declare the port on which the webserver will be exposed.
# As we're going to run the executable as an unprivileged user, we can't bind
# to ports below 1024.
EXPOSE 8080

# Perform any further action as an unprivileged user.
USER nobody:nobody

# Metadata params
ARG DOCKER_REGISTRY
ARG DOCKER_CONTEXT_PATH=xmlking
ARG BUILD_DATE
ARG VCS_URL=grpc-starter-kit
ARG VCS_REF=1
ARG VENDOR=sumo

# Metadata
LABEL org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.name="${TARGET}-${TYPE}" \
    org.label-schema.description="Example of multi-stage docker build" \
    org.label-schema.url="https://example.com" \
    org.label-schema.vcs-url=https://github.com/xmlking/$VCS_URL \
    org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.vendor=$VENDOR \
    org.label-schema.version=$VERSION \
    org.label-schema.docker.schema-version="1.0" \
    org.label-schema.docker.cmd="docker run -it -e CONFIG_ENV=production -p 8080:8080  ${DOCKER_REGISTRY:+${DOCKER_REGISTRY}/}${DOCKER_CONTEXT_PATH}/${TARGET}-${TYPE}:${VERSION}"

# Run the compiled binary.
ENTRYPOINT ["/usr/bin/dumb-init", "/app"]
