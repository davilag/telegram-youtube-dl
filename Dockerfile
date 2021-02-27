FROM golang:1.15-alpine AS go-builder

RUN apk update && apk add ca-certificates

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o telegram-youtube-dl .

FROM python:3.9-alpine AS final-image

# Install youtube-dl
RUN wget https://yt-dl.org/downloads/latest/youtube-dl
RUN mv youtube-dl /usr/local/bin/youtube-dl
RUN chmod a+rx /usr/local/bin/youtube-dl

# Copy binaries from go-builder
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go-builder /app/.env /.env
COPY --from=go-builder /app/telegram-youtube-dl /telegram-youtube-dl



ENTRYPOINT [ "/telegram-youtube-dl" ]