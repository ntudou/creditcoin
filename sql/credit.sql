alter user work with encrypted password '5J7zqnxY8HBIa00_5zJg5AlLOsdvODpn';

create user credit_user_u with ENCRYPTED password '5J7zqm59fHeho00_r35ger1jwnWoQvDA';
CREATE DATABASE credit_user WITH OWNER credit_user_u ENCODING UTF8 TEMPLATE template0;
\c credit_user
CREATE SCHEMA user_info;
ALTER SCHEMA user_info OWNER to credit_user_u;
revoke create on schema public from public;

CREATE TYPE e_user_status AS ENUM ('low','normal','lock');

CREATE TABLE t_user_base(
user_id bigint NOT NULL ,
user_name VARCHAR(256) NOT NULL ,
nike_name VARCHAR(256),
crc_code VARCHAR(512) NOT NULL ,
pwd VARCHAR(512) NOT NULL ,
tel varchar(512) NOT NULL ,
user_code varchar(512) NOT NULL ,
user_status e_user_status NOT NULL ,
create_user bigint NOT NULL ,
create_datetime TIMESTAMP without time zone NOT NULL ,
PRIMARY KEY (user_id),
UNIQUE (user_name)
);



usercoin
create user user_coin_u with ENCRYPTED password '5J7zqm59fHeho00_r35ger1jwnWoQvDA';
CREATE DATABASE user_coin WITH OWNER user_coin_u ENCODING UTF8 TEMPLATE template0;
\c user_coin
CREATE SCHEMA user_coin;
ALTER SCHEMA user_coin OWNER to user_coin_u;
revoke create on schema public from public;

CREATE TABLE t_user_coin(
user_id bigint NOT NULL ,
user_coin bigint NOT NULL DEFAULT 0,
crc_coin VARCHAR(512) NOT NULL '',
limited_coin bigint NOT NULL DEFAULT 0,
crc_limited VARCHAR(512) NOT NULL '',
update_datetime TIMESTAMP without time zone NOT NULL ,
PRIMARY KEY (user_id),
);

coinlog
create user coin_log_u with ENCRYPTED password '5J7zqm59fHeho00_r35ger1jwnWoQvDA';
CREATE DATABASE coin_log WITH OWNER coin_log_u ENCODING UTF8 TEMPLATE template0;
\c coin_log
CREATE SCHEMA coin_log;
ALTER SCHEMA coin_log OWNER to coin_log_u;
revoke create on schema public from public;

CREATE TYPE coin_log_type AS ENUM ('low','normal','lock');

CREATE TABLE t_coin_log(
user_id bigint NOT NULL ,
create_datetime TIMESTAMP without time zone NOT NULL ,
user_coin bigint NOT NULL DEFAULT 0,
log_info VARCHAR(2048) NOT NULL '',
PRIMARY KEY (user_id),
);