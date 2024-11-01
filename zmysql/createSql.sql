-- 用户信息表
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


-- 质押池表
create table sc_stake_pool
(
    id                      bigint unsigned auto_increment comment '主键'
        primary key,
    pool_id                 bigint       default -1                not null comment '质押池id',
    pool_code               varchar(128) default ''                not null comment '质押池唯一code',
    token_address           varchar(128) default ''                not null comment 'token地址',
    lock_stake_block_number bigint       default 0                 not null comment '解锁质押的区块数',
    min_stake_amount        bigint(64)   default 0                 not null comment '最小质押金额,单位：wei',
    wight                   bigint       default 0                 not null comment '池权重',
    creator                 varchar(64)  default ''                not null comment '创建人',
    is_deleted              tinyint      default 0                 not null comment '是否已删除；1=是，0=否',
    mtime                   datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    ctime                   datetime     default CURRENT_TIMESTAMP not null comment '创建时间'
)
    comment '质押池表';

create index ix_pool_id
    on sc_stake_pool (pool_id);

create index ix_token_address
    on sc_stake_pool (token_address);


-- 质押交易记录表
create table sc_transaction_logs
(
    id               bigint unsigned auto_increment comment '主键'
        primary key,
    trans_type       varchar(64)  default ''                not null comment '交易类型',
    trans_hash       varchar(128) default ''                not null comment '交易hash',
    gas              bigint       default 0                 not null comment '消耗的gas',
    gas_price        bigint       default 0                 not null comment 'gas价格',
    trans_status     tinyint      default 0                 not null comment '交易状态',
    trans_logs       varchar(512) default ''                not null comment '交易日志',
    trans_cost       bigint       default 0                 not null comment '总的交易的费用',
    trans_from       varchar(128) default ''                not null comment '交易from地址',
    trans_to        varchar(128) default ''                not null comment '交易to地址',
    creator          varchar(64)  default ''                not null comment '创建人',
    is_deleted       tinyint      default 0                 not null comment '是否已删除；1=是，0=否',
    mtime            datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    ctime            datetime     default CURRENT_TIMESTAMP not null comment '创建时间'
)
    comment '质押交易记录表';

create index ix_pool_id
    on sc_transaction_logs (trans_hash);