FROM golang:1.20-alpine AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /
COPY services.waitlist services.waitlist
COPY services.shared ../services.shared
WORKDIR /services.waitlist
ENV CGO_ENABLED=0
COPY ./services.waitlist/go.mod ./services.waitlist/go.sum ./
RUN  --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . . 
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o main ./src/cmd/main.go

FROM scratch

ENV PORT 8080

COPY --from=builder /services.waitlist/main .
COPY --from=builder /services.waitlist/src/locales ./src/locales

EXPOSE $PORT

CMD ["/main"]