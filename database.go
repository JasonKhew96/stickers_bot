package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/JasonKhew96/stickers_bot/models"
	"github.com/aarondl/opt/omit"
	"github.com/pkg/errors"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/sqlite/im"
	"github.com/stephenafamo/bob/dialect/sqlite/sm"
	_ "modernc.org/sqlite"
)

/*
CREATE TABLE IF NOT EXISTS sticker (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    file_id TEXT UNIQUE NOT NULL,
    sticker_type TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS keyword (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    keyword TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS sticker_keyword (
    sticker_id INTEGER NOT NULL,
    keyword_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (sticker_id, keyword_id),
    FOREIGN KEY (sticker_id) REFERENCES sticker (id) ON DELETE CASCADE,
    FOREIGN KEY (keyword_id) REFERENCES keyword (id) ON DELETE CASCADE
);
*/

type Database struct {
	db  bob.DB
	ctx context.Context
}

func NewDatabase() (*Database, error) {
	bobExec, err := bob.Open("sqlite", "./data.db?_pragma=foreign_keys(1)&cache=shared")
	if err != nil {
		return nil, err
	}
	return &Database{db: bobExec, ctx: context.TODO()}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetKeywordsFromFileId(fileId string) (models.KeywordSlice, error) {
	return models.Keywords.Query(sm.InnerJoin("sticker_keyword on sticker_keyword.keyword_id = keyword.id"), sm.InnerJoin("sticker on sticker.id = sticker_keyword.sticker_id"), models.SelectWhere.Stickers.FileID.EQ(fileId)).All(d.ctx, d.db)
}

func (d *Database) GetStickerFromFileId(fileId string) (*models.Sticker, error) {
	return models.Stickers.Query(models.SelectWhere.Stickers.FileID.EQ(fileId)).One(d.ctx, d.db)
}

func (d *Database) GetStickersFromKeyword(keyword string) (models.StickerSlice, error) {
	return models.Stickers.Query(sm.InnerJoin("sticker_keyword on sticker_keyword.sticker_id = sticker.id"), sm.InnerJoin("keyword on keyword.id = sticker_keyword.keyword_id"), models.SelectWhere.Keywords.Keyword.EQ(keyword), sm.OrderBy("updated_at DESC"), sm.Limit(50)).All(d.ctx, d.db)
}

func (d *Database) SaveSticker(fileId, stickerType string, keywords []string) error {
	s, err := d.GetStickerFromFileId(fileId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return errors.Wrap(err, "get sticker failed")
	}
	if err == sql.ErrNoRows {
		s, err = models.Stickers.Insert(&models.StickerSetter{
			ID:          omit.From(s.ID),
			StickerType: omit.From(stickerType),
			UpdatedAt:   omit.From(time.Now()),
		}, im.OnConflict("file_id").DoUpdate(im.SetExcluded("updated_at"))).One(d.ctx, d.db)
		if err != nil {
			return errors.Wrap(err, "insert sticker failed")
		}
	} else {
		_, err = models.StickerKeywords.Delete(models.DeleteWhere.StickerKeywords.StickerID.EQ(s.ID)).Exec(d.ctx, d.db)
		if err != nil {
			return errors.Wrap(err, "delete sticker keyword failed")
		}
	}

	for _, keyword := range keywords {
		k, err := models.Keywords.Insert(&models.KeywordSetter{
			Keyword:   omit.From(keyword),
			UpdatedAt: omit.From(time.Now()),
		}, im.OnConflict("keyword").DoUpdate(im.SetExcluded("updated_at"))).One(d.ctx, d.db)
		if err != nil {
			return errors.Wrap(err, "insert keyword failed")
		}
		if s.ID == 0 || k.ID == 0 {
			continue
		}
		_, err = models.StickerKeywords.Insert(&models.StickerKeywordSetter{
			StickerID: omit.From(s.ID),
			KeywordID: omit.From(k.ID),
			UpdatedAt: omit.From(time.Now()),
		}, im.OnConflict("sticker_id", "keyword_id").DoUpdate(im.SetExcluded("updated_at"))).One(d.ctx, d.db)
		if err != nil {
			return errors.Wrap(err, "insert sticker keyword failed")
		}
	}
	return nil
}

func (d *Database) RemoveSticker(fileId string) error {
	s, err := d.GetStickerFromFileId(fileId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return errors.Wrap(err, "get sticker failed")
	}
	_, err = models.StickerKeywords.Delete(models.DeleteWhere.StickerKeywords.StickerID.EQ(s.ID)).Exec(d.ctx, d.db)
	if err != nil {
		return errors.Wrap(err, "delete sticker keyword failed")
	}
	return s.Delete(d.ctx, d.db)
}
