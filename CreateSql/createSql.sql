create database jdColly;

create table comments
(
    id                int auto_increment
        primary key,
    context           varchar(200)     null,
    en_context        varchar(200)     null,
    old_score         int              null,
    product_id        bigint           null,
    reference_time    bigint default 0 not null,
    useful_vote_count int              null
);

create table infoColly
(
    id         int auto_increment
        primary key,
    img        varchar(100) null,
    price      varchar(20)  null,
    name       varchar(100) null,
    product_id varchar(100) null,
    title      varchar(100) null,
    url        varchar(100) null,
    `key`      varchar(100) null
);


