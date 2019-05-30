FROM golang:1.12-alpine AS backend-build
WORKDIR /hello/
COPY ./ ./
RUN go build -o hello
# ---- Release Alpine Image ----
FROM alpine:latest AS deploy
RUN apk --no-cache add ca-certificates
WORKDIR /hello/
COPY --from=backend-build /hello/hello ./
CMD [ "sh", "-c", "./hello" ]
