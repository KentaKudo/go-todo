
-- +migrate Up
CREATE TABLE IF NOT EXISTS todos (
    id    INTEGER AUTO_INCREMENT NOT NULL,
    title TEXT                   NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB;

-- +migrate Down
DROP TABLE todos;