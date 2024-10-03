-- create user 'chatbot'
create user chatbot with password 'chatbot';

-- create database 'chatbot'
create database chatbot with owner chatbot;

-- create schema 'chatbot'
create schema chatbot authorization chatbot;