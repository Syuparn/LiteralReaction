#!/bin/bash
UTF8_IPADIC_PATH=./db/ipadic
SHAPED_CSV_PATH=./db/csv
SQL_PATH=./db/sqlite

PATH_TO=$UTF8_IPADIC_PATH \
    db/pickup-ipadic-csv.bash
PATH_FROM=$UTF8_IPADIC_PATH PATH_TO=$SHAPED_CSV_PATH \
    db/format-ipadic-csv.bash
CSV_PATH=$SHAPED_CSV_PATH SQL_PATH=$SQL_PATH SQL_NAME=$SQL_NAME\
    db/csv2sqlite.bash

cp -r $SQL_PATH/words.sqlite3 $COPY_TO
