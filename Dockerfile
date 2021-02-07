FROM golang:1.15-alpine AS go-builder

RUN apk update && apk add ca-certificates

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o telegram-twitter-dl .

FROM scratch AS final-image

COPY --from=go-builder /app/.env /.env
COPY --from=go-builder /app/telegram-twitter-dl /telegram-twitter-dl
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


ENTRYPOINT [ "/telegram-twitter-dl" ]