create database jdColly;

create table comments
(
    id         int auto_increment
        primary key,
    context    varchar(200) null,
    en_context varchar(200) null,
    old_score  int          null,
    product_id bigint       null
);
