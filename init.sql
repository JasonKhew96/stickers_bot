-- sticker_type
--   document (maybe)
--   gif
--   mpeg4gif
--   photo
--   sticker
--   video

CREATE TABLE IF NOT EXISTS sticker (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    file_id TEXT UNIQUE NOT NULL,
    sticker_type TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS keyword (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    keyword TEXT UNIQUE NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS sticker_keyword (
    sticker_id INTEGER NOT NULL,
    keyword_id INTEGER NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    PRIMARY KEY (sticker_id, keyword_id),
    FOREIGN KEY (sticker_id) REFERENCES sticker (id) ON DELETE CASCADE,
    FOREIGN KEY (keyword_id) REFERENCES keyword (id) ON DELETE CASCADE
);

