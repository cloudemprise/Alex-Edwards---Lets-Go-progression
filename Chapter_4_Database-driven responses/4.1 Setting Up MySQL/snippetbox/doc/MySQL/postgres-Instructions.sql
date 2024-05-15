# Use this to list all the roles:
SELECT rolname FROM pg_roles;\

# Display all tables
\l

#---

# Superuser access to client:
>:$ sudo -u postgres psql

#CREATE ROLE solar WITH LOGIN PASSWORD 'abcd1234';
CREATE ROLE web WITH LOGIN PASSWORD 'abcd1234';
CREATE DATABASE snippetbox;
GRANT ALL PRIVILEGES ON DATABASE snippetbox TO solar;

#---

>:$ psql -U solar snippetbox


-- Create a `snippets` table.
CREATE TABLE snippets (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMP NOT NULL,
    expires TIMESTAMP NOT NULL
);


-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);


-- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO snippets (title, content, created, expires) VALUES (
    'An old silent pond',
    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '365 DAY'
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'Over the wintry forest',
    'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '365 DAY'
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'First autumn morning',
    'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '7 DAY'
);


# To validate table:
SELECT id, title, expires FROM snippets;