# Etapa de build
FROM golang:1.23.2-alpine AS builder

RUN apk --update-cache upgrade && apk add --no-cache git mercurial ca-certificates tzdata gcc musl-dev sqlite-dev

ENV APPNAME teste
ENV SRCPATH /src/${APPNAME}
ENV BINPATH /${APPNAME}

COPY . ${SRCPATH}
WORKDIR ${SRCPATH}/cmd

RUN go mod tidy \
    && CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -ldflags "-extldflags '-static'" -o ${BINPATH}/${APPNAME} .

# Etapa final
FROM alpine

RUN apk --no-cache add sqlite

ENV APPNAME teste
ENV DB_FILE /${APPNAME}/test.db

COPY --from=builder /${APPNAME} /${APPNAME}
WORKDIR /${APPNAME}

CMD ["./teste"]



