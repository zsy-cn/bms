FROM 192.168.1.1 as builder
WORKDIR ${GOPATH}/src/github.com/zsy-cn/bms/loraclient
COPY . .
RUN go get -d -v
RUN set -x && CGO_ENABLED=0 GOOS=linux go build -a -o /app 

FROM 192.168.1.1/alpine

ENV GOPATH /root/go
COPY --from=builder /app .
COPY --from=builder ${GOPATH}/src/github.com/zsy-cn/bms/loraclient/config/http.pem /config/http.pem

CMD ["/app"]