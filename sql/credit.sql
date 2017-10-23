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



creditcoin