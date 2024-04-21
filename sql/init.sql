CREATE TABLE employee (
    id serial,
    name text,
    surname text
);

CREATE TABLE attendance (
    id serial,
    employee_id int,
    time timestamp,
    work_event_type varchar(1)
);

CREATE TABLE salary (
    id serial,
    employee_id int,
    salary_per_hour int
);
