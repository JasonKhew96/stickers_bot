package main

import (
	"context"

	"github.com/JasonKhew96/stickers_bot/models"
	"github.com/aarondl/opt/omit"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/sqlite/im"
	"github.com/stephenafamo/bob/dialect/sqlite/sm"
	_ "modernc.org/sqlite"
)

/*
CREATE TABLE IF NOT EXISTS sticker (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    file_id TEXT UNIQUE NOT NULL,
    sticker_type TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS keyword (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
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
	s, err := models.Stickers.Insert(&models.StickerSetter{
		FileID:      omit.From(fileId),
		StickerType: omit.From(stickerType),
	}, im.OnConflict("file_id").DoNothing()).One(d.ctx, d.db)
	if err != nil {
		return err
	}
	for _, keyword := range keywords {
		k, err := models.Keywords.Insert(&models.KeywordSetter{
			Keyword: omit.From(keyword),
		}, im.OnConflict("keyword").DoNothing()).One(d.ctx, d.db)
		if err != nil {
			return err
		}
		if s.ID == 0 || k.ID == 0 {
			continue
		}
		_, err = models.StickerKeywords.Insert(&models.StickerKeywordSetter{
			StickerID: omit.From(s.ID),
			KeywordID: omit.From(k.ID),
		}, im.OnConflict("sticker_id", "keyword_id").DoNothing()).One(d.ctx, d.db)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) RemoveSticker(fileId string) error {
	// todo: remove any keywords that are no longer used
	_, err := models.Stickers.Delete(models.DeleteWhere.Stickers.FileID.EQ(fileId)).All(d.ctx, d.db)
	return err
}
