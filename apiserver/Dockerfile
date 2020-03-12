FROM golang:latest as builder

ENV IPADIC_URL https://drive.google.com/uc?export=download&id=0B4y35FiV1wh7MWVlSDBCSXZMTXM
ENV TMP_SQL_DIR /tmp/sqlite
ENV TMP_EXEC_DIR /tmp/apiserver
ENV SQL_NAME words.sqlite3
# necessary to work go-sqlite3 in alpine
# (static link to cgo are required to work without gcc)
ENV CGO_ENABLED=1

RUN mkdir /go/src/apiserver
WORKDIR /go/src/apiserver

# copy files from host to container
ADD . /go/src/apiserver

RUN apt-get update \
    && apt-get install -y sqlite3 libsqlite3-dev \
    && apt-get clean

# create sqlite from mecab ipadic
RUN wget -O ipadic.tar.gz ${IPADIC_URL} \
    && tar zxf ipadic.tar.gz \
    && mkdir -p ${TMP_SQL_DIR} \
    && COPY_TO=${TMP_SQL_DIR} SQL_NAME=${SQL_NAME} db/setup-db.bash

# build app
RUN cp ${TMP_SQL_DIR}/${SQL_NAME} ./main \
    && cd ./main \
    # options make go-sqlite3 work on alpine
    # (static link to cgo are required to work without gcc)
    && go build -a -ldflags '-linkmode external -extldflags "-static"'

FROM alpine:latest as production

RUN mkdir /main
WORKDIR /main
# NOTE: this copies all files inside, NOT directory
COPY --from=builder /go/src/apiserver/main .
ENTRYPOINT [ "./apiserver" ]
