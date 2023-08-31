create schema if not exists avito;

create table avito.segments
(
    segment         text not null
        primary key,
    user_percentage numeric(5, 2) default 0.0
        constraint segments_user_percentage_check
            check ((user_percentage >= 0.0) AND (user_percentage <= 100.0))
);

create table avito.user_segment
(
    user_id    bigint                                 not null,
    segment    text                                   not null,
    created_at timestamp with time zone default now() not null,
    deleted_at timestamp with time zone,
    expired_at timestamp with time zone,
    primary key (segment, user_id)
);
