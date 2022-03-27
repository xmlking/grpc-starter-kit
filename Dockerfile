# Accept the Go Micro version for the image to be set as a build argument.
ARG BASE_VERSION=latest

# First stage: build the executable.
FROM ghcr.io/xmlking/grpc-starter-kit/base:${BASE_VERSION} AS builder

# Set the environment variables for the go command:
# * CGO_ENABLED=0 to build a statically-linked executable
# * GOFLAGS=-mod=vendor to force `go build` to look into the `/vendor` folder.
# * GOPROXY="https://nexus.mycomp.com,direct" to set corp nexus goproxy and fallback to direct.
# * GOWORK=off to desable go.work which is used for local development
# e.g., ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor
ENV CGO_ENABLED=1 GOOS=linux GOPROXY="https://proxy.golang.org,direct" GOWORK=off
# when runing behind corporate proxy, you may need following settings
#ENV GOPRIVATE=github.com/mycomp
#ENV GONOSUMDB=github.com/*,gopkg.in/*,google.golang.org/*,cloud.google.com/*
#ENV GOPROXY="https://nexus.mycomp.com,direct"

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY go.mod go.sum ./
# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# COPY the source code as the last step
COPY ./ ./

# Build the executable to `/app`. Mark the build as statically linked.
ARG VERSION=0.0.1
ARG TYPE=service
ARG TARGET=account

RUN go build -a \
    -ldflags="-w -s -linkmode external -extldflags '-static'" \
    -o /app ./$TYPE/$TARGET/main.go

# Compress the binary
RUN upx -v --ultra-brute --best /app || true

# Final stage: the running container.

# Use `static:debug` to create debuggable container. for prod, use `static:nonroot`
FROM gcr.io/distroless/static:debug AS final
#FROM gcr.io/distroless/static:nonroot AS final

# copy grpc-health-probe to use with readinessProbe
COPY --from=builder --chown=nonroot:nonroot /bin/grpc_health_probe /bin/grpc_health_probe
# copygrpcurl to use with livenessProbe
COPY --from=builder --chown=nonroot:nonroot /bin/grpcurl /bin/grpcurl
# copy dumb-ini from base
COPY --from=builder --chown=nonroot:nonroot /usr/bin/dumb-init /bin/dumb-init
# copy the compiled executable from the second stage.
ARG VERSION=0.0.1
ARG TYPE=service
ARG TARGET=account
COPY --from=builder --chown=nonroot:nonroot /app /app
COPY --from=builder --chown=nonroot:nonroot src/config /config

# Declare the port on which the webserver will be exposed.
# As we're going to run the executable as an unprivileged user, we can't bind
# to ports below 1024.
EXPOSE 8080

# Perform any further action as an unprivileged user.
USER nonroot:nonroot

# Health Check
HEALTHCHECK --interval=15s --timeout=5s --retries=3 CMD /bin/grpc_health_probe -addr=:8080 -connect-timeout 250ms -rpc-timeout 100ms || exit 1

# Metadata params
ARG DOCKER_REGISTRY
ARG DOCKER_CONTEXT_PATH=xmlking
ARG BUILD_DATE
ARG VCS_URL=grpc-starter-kit
ARG VCS_REF=1
ARG VENDOR=sumo

# Metadata
LABEL org.opencontainers.image.created=$BUILD_DATE \
    org.opencontainers.image.name="${TARGET}-${TYPE}" \
    org.opencontainers.image.title="${TARGET}-${TYPE}" \
    org.opencontainers.image.description="Example of multi-stage docker build" \
    org.opencontainers.image.url=https://github.com/xmlking/$VCS_URL \
    org.opencontainers.image.source=https://github.com/xmlking/$VCS_URL \
    org.opencontainers.image.revision=$VCS_REF \
    org.opencontainers.image.version=$VERSION \
    org.opencontainers.image.authors=sumanth \
    org.opencontainers.image.vendor=$VENDOR \
    org.opencontainers.image.ref.name=$VCS_REF \
    org.opencontainers.image.licenses=MIT \
    org.opencontainers.image.documentation="docker run -it -e CONFY_ENV=production -p 8080:8080  ${DOCKER_REGISTRY:+${DOCKER_REGISTRY}/}${DOCKER_CONTEXT_PATH}/${TARGET}-${TYPE}:${VERSION}"

# Run the compiled binary.
ENTRYPOINT ["/bin/dumb-init", "/app"]
