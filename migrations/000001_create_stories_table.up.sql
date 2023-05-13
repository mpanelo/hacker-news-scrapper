CREATE TABLE IF NOT EXISTS stories (
    id integer PRIMARY KEY,
    type text NOT NULL,
    by text NOT NULL,
    time timestamp(0) with time zone NOT NULL,
    kids integer[] NOT NULL,
    url text NOT NULL,
    score integer NOT NULL,
    title text NOT NULL,
    descendants integer NOT NULL
);
