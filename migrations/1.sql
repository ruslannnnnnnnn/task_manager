create table tasks (
    id serial constraint tasks_pk primary key,
    title varchar(255) not null,
    description text,
    created_at timestamp with time zone default now() not null
);