#!/bin/bash

mkdir -p $PATH_TO

# make adjective
cat $PATH_FROM/Adj.csv       | awk -F"," '$10~/^基本形$/ {print $1}' >  $PATH_TO/adj.txt
cat $PATH_FROM/Noun.adjv.csv | awk -F"," '{print $1 "な"}'         >> $PATH_TO/adj.txt
cat $PATH_FROM/Noun.nai.csv  | awk -F"," '{print $1 "ない"}'       >> $PATH_TO/adj.txt

# make noun
cat $PATH_FROM/Noun.csv          | awk -F"," '{print $1}' >  $PATH_TO/noun.txt
cat $PATH_FROM/Noun.verbal.csv   | awk -F"," '{print $1}' >> $PATH_TO/noun.txt
cat $PATH_FROM/Noun.adverbal.csv | awk -F"," '{print $1}' >> $PATH_TO/noun.txt

# make adverb
cat $PATH_FROM/Adverb.csv        | awk -F"," '{print $1}'                    >  $PATH_TO/adverb.txt
cat $PATH_FROM/Noun.adverbal.csv | awk -F"," '{print $1}'                    >> $PATH_TO/adverb.txt
## 連用形接続形容詞（小さい「っ」で終わるものは動詞が後続できないので削除）
cat $PATH_FROM/Adj.csv           | awk -F"," '$10~/^連用テ接続$/ {print $1}' \
                                 | awk       '!/っ$/'                        >> $PATH_TO/adverb.txt


# make verb
cat $PATH_FROM/Verb.csv        | awk -F"," '$10~/^基本形$/ {print $1}'        >  $PATH_TO/verb.txt
cat $PATH_FROM/Noun.verbal.csv | awk -F"," '{print $1 "する"}'                >> $PATH_TO/verb.txt

# convert txt to csv (enumerate and add headers)
for txt in $PATH_TO/*.txt; do
    # NOTE: 同じ表記で読みだけ違う単語はまとめる(uniqはsortと併用必須のためsort -u使用)
    cat $txt | sort -u | awk '{print NR "," $1}' > ${txt%.txt}.csv
    echo created ${txt%.txt}.csv
done

# remove tmp txts
for txt in $PATH_TO/*.txt; do
    rm $txt
done