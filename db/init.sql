CREATE TABLE public.restaurant (
    id serial4 NOT NULL,
    "name" text NULL,
    address text NULL,
    point int4 NULL,
    CONSTRAINT restaurant_pkey PRIMARY KEY (id)
);

CREATE TABLE public.users (
    id serial4 NOT NULL,
    "name" varchar(30) NULL,
    email text NOT NULL,
    "type" int4 NOT NULL DEFAULT 1,
    "password" text NOT NULL,
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_pkey PRIMARY KEY (id)
);
