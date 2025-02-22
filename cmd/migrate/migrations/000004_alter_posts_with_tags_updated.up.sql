ALTER TABLE posts ADD COLUMN tags varchar(128) [];
ALTER TABLE posts ADD COLUMN updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW();