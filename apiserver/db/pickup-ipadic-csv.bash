#!/bin/bash

IPADIC_PATH=./mecab-ipadic-2.7.0-20070801

REQUIRED_CSVS="
Adj.csv
Adverb.csv
Noun.adjv.csv
Noun.adverbal.csv
Noun.csv
Noun.nai.csv
Noun.verbal.csv
Verb.csv
"

# "-p" means "make dir if it doesn't exist"
mkdir -p $PATH_TO

# copy csv files to db/csv and convert them to utf-8
for csv in $REQUIRED_CSVS; do
    iconv $IPADIC_PATH/$csv -f EUC-JP -t UTF-8 -o $PATH_TO/$csv
    echo convert $csv to utf-8
done