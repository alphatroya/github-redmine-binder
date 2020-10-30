FROM alpine:3.12.1

RUN apk add --no-cache git make musl-dev go

ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir /app
COPY . /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]

EXPOSE 8933
