--pg:table
create table if not exists orders (
    id serial primary key,
    store_id uuid not null references stores(id) on delete cascade,
    customer_id uuid references store_customers(id) on delete set null, 
    payment_method_id uuid references store_payment_methods(id) on delete set null,
    delivery_method_id uuid references store_delivery_methods(id) on delete set null,
    delivery_cost decimal(10, 2) default 0,
    total_cost decimal(10, 2) not null,
    note text,
    store_note text,
    order_status text not null default 'pending' check (order_status in ('pending', 'processing', 'shipping', 'completed', 'cancelled')),
    paid boolean not null default false,
    billing jsonb not null,
    shipping jsonb,
    payment_date timestamp,
    created_at timestamp not null,
    processed_at timestamp,
    cancelled_at timestamp,
    completed_at timestamp
);

--pg:table
create table if not exists order_products (
    id serial primary key,
    order_id int not null references orders(id) on delete cascade,
    customer_id uuid references store_customers(id) on delete set null, 
    store_id uuid not null references stores(id) on delete cascade,
    product_id uuid not null references products(id) on delete cascade,
    quantity int not null,
    price decimal(10, 2) not null,
    total_price decimal(10, 2) not null,
    created_at timestamp not null
);

--pg:table
create table if not exists product_inventorys(
    id serial primary key,
    product_id uuid not null references products(id) on delete cascade,
    quantity int not null default 0,
    minimum_quantity int not null default 0, 
    status text not null default 'in_stock' check (status in ('in_stock','out_of_stock', 'backordered')),
    unlimited_stock boolean not null default 'f',
    updated_at timestamp not null,
    created_at timestamp not null
);

--pg:table
create table if not exists notifications(
    id serial primary key,
    sender_id uuid references users(id) on delete cascade,
    store_sender_id uuid references stores(id) on delete cascade,
    receipient_id uuid not null references users(id) on delete cascade,
    store_recipient_id uuid references stores(id) on delete cascade,
    title text not null,
    message varchar not null,
    status text not null default 'unread' check (status in ('read', 'unread')),
    created_at timestamp not null
);

