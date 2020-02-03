FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get github.com/gorilla/mux
RUN go get github.com/BurntSushi/toml
RUN go get gopkg.in/mgo.v2/bson
RUN go get gopkg.in/mgo.v2
RUN go get github.com/patrickmn/go-cache
RUN go build -o main . 
CMD ["/app/main"]
EXPOSE 9090