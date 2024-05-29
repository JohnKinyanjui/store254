--pg:table
create table if not exists carts (
    id uuid primary key default uuid_generate_v4(),
    store_id uuid not null references stores(id) on delete cascade,
    customer_id uuid references store_customers(id) on delete set null, 
    created_at timestamp not null default timezone('utc', now()),
    updated_at timestamp
);

--pg:index
create index idx_carts_customer_id on carts(customer_id);
create index idx_carts_store_id on carts(store_id);

--pg:table
create table if not exists cart_items (
    id uuid primary key default uuid_generate_v4(),
    cart_id uuid not null references carts(id) on delete cascade,
    product_id uuid not null references products(id) on delete cascade,
    variant_id uuid references product_variants(id) on delete set null,
    quantity integer not null,
    created_at timestamp not null default timezone('utc', now())
);

--pg:index
create index idx_cart_items_cart_id on cart_items(cart_id);
create index idx_cart_items_product_id on cart_items(product_id);
create index idx_cart_items_variant_id on cart_items(variant_id);

--pg:table 
create table if not exists customer_product_views (
    id uuid primary key default uuid_generate_v4(),
    customer_id uuid not null references store_customers(id) on delete cascade,
    product_id uuid not null references products(id),
    last_viewed timestamp not null,
    view_count integer not null default 1
);

--pg:index
create index idx_customer_product_views_customer_id on customer_product_views(customer_id);
create index idx_customer_product_views_product_id on customer_product_views(product_id);

--pg:table
CREATE TABLE IF NOT EXISTS product_views (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id uuid NOT NULL,
    view_count integer NOT NULL DEFAULT 1,
    last_viewed timestamp NOT NULL DEFAULT timezone('utc', now())
);

--pg: indexes
create index idx_product_views_product_id on product_views(product_id);
create index idx_product_views_last_viewed on product_views(last_viewed);

--pg:table
create table if not exists order_transactions (
    id serial primary key,
    order_id int not null references orders(id) on delete cascade,
    payment_method_id uuid not null references store_payment_methods(id) on delete cascade,
    payment_date timestamp,
    amount decimal(10, 2) not null,
    currency varchar(3),
    transaction_status varchar(50),
    transaction_reference_id varchar(255),
    metadata jsonb default '{}',
    created_at timestamp not null,
    updated_at timestamp with time zone default (current_timestamp at time zone 'utc')
);

--pg:index
create index idx_order_transactions_order_id on order_transactions(order_id);
create index idx_order_transactions_payment_date on order_transactions(payment_date);
create index idx_order_transactions_status on order_transactions(transaction_status);
create index idx_order_transactions_order_id_status on order_transactions(order_id, transaction_status);
