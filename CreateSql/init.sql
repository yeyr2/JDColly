create table if not exists comments
(
    id                bigint auto_increment
        primary key,
    context           varchar(3000)     null,
    en_context        varchar(3000)     null,
    old_score         int              null,
    product_id        bigint           null,
    reference_time    bigint default 0 null,
    useful_vote_count int              null
);

create table if not exists infoColly
(
    id         bigint auto_increment
        primary key,
    img        varchar(100) null,
    price      varchar(20)  null,
    name       varchar(500) null,
    product_id varchar(100) null,
    title      varchar(500) null,
    url        varchar(100) null,
    `key`      varchar(100) null
);

create table if not exists user(
    id                  bigint auto_increment primary key ,
    nickname            varchar(100) not null ,
    username            varchar(50) not null,
    password            varchar(100) not null,
    sex                 varchar(2) not null default '未知',
    phone_number        varchar(20)  null default '无',
    email               varchar(30)  null default '无',
    address             varchar(100) null default '无',
    emergency_contact   varchar(20) default '无'
#     login int default 0 check ( login in (0,1))
);

create table if not exists search(
    id bigint auto_increment primary key ,
    user_id bigint ,
    `key` varchar(100),
    create_time bigint,
    update_time bigint
);