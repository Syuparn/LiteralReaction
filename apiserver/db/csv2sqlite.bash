mkdir -p $SQL_PATH

# create sqlite file
# NOTE: sqlite can receive string(heredoc) as manipulation
sqlite3 $SQL_PATH/$SQL_NAME << EOS
/* create tables */
.read ./db/init.sql

/* let col separator "," */
.separator ','

/* import word data csvs */
.mode csv
.import $CSV_PATH/adj.csv adjectives
.import $CSV_PATH/adverb.csv adverbs
.import $CSV_PATH/noun.csv nouns
.import $CSV_PATH/verb.csv verbs

/* insert each table count into counts */
insert into counts select 'adjectives', count(*) from adjectives;
insert into counts select 'adverbs', count(*) from adverbs;
insert into counts select 'nouns', count(*) from nouns;
insert into counts select 'verbs', count(*) from verbs;
EOS

echo created $SQL_NAME
