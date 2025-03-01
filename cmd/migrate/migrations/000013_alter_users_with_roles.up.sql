ALTER TABLE IF EXISTS users
ADD COLUMN role_id int references roles(id) default 1;

UPDATE users SET role_id = (
    SELECT id from roles where name = 'user'
);

ALTER TABLE users ALTER COLUMN role_id drop default;

ALTER TABLE users ALTER COLUMN role_id SET NOT NULL;
