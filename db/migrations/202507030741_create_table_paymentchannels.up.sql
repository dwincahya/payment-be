CREATE TABLE payment_channels (
 id serial4 NOT NULL,
 payment_method_id int4 NULL,
 code varchar(255) NOT NULL,
 "name" varchar(50) NOT NULL,
 icon_url varchar(255) NULL,
 order_num int4 DEFAULT 1 NULL,
 lib_name varchar(255) NULL,
 user_action varchar(25) NOT NULL,
 created_at timestamptz NULL,
 updated_at timestamptz NULL,
 mdr varchar(255) DEFAULT '0'::character varying NULL,
 fixed_fee numeric DEFAULT 0 NULL,
 CONSTRAINT payment_channels_code_unique UNIQUE (code),
 CONSTRAINT payment_channels_name_unique UNIQUE (name),
 CONSTRAINT payment_channels_pkey PRIMARY KEY (id)
);

ALTER TABLE payment_channels ADD CONSTRAINT payment_channels_payment_method_id_foreign FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id);