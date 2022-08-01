FROM golang:1.17-buster

#ENV TZ=Europe/Moscow
#RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone


WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go build -o main ./cmd/

CMD [ "./main" ]