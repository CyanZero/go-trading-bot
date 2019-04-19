# docker build -t bot . \
#    --build-arg CAL_BOT_AUTH= \
#    --build-arg CAL_BOT_SECRET= \
#    --build-arg CAL_BOT_LOG_PATH=
#
# docker run -it --rm -p 8080:8080 --name trading_bot trading_bot
# export GOPATH=/Users/calvinzhou/golang/trading_bot:/Users/calvinzhou/go
# Delete all containers
# docker rm $(docker ps -a -q)
# Delete all images
# docker rmi $(docker images -q)

FROM golang

ARG app_env
ARG CAL_BOT_AUTH
ARG CAL_BOT_SECRET
ARG CAL_BOT_LOG_PATH

ENV APP_ENV $app_env
ENV CAL_BOT_AUTH $CAL_BOT_AUTH
ENV CAL_BOT_SECRET $CAL_BOT_SECRET
ENV CAL_BOT_LOG_PATH $CAL_BOT_LOG_PATH

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

#CMD trading_bot
CMD ["sleep", "infinity"]

EXPOSE 8080