create table sc_user
(
    id                  bigint unsigned auto_increment comment '主键'
        primary key,
    user_wallet_address varchar(128) default ''                not null comment '用户钱包地址',
    user_UID            varchar(128) default ''                not null comment '用户UID',
    user_telephone      varchar(32)  default ''                not null comment '用户电话',
    user_name           varchar(64)  default ''                not null comment '用户名称',
    user_password       varchar(64)  default ''                not null comment '用户密码',
    user_email          varchar(64)  default ''                not null comment '用户邮箱',
    user_country        varchar(128) default ''                not null comment '用户国家',
    user_address        varchar(128) default ''                not null comment '用户住址',
    remark              varchar(512) default ''                not null comment '备注信息',
    is_deleted          tinyint      default 0                 not null comment '是否已删除；1=是，0=否',
    mtime               datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    ctime               datetime     default CURRENT_TIMESTAMP not null comment '创建时间'
)
    comment '用户信息表';

create index ix_user_wallet_address
    on sc_user (user_wallet_address);

create index ix_user_telephone
    on sc_user (user_telephone);