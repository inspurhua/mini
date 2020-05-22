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

--产品信息
drop table if exists sys_material;
create table sys_material
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
    manufacture     text,--生产商
    type_id         int,--物料类型
    tenant_id       int not null
);
comment on COLUMN sys_material.tenant_id is '租户';

--物料和参数体系

drop table if exists sys_quality;
create table sys_quality
(
    id        bigserial primary key,
    key_id    int,--可以是产品和物料,中间品
    batch     text,--默认是批次
    --五项与key有关的字段
    key_t00   text,
    key_t01   text,
    key_t02   text,
    key_t03   text,
    key_t04   text,

    --20项文本数据
    data_t00  text,
    data_t01  text,
    data_t02  text,
    data_t03  text,
    data_t04  text,
    data_t05  text,
    data_t06  text,
    data_t07  text,
    data_t08  text,
    data_t09  text,
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
    data_i00  int,
    data_i01  int,
    data_i02  int,
    data_i03  int,
    data_i04  int,
    data_i05  int,
    data_i06  int,
    data_i07  int,
    data_i08  int,
    data_i09  int,
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
    data_n00  numeric,
    data_n01  numeric,
    data_n02  numeric,
    data_n03  numeric,
    data_n04  numeric,
    data_n05  numeric,
    data_n06  numeric,
    data_n07  numeric,
    data_n08  numeric,
    data_n09  numeric,
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
    qc_date   date,--质检日期
    create_at timestamp without time zone default localtimestamp(0),
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
    id          bigserial primary key,
    key_id      int,--物品信息
    col_name    text,--列名
    col_title   text,--中文显示
    group_title text default '',--组名非空且相同,合并单元格
    refer_text  text,--参数文本显示
    refer_expr  text,--左闭右开区间表示:3,(>=3);,3(<3);3,5(>=3 <5),可以自动解析text 生成
    refer_unit  text,--单位
    sort        int  default 1,
    tenant_id   int not null
);
--以上生成的sql需要缓存,修改配置的时候直接生成即可.
--前端自动匹配data_n2和他的中文



