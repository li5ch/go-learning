create table link_tab
(
    id           bigint unsigned auto_increment
        primary key,
    link         varchar(1024)           not null comment 'affiliate link'
)