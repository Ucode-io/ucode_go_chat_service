FROM golang:1.24.0 as builder

RUN mkdir -p $GOPATH/src/gitlab.udevs.io/ucode/ucode_go_chat_service 
WORKDIR $GOPATH/src/gitlab.udevs.io/ucode/ucode_go_chat_service

COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/ucode_go_chat_service /

ENTRYPOINT ["/ucode_go_chat_service"]
