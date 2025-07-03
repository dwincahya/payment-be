CREATE TABLE payment_methods (
 id serial4 NOT NULL,
 "name" varchar(50) NOT NULL,
 "desc" text NULL,
 order_num int4 DEFAULT 1 NOT NULL,
 user_action varchar(25) NOT NULL,
 created_at timestamptz NULL,
 updated_at timestamptz NULL,
 code varchar(25) NULL,
 CONSTRAINT payment_methods_name_unique UNIQUE (name),
 CONSTRAINT payment_methods_pkey PRIMARY KEY (id)
);