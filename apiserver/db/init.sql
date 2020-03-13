drop table if exists adjectives;
create table adjectives (
    id integer primary key,
    word text
);

drop table if exists adverbs;
create table adverbs (
    id integer primary key,
    word text
);

drop table if exists nouns;
create table nouns (
    id integer primary key,
    word text
);

drop table if exists verbs;
create table verbs (
    id integer primary key,
    word text
);

drop table if exists counts;
create table counts
(
    table_name varchar(255),
    row_count int
);

drop table if exists favorite_sentences;
create table favorite_sentences (
    id integer primary key autoincrement,
    former_pos text,
    latter_pos text,
    particle text,
    former_word text,
    latter_word text
);
