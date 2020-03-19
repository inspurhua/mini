drop table sys_role;
create table sys_role
(
    id     bigserial primary key,
    name   text,
    status integer default 1 --1有效,2无效
);
insert into sys_role
values (1, '管理员', 1),
       (2, '操作员', 1);

drop table sys_org;
create table sys_org
(
    id   bigserial primary key,
    pid  integer,
    code text,--333编码
    name text,
    sort int
);
insert into sys_org
values (1, 0, '100', '我的组织', 1),
       (2, 1, '100101', 'A部门', 1),
       (3, 1, '100102', 'B部门', 2),
       (4, 1, '100103', 'C部门', 3),
       (5, 4, '100103101', 'C1部门', 1),
       (6, 4, '100103102', 'C2部门', 2);
drop table sys_entry;
create table sys_entry
(
    id     bigserial primary key,
    title  text,
    pid    integer,
    type   integer default 1,--1菜单2功能
    method text,
    href    text,
    icon   text,
    target text    default '_self',
    sort   int
);
insert into sys_entry
values (1, '系统管理', 0, 1, '', '', '', '', 1),
       (1000, '角色', 1, 1, '', 'sys/role.html', 'fa fa-group', '_self', 1),
       (1001, '角色列表', 1000, 2, 'GET', '/api/role', '', '', 1),
       (1002, '添加角色', 1000, 2, 'POST', '/api/role', '', '', 2),
       (1003, '编辑角色', 1000, 2, 'PUT', '/api/role/:id', '', '', 3),
       (1004, '删除角色', 1000, 2, 'DELETE', '/api/role/:id', '', '', 4),
       (1005, '查看角色', 1000, 2, 'GET', '/api/role/:id', '', '', 5),
       (1006, '查看授权', 1000, 2, 'GET', '/api/role/:id/auth', '', '', 6),
       (1007, '设置授权', 1000, 2, 'PUT', '/api/role/:id/auth', '', '', 7),

       (1100, '组织', 1, 1, '', 'sys/org.html', 'fa fa-sitemap', '_self', 1),
       (1101, '组织列表', 1100, 2, 'GET', '/api/org', '', '', 1),
       (1102, '添加组织', 1100, 2, 'POST', '/api/org', '', '', 2),
       (1103, '编辑组织', 1100, 2, 'PUT', '/api/org/:id', '', '', 3),
       (1104, '删除组织', 1100, 2, 'DELETE', '/api/org/:id', '', '', 4),
       (1105, '查看组织', 1100, 2, 'GET', '/api/org/:id', '', '', 5),


       (1200, '用户', 1, 1, '', 'sys/user.html', 'fa fa-user', '_self', 1),
       (1201, '用户列表', 1200, 2, 'GET', '/api/user', '', '', 1),
       (1202, '添加用户', 1200, 2, 'POST', '/api/user', '', '', 2),
       (1203, '编辑用户', 1200, 2, 'PUT', '/api/user/:id', '', '', 3),
       (1204, '删除用户', 1200, 2, 'DELETE', '/api/user/:id', '', '', 4),
       (1205, '查看用户', 1200, 2, 'GET', '/api/user/:id', '', '', 5),

       (1300, '日志', 1, 1, '', 'sys/log.html', 'fa fa-list', '_self', 1),
       (1301, '组织列表', 1300, 2, 'GET', '/api/log', '', '', 1),
       (2, '业务管理', 0, 1, '', '', '', '', 1);

drop table sys_auth;
create table sys_auth
(
    id       bigserial primary key,
    role_id  integer,
    entry_id integer
);


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

insert into sys_user(id, account, password, role_id, org_id)
values (1, 'zhanghua', md5('abc123456'), 1, 1);

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
drop table sys_file;
create table sys_file
(
    id        bigserial primary key,
    name      text,
    save_name text,
    save_path text,
    url       text,
    create_at timestamp without time zone default localtimestamp(0)
);

drop table sys_log;
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


