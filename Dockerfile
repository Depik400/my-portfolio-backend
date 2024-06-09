FROM golang:1.22.1

RUN apt-get update
RUN apt-get install curl -y

RUN  curl -sL https://deb.nodesource.com/setup_20.x | bash

RUN apt-get install -y nodejs

EXPOSE 80

WORKDIR /application

COPY . .

RUN cp .env.docker .env

RUN cd public/frontend && npm i && npm run build && cd ../../ && go build --tags=prod  -o /main cmd/main.go

CMD /main