--pg:table
create table if not exists stores(
    id uuid primary key default uuid_generate_v4(),
    owner_id uuid not null references users(id) on delete cascade, 
    image text,
    name text not null,
    verified boolean default 'f',
    rating decimal(10,1) not null default 4.0,
    followers bigint not null default 0,
    reviews bigint not null default 0,
    country text not null default 'KE',
    currency text not null default 'KES',
    status text not null default 'closed' check (status in ('open', 'closed', 'banned')),
    contact_info jsonb default '{}',
    created_at timestamp not null default timezone('utc', now())
);


--pg:table
create table if not exists store_customers(
    id uuid primary key default uuid_generate_v4(),
    store_id uuid not null references stores(id) on delete cascade,
    profile_image text,
    name text not null,
    created_at timestamp not null default timezone('utc', now())
);

--pg:table
create table if not exists store_customer_auths(
    customer_id uuid primary key references store_customers(id) on delete cascade,
    email citext unique,
    phone_number text,
    password text,
    google_uid text unique,
    facebook_uid text unique
);

--pg:tables
create table if not exists store_delivery_methods (
    id uuid primary key default uuid_generate_v4(),
    store_id uuid not null references stores(id) on delete cascade,
    name varchar(255) not null,
    description text,
    tag text, 
    estimated_duration bigint default 0,
    price float default 0,
    credentials text not null,
    is_active boolean default true,
    created_at timestamp with time zone default (current_timestamp at time zone 'utc'),
    updated_at timestamp with time zone default (current_timestamp at time zone 'utc')
);


--pg:table
create table if not exists store_customer_addresses(
    customer_id uuid primary key references store_customers(id) on delete cascade,
    street text not null,
    city text not null,
    state text not null,
    postal_code text,
    country text not null
);

--pg:table
create table if not exists store_apis(
    id uuid primary key default uuid_generate_v4(),
    store_id uuid not null references stores(id) on delete cascade,
    label text not null,
    token text not null,
    created_at timestamp not null default timezone('utc', now())
);

--pg:table
create table if not exists store_payment_methods (
    id uuid primary key default uuid_generate_v4(),
    store_id uuid not null references stores(id) on delete cascade,
    name varchar(255) not null,
    tag text, 
    description text,
    credentials jsonb not null,
    is_active boolean default true,
    created_at timestamp with time zone default (current_timestamp at time zone 'utc'),
    updated_at timestamp with time zone default (current_timestamp at time zone 'utc')
);

--pg:index
create index idx_store_payment_methods_store_id on store_payment_methods(store_id);
create index idx_store_payment_methods_name on store_payment_methods(name);


