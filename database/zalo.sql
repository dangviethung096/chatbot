drop table if exists zalo_messages;

drop table if exists zalo_sessions;

-- Create table zalo_message
create table if not exists zalo_messages (
    id SERIAL primary key,
    message_id VARCHAR(255),
    sender_id VARCHAR(255),
    recipient_id VARCHAR(255),
    message TEXT,
    thread_id VARCHAR(255),
    index INTEGER,
    created_at TIMESTAMP default current_timestamp
);

-- Create table zalo_session
create table if not exists zalo_sessions (
    id SERIAL primary key,
    sender_id VARCHAR(255),
    thread_id VARCHAR(255),
    updated_at TIMESTAMP default current_timestamp,
    created_at TIMESTAMP default current_timestamp
);