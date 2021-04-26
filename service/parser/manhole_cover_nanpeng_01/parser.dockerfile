FROM 192.168.1.1 as builder
WORKDIR ${GOPATH}/src/github.com/zsy-cn/bms/service/parser/manhole_cover_nanpeng_01
COPY . .
RUN go get -d -v
RUN set -x && CGO_ENABLED=0 GOOS=linux go build -a -o /app ./cmd

FROM 192.168.1.1/alpine

ENV GOPATH /root/go
COPY --from=builder /app .

CMD ["/app"]