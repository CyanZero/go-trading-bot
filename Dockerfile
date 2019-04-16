# docker build -t bot .
# docker run -it --rm -p 8080:8080 --name trading_bot bot
# export GOPATH=/Users/calvinzhou/golang/trading_bot:/Users/calvinzhou/go
# Delete all containers
# docker rm $(docker ps -a -q)
# Delete all images
# docker rmi $(docker images -q)

FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./src/cyan.io/ /go/src/cyan.io/
WORKDIR /go/src/cyan.io/trading_bot/

RUN go get ./
RUN go build

# CMD if [ ${APP_ENV} = production ]; \
# 	then \
# 	app; \
# 	else \
# 	go get github.com/pilu/fresh && \
# 	fresh; \
# 	fi

CMD trading_bot

EXPOSE 8080