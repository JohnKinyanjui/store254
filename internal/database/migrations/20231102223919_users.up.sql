--pg:split
create extension if not exists "uuid-ossp";

--pg:split
create extension if not exists "citext";

--pg:split
create table if not exists users(
    id uuid primary key default uuid_generate_v4(),
    profile_image text not null,
    full_name text not null,
    verified boolean default 'f',
    status varchar(15) default 'ok' check (status in ('ok', 'unverified','suspended', 'banned')),
    address jsonb
);

--pg:split
create table if not exists user_auths(
    user_id uuid unique not null references users(id) on delete cascade,
    google_uid text unique,
    github_uid text unique,
    phone_number text unique,
    password text default 'none',
    email text unique,
    roles text[],
    credentials text
);

--pg:table
create table if not exists user_waitlists(
    id uuid primary key default uuid_generate_v4(),
    email citext not null,
    approved boolean default 'f',
    created_at timestamp not null default timezone('utc', now())
);