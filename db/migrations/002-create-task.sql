CREATE TYPE task_status AS ENUM('todo', 'doing', 'done', 'canceled');


CREATE TABLE task (
    id serial PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    created_by INTEGER NOT NULL REFERENCES "users" (id),
    status task_status NOT NULL,
    due_date date NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
