FROM golang:1.18.3-alpine 
WORKDIR /app
COPY ./../ /app/
WORKDIR /app/cmd/
RUN go build -o /server
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh
RUN apk add --no-cache bash
EXPOSE 8080

CMD ["/wait-for-it.sh","auth_db:3306", "--","/server"]