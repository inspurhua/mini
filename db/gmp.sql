--产品信息
drop table if exists sys_product;
create table sys_product
(
    id              bigserial primary key,
    code            text,--代号
    name            text,--名称中英文
    specs           text,--规格
    pack_specs      text,--包装规格
    batch           text,--批量
    batch_unit      text,--产量单位
    storage         text,--贮藏条件
    validity        text,--有效期
    standard        text,--执行标准
    approval_number text,--批准文号
    quality_name    text,--质量标准文件编号
    process_name    text,--生产工艺文件编号
    process_file    int,--工序描述文件,word表格
    status          int default 1,--状态|radio|1:正常,2:停产
    tenant_id       int not null
);
comment on COLUMN sys_product.tenant_id is '租户';


--物料类别
drop table if exists sys_material_type;
create table sys_material_type
(
    id        bigserial primary key,
    pid       int default 0,
    code      text,
    name      text,
    sort      int default 1,
    tenant_id int not null
);

insert into sys_material_type
values (1, 0, '100', '物料类别', 1, 1);
--物料
drop table if exists sys_material;
create table sys_material
(
    id          bigserial primary key,
    name        text,
    manufacture text,
    type_id     int,
    tenant_id   int not null
);

drop table if exists sys_quality;
create table sys_quality
(
    id        bigserial primary key,
    key_id    int,--可以是产品和物料,中间品
    --五项与key有关的字段
    key_t0    text,
    key_t1    text,
    key_t2    text,
    key_t3    text,
    key_t4    text,

    --20项文本数据
    data_t0   text,
    data_t1   text,
    data_t2   text,
    data_t3   text,
    data_t4   text,
    data_t5   text,
    data_t6   text,
    data_t7   text,
    data_t8   text,
    data_t9   text,
    data_t10  text,
    data_t11  text,
    data_t12  text,
    data_t13  text,
    data_t14  text,
    data_t15  text,
    data_t16  text,
    data_t17  text,
    data_t18  text,
    data_t19  text,

--20项整型数据
    data_i0   int,
    data_i1   int,
    data_i2   int,
    data_i3   int,
    data_i4   int,
    data_i5   int,
    data_i6   int,
    data_i7   int,
    data_i8   int,
    data_i9   int,
    data_i10  int,
    data_i11  int,
    data_i12  int,
    data_i13  int,
    data_i14  int,
    data_i15  int,
    data_i16  int,
    data_i17  int,
    data_i18  int,
    data_i19  int,

--20项小数数据
    data_n0   numeric,
    data_n1   numeric,
    data_n2   numeric,
    data_n3   numeric,
    data_n4   numeric,
    data_n5   numeric,
    data_n6   numeric,
    data_n7   numeric,
    data_n8   numeric,
    data_n9   numeric,
    data_n10  numeric,
    data_n11  numeric,
    data_n12  numeric,
    data_n13  numeric,
    data_n14  numeric,
    data_n15  numeric,
    data_n16  numeric,
    data_n17  numeric,
    data_n18  numeric,
    data_n19  numeric,

    tenant_id int not null
);
drop table if exists sys_quality_chart;
create table sys_quality_chart
(
    id         bigserial primary key,
    key_id     int,--可以是产品和物料,中间品
    quality_id int,--数据id
    col_name   text,--列名比如data_n3
    chart_file text,--图标图片路径
    tenant_id  int not null
);
--为产品分配指标
drop table if exists sys_quality_info;
create table sys_quality_info
(
    id           bigserial primary key,
    key_id       int,--物品信息
    col_name     text,--列名
    col_title    text,--中文显示
    group_title  text default '',--组名非空且相同,合并单元格
    refer_text   text,--参数文本显示
    refer_expr   text,--左闭右开区间表示:3,(>=3);,3(<3);3,5(>=3 <5),可以自动解析text 生成
    refer_unit   text,--单位
    zero_value   text default '未检出',
    chart_column text,--chart水平轴
    chart_type   int  default 1,--1.趋势图
    sort         int  default 1,
    tenant_id    int not null
);
--以上生成的sql需要缓存,修改配置的时候直接生成即可.
--前端自动匹配data_n2和他的中文



