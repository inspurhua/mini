drop table sys_user;
create table sys_user
(
    id        bigserial primary key,
    account   text    not null,
    password  text    not null,
    role_id   integer not null,
    org_id    integer not null,
    status    integer                     default 1,--1有效2无效
    real_name text    null,
    code      text    null,
    position  text    null,
    email     text    null,
    tel       text    null,
    avatar    text    null,
    address   text    null,
    gender    integer                     default 1,--1男2女
    state     integer                     default 1,--1在职2离职
    note      text    null,
    open_id   text    null,
    create_at timestamp without time zone default localtimestamp(0),
    update_at timestamp without time zone,
    delete_at timestamp without time zone
);

create table sys_role
(
    id     bigserial primary key,
    name   text,
    status integer default 1 --1有效,2无效
);

create table sys_entry
(
    id     bigserial primary key,
    title  text,
    pid    integer,
    type   integer default 1,--1菜单2功能
    url    text,
    icon   text,
    target text    default '_self',
    sort   int
);

create table sys_auth
(
    id       bigserial primary key,
    role_id  integer,
    entry_id integer
);

-- create table sys_token
-- (
--     id        bigserial primary key,
--     account   text,
--     token     text,
--     jwt_token text,
--     iat       integer,
--     exp       integer,
--     create_at timestamp without time zone default localtimestamp(0)
-- );

create table sys_file
(
    id        bigserial primary key,
    name      text,
    save_name text,
    save_path text,
    url       text,
    create_at timestamp without time zone default localtimestamp(0)
);

create table sys_log
(
    id        bigserial primary key,
    user_id   integer,
    uri       text,
    data      text,
    ip        text,
    ua        text,
    method    text,
    create_at timestamp without time zone default localtimestamp(0)
);


create table sys_org
(
    id   bigserial primary key,
    pid  integer,
    code text,--333编码
    name text
);
