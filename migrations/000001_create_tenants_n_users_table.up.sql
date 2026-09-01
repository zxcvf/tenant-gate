-- public.tenants
CREATE TABLE public.tenants (
	id BIGINT NOT NULL,
	tenant_name varchar NOT NULL UNIQUE,
	email varchar NOT NULL UNIQUE,
	created_by varchar NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT tenants_pk PRIMARY KEY (id)
);


-- public.users
CREATE TABLE public.users (
	id BIGINT NOT NULL,
	username varchar NOT NULL UNIQUE,
	email varchar NOT NULL UNIQUE,
	phone varchar NOT NULL UNIQUE,
	password_hash varchar NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
);


-- public.tenants_users
CREATE TABLE public.tenants_users (
	id int NOT NULL,
	tenant_id BIGINT NOT NULL,
	user_id BIGINT NOT NULL,
	role_code int NOT NULL,
	CONSTRAINT tenants_users_pk PRIMARY KEY (id),
	CONSTRAINT tenants_users_unique UNIQUE (tenant_id,user_id)
);