FROM golang:1.13 as builder
COPY . /workspace 
WORKDIR /workspace
RUN ["make", "build"]

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
COPY --from=builder /workspace/word-count /app/
WORKDIR /app/
