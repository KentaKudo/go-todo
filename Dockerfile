# https://medium.com/@pierreprinetti/the-go-1-11-dockerfile-a3218319d191

ARG GO_VERSION=1.11

FROM golang:${GO_VERSION}-alpine AS builder

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN apk add --no-cache git

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app \
    github.com/KentaKudo/goapi-skel/cmd/todod

FROM scratch AS final

COPY --from=builder /user/group /user/passwd /etc/

COPY --from=builder /app /app

EXPOSE 8080

USER nobody:nobody

ENTRYPOINT [ "/app" ]
