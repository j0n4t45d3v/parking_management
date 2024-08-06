FROM golang:1.22.5-alpine as build 

WORKDIR /parking_management

COPY . /parking_management

RUN apk add --no-cache make

RUN go mod tidy

RUN make build

FROM scratch 

WORKDIR /api

COPY --from=build /parking_management/bin/* /api/

EXPOSE 8000

ENTRYPOINT [ "./main" ]
