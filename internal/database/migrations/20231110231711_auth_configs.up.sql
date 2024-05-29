
--pg:table
create table if not exists store_oauth_configs(
    id uuid primary key default uuid_generate_v4(),
    store_id uuid not null references stores(id) on delete cascade,
    auth_type text not null,
    credentials text not null,
    is_active boolean not null,
    created_at timestamp with time zone default (current_timestamp at time zone 'utc')
);



-- pg:table
create table if not exists order_shipments (
    id uuid primary key default uuid_generate_v4(),
    order_id bigint not null references orders(id),
    delivery_method_id uuid not null references store_delivery_methods(id),
    tracking_number varchar(255),
    estimated_delivery_date timestamp,
    shipped_at timestamp,
    price float default 0,
    status text default 'intransit' check (status in ('intransit', 'delivered', 'delayed', 'lost')),
    created_at timestamp with time zone default (current_timestamp at time zone 'utc'),
    updated_at timestamp with time zone default (current_timestamp at time zone 'utc')
);

