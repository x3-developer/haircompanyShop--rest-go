CREATE TABLE categories
(
    id                 SERIAL PRIMARY KEY,
    name               VARCHAR(255) NOT NULL UNIQUE,
    description        TEXT,
    image              VARCHAR(255),
    header_image       VARCHAR(255),
    slug               VARCHAR(255) NOT NULL UNIQUE,
    parent_id          INTEGER      REFERENCES categories (id) ON DELETE SET NULL,
    sort_index         INTEGER      NOT NULL DEFAULT 100,
    seo_title          VARCHAR(255),
    seo_description    TEXT,
    seo_keys           TEXT,
    is_active          BOOLEAN      NOT NULL DEFAULT TRUE,
    is_shade           BOOLEAN      NOT NULL DEFAULT FALSE,
    is_visible_in_menu BOOLEAN      NOT NULL DEFAULT TRUE,
    is_visible_on_main BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at         TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_categories_parent_id ON categories (parent_id);