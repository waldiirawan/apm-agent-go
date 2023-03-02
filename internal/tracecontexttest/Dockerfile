FROM golang:latest
ADD . /go/src/github.com/waldiirawan/apm-agent-go
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org
WORKDIR /go/src/github.com/waldiirawan/apm-agent-go/internal/tracecontexttest
RUN go build -o /trace-context-service main.go

EXPOSE 5000/tcp
HEALTHCHECK CMD curl -X POST -H "Content-Type: application/json" -d "{}" http://localhost:5000
CMD /trace-context-service
