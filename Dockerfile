FROM  golang:alpine
RUN mkdir app
COPY . /app
WORKDIR /app/src
# RUN  go mod init
RUN go build -o main .
CMD ["app/main"]
