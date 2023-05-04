FROM --platform=$BUILDPLATFORM golang:1.20.4-alpine3.17 AS build

WORKDIR /src

COPY . .

ARG TARGETOS TARGETARCH

RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /out/task .

FROM gcr.io/distroless/static-debian11

COPY --from=build /out/task /bin

ENTRYPOINT ["/bin/task"]
