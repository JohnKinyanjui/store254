--pg:table
create table if not exists templates(
    id uuid primary key default uuid_generate_v4(),
    user_id uuid unique not null references users(id) on delete cascade,
    images text[] not null,
    title text not null,
    description text,
    regular_price decimal(10,2) not null default 0,
    sales_price decimal(10,2) not null default 0,
    source text not null,
    tag text[],
    created_at timestamp with time zone default (current_timestamp at time zone 'utc'),
    updated_at timestamp with time zone default (current_timestamp at time zone 'utc')
);

--pg:table
create table if not exists store_deployments(
    id uuid primary key default uuid_generate_v4(),
    template_id uuid not null references templates(id),
    deployment_id text not null,
    template_tag text not null,
    enviroments jsonb
);

--pg:table
create table if not exists subscriptions(
    id uuid primary key default uuid_generate_v4(),
    user_id uuid unique not null references users(id) on delete cascade,
    cost decimal(10,2) not null default 0,
    currency varchar(3) not null,
    subscription_type text not null default 'trial' check (subscription_type in ('triel', 'subscription')),
    subscription_info_type text not null,
    created_at timestamp with time zone default (current_timestamp at time zone 'utc'),
    expiry_at timestamp with time zone default (current_timestamp at time zone 'utc')
);

--pg:table
create table if not exists subscription_invoices(
    id uuid primary key default uuid_generate_v4(),
    user_id uuid not null references users(id) on delete cascade,
    cost decimal(10,2) not null default 0,
    currency varchar(3) not null,
    subscription_type text not null default 'trial' check (subscription_type in ('triel', 'subscription')),
    subscription_info_type text not null,
    created_at timestamp with time zone default (current_timestamp at time zone 'utc'),
    expiry_at timestamp with time zone default (current_timestamp at time zone 'utc')
);