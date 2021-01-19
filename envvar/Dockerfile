FROM golang:alpine as builder
WORKDIR /app
ADD main.go /app/main.go
RUN go build -o goserver main.go

FROM alpine
LABEL maintainer "Sofyan Sugianto <sofyanartem@gmail.com>"
WORKDIR /app
COPY --from=builder /app /app
ENV PORT 8000
EXPOSE 8000
CMD ["./goserver"]
