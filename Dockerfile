ARG GO_VERSION=1.24.2
ARG ALPINE_VERSION=3.21.3


FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS builder

LABEL maintainer="miguel"
LABEL env="production"

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o server ./cmd/main.go

# Final stage
FROM alpine:${ALPINE_VERSION}
COPY --from=builder /app/server /

ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

HEALTHCHECK --interval=10s --timeout=5s --start-period=5s --retries=3 \
    CMD [ "curl","-f","http://localhost:8080/ping" ]

ENTRYPOINT ["./server"]
CMD ["8080"]
