drop table if exists facebook_messages cascade;

drop table if exists facebook_sessions cascade;

create table if not exists facebook_sessions(
    id bigserial primary key,
    sender text not null,
    thread_id text not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table if not exists facebook_messages(
    id bigserial primary key,
    sender text not null,
    receiver text not null,
    message text not null,
    thread_id text not null,
    created_at timestamp default now()
);