--多租户,超级管理员进来设置租户信息,新建租户默认形成当前租户
--下面的管理员角色,组织,管理员用户,并且这些不允许删除
--形成管理员角色,根组织后,会写admin_role,root_org_id,root_org_code,
--超级管理员,租户id 0 角色0 组织0
--TODO
drop table if exists sys_tenant;
create table sys_tenant
(
    id            bigserial primary key,
    name          text,
    status        integer      default 1, --1有效,2无效
    role_admin    int not null default 1,--管理员角色id
    root_org_id   int not null default 1,--根组织编号
    root_org_code int not null default '100'
);
insert into sys_tenant(id, name)
values (1, 'default');
SELECT setval('sys_tenant_id_seq', (select max(id) from sys_tenant), true);
--角色
drop table if exists sys_role;
create table sys_role
(
    id        bigserial primary key,
    name      text,
    status    integer default 1, --1有效,2无效
    tenant_id int not null
);

insert into sys_role(id, name, status, tenant_id)
values (1, '管理员', 1, 1),
       (2, '操作员', 1, 1);
SELECT setval('sys_role_id_seq', (select max(id) from sys_role), true);
--组织
drop table if exists sys_org;
create table sys_org
(
    id        bigserial primary key,
    pid       integer,
    code      text,--333编码,后台需要检测!!
    name      text,
    sort      int,
    tenant_id int not null
);
insert into sys_org(id, pid, code, name, sort, tenant_id)
values (1, 0, '100', '我的组织', 1, 1),
       (2, 1, '100101', 'A部门', 1, 1),
       (3, 1, '100102', 'B部门', 2, 1),
       (4, 1, '100103', 'C部门', 3, 1),
       (5, 4, '100103101', 'C1部门', 1, 1),
       (6, 4, '100103102', 'C2部门', 2, 1);
SELECT setval('sys_org_id_seq', (select max(id) from sys_org), true);
--菜单
drop table if exists sys_entry;
create table sys_entry
(
    id     bigserial primary key,
    title  text,
    pid    integer,
    type   integer      default 1,--1菜单2功能
    method text,
    href   text,
    icon   text,
    target text         default '_self',
    sort   int,
    kind   int not null default 0 --0.common 1.chaoji 2.putong
);
insert into sys_entry
values (1, '系统管理', 0, 1, '', '', '', '', 1, 0),

       (900, '租户', 1, 1, '', 'sys/tenant.html', 'fa fa-group', '_self', 1, 1),
       (901, '租户列表', 900, 2, 'GET', '/api/tenant', '', '', 1, 1),
       (902, '添加租户', 900, 2, 'POST', '/api/tenant', '', '', 2, 1),
       (903, '编辑租户', 900, 2, 'PUT', '/api/tenant/:id', '', '', 3, 1),
       (904, '删除租户', 900, 2, 'DELETE', '/api/tenant/:id', '', '', 4, 1),
       (905, '查看租户', 900, 2, 'GET', '/api/tenant/:id', '', '', 5, 1),

       (1000, '角色', 1, 1, '', 'sys/role.html', 'fa fa-group', '_self', 1, 2),
       (1001, '角色列表', 1000, 2, 'GET', '/api/role', '', '', 1, 2),
       (1002, '添加角色', 1000, 2, 'POST', '/api/role', '', '', 2, 2),
       (1003, '编辑角色', 1000, 2, 'PUT', '/api/role/:id', '', '', 3, 2),
       (1004, '删除角色', 1000, 2, 'DELETE', '/api/role/:id', '', '', 4, 2),
       (1005, '查看角色', 1000, 2, 'GET', '/api/role/:id', '', '', 5, 2),
       (1006, '查看授权', 1000, 2, 'GET', '/api/role/:id/auth', '', '', 6, 2),
       (1007, '设置授权', 1000, 2, 'PUT', '/api/role/:id/auth', '', '', 7, 2),

       (1100, '组织', 1, 1, '', 'sys/org.html', 'fa fa-sitemap', '_self', 1, 2),
       (1101, '组织列表', 1100, 2, 'GET', '/api/org', '', '', 1, 2),
       (1102, '添加组织', 1100, 2, 'POST', '/api/org', '', '', 2, 2),
       (1103, '编辑组织', 1100, 2, 'PUT', '/api/org/:id', '', '', 3, 2),
       (1104, '删除组织', 1100, 2, 'DELETE', '/api/org/:id', '', '', 4, 2),
       (1105, '查看组织', 1100, 2, 'GET', '/api/org/:id', '', '', 5, 2),


       (1200, '用户', 1, 1, '', 'sys/user.html', 'fa fa-user', '_self', 1, 2),
       (1201, '用户列表', 1200, 2, 'GET', '/api/user', '', '', 1, 2),
       (1202, '添加用户', 1200, 2, 'POST', '/api/user', '', '', 2, 2),
       (1203, '编辑用户', 1200, 2, 'PUT', '/api/user/:id', '', '', 3, 2),
       (1204, '删除用户', 1200, 2, 'DELETE', '/api/user/:id', '', '', 4, 2),
       (1205, '查看用户', 1200, 2, 'GET', '/api/user/:id', '', '', 5, 2),

       (1300, '日志', 1, 1, '', 'sys/log.html', 'fa fa-list', '_self', 1, 0),
       (1301, '日志列表', 1300, 2, 'GET', '/api/log', '', '', 1, 0),
       (2, '业务管理', 0, 1, '', '', '', '', 1, 2);
SELECT setval('sys_entry_id_seq', (select max(id) from sys_entry), true);
drop table if exists sys_auth;
create table sys_auth
(
    id       bigserial primary key,
    role_id  integer,
    entry_id integer
);


drop table if exists sys_user;
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
    delete_at timestamp without time zone,
    tenant_id int     not null
);

insert into sys_user(id, account, password, role_id, org_id, tenant_id)
values (1, 'admin', md5('abc123456'), 0, 0, 0),
       (2, 'zhanghua', md5('abc123456'), 1, 1, 1);
SELECT setval('sys_user_id_seq', (select max(id) from sys_user), true);
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
drop table if exists sys_file;
create table sys_file
(
    id        bigserial primary key,
    name      text,
    save_name text,
    save_path text,
    url       text,
    create_at timestamp without time zone default localtimestamp(0)
);

drop table if exists sys_log;
create table sys_log
(
    id        bigserial primary key,
    user_id   integer,
    uri       text,
    data      text,
    ip        text,
    ua        text,
    method    text,
    tenant_id int not null,
    create_at timestamp without time zone default localtimestamp(0)
);

drop table if exists sys_config;
create table sys_config
(
    id   bigserial primary key,
    key  text,
    data text
);


