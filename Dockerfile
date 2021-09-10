FROM golang:alpine as builder
COPY . /app
WORKDIR /app
RUN ls
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOARM=6 go build -ldflags '-w -s' -o githubDown


FROM alpine:latest
MAINTAINER jw-star
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata
RUN mkdir /app /down /config
COPY --from=builder /app/githubDown /app/
COPY  --from=builder     /app/conf.yml /config
WORKDIR /app
CMD ["/app/githubDown"]

#docker rmi jwstar/github_down:latest
# docker build -t jwstar/github_down .
#docker run --name down   jwstar/github_down



#docker rm -f down