CREATE TABLE IF NOT EXISTS carts
(
    id               INT GENERATED BY DEFAULT AS IDENTITY,
    guid             UUID      default uuid_generate_v4()                             NOT NULL,
    user_id          INT REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE    NOT NULL,
    product_id       INT REFERENCES products (id) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
    qty              INT                                                              NOT NULL,
    created_at       TIMESTAMP default now(),
    last_modified_at TIMESTAMP,
    deleted_at       TIMESTAMP,

    PRIMARY KEY (id),
    UNIQUE (guid)
);
