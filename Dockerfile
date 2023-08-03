# syntax=docker/dockerfile:1
FROM golang:1.20 AS build-stage
ARG APP_VERSION
ENV APP_VERSION=${APP_VERSION:-unknown}
RUN echo "APP_VERSION=$APP_VERSION"


#LABEL authors="hschellenberger"
LABEL org.opencontainers.image.source=https://github.com/hkonitzer/OHMAB
LABEL org.opencontainers.image.description="OHMAB $APP_VERSION"
LABEL org.opencontainers.image.version=$APP_VERSION

WORKDIR /app
COPY go.mod go.sum ./
COPY templates/ ./templates/
COPY ent/ ./ent/
COPY internal/ ./internal/
COPY cmd/ ./cmd/
COPY *.graphql ./
COPY *.go ./
COPY config.yml /

RUN go mod download

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-X 'github.com/hkonitzer/ohmab/internal/pkg/common/config.Version=$APP_VERSION'" -o /server cmd/server.go


#FROM build-stage AS run-test-stage
#RUN go test -v .

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11:latest-amd64 AS build-release-stage

WORKDIR /

COPY --from=build-stage /server /server
COPY --from=build-stage /config.yml /config.yml

EXPOSE 8184

USER nonroot:nonroot

ENTRYPOINT  ["/server"]

