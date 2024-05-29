--pg:table
create table if not exists products(
    id uuid primary key default uuid_generate_v4(),
    store_id uuid not null references stores(id) on delete cascade,
    images text[],
    name text not null,
    description text not null,
    regular_price decimal(10,2) not null,
    sales_price decimal(10,2) not null default 0,
    rating decimal(10,1) not null default 4.0,
    active boolean default 'f',
    downloadable boolean default 'f',
    reference_note text,
    attachments jsonb[],
    deleted bool not null default 'f',
    deleted_at timestamp,
    created_at timestamp not null default timezone('utc', now())
);

--pg:index
create index idx_products_store_id on products(store_id);
create index idx_products_name on products(name);
create index idx_products_active_deleted on products(deleted);
create index idx_products_created_at on products(created_at);

--pg:table
create table if not exists product_variants (
    id uuid primary key default uuid_generate_v4(),
    product_id uuid not null references products(id) on delete cascade,
    name text not null,
    price decimal(10,2) not null default 0,
    stock_level decimal(10,2) not null default 0,
    attributes jsonb, 
    created_at timestamp not null default timezone('utc', now())
);

--pg:index
create index idx_product_variants_product_id on product_variants(product_id);

--pg:table
create table if not exists product_categories(
    id serial primary key,
    store_id uuid not null references stores(id) on delete cascade,
    category_name text not null,
    slug text not null,
    parent_category_id int references product_categories(id) on delete set null
);

--pg:index
create index idx_pc_slug on product_categories(slug);
create index idx_pc_parent_category_id on product_categories(parent_category_id);
create index idx_pc_store_id on product_categories(store_id);

--pg:table
create table if not exists product_rel_categories(
    product_id uuid not null references products(id) on delete cascade,
    category_id integer not null references product_categories(id) on delete cascade
);

--pg:index
create index idx_prc_product_id on product_rel_categories(product_id);
create index idx_prc_category_id on product_rel_categories(category_id);

--pg:table
create table if not exists product_tags(
    id serial primary key,
    store_id uuid not null references stores(id) on delete cascade,
    tags_name text not null
);

--pg:table
create table if not exists product_rel_tags(
    product_id uuid references products(id) on delete cascade,
    tag_id integer references product_tags(id) on delete cascade
);

