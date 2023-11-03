package main

import (
	"context"
	"database/sql"

	"github.com/JasonKhew96/stickers_bot/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	_ "modernc.org/sqlite"
)

type Database struct {
	db  *sql.DB
	ctx context.Context
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("sqlite", "./data.db?_pragma=foreign_keys(1)")
	if err != nil {
		return nil, err
	}
	return &Database{db: db, ctx: context.TODO()}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetKeywordsFromFileId(fileId string) (models.KeywordSlice, error) {
	return models.Keywords(qm.InnerJoin("sticker_keyword on sticker_keyword.keyword_id = keyword.id"), qm.InnerJoin("sticker on sticker.id = sticker_keyword.sticker_id"), models.StickerWhere.FileID.EQ(fileId)).All(d.ctx, d.db)
}

func (d *Database) GetStickerFromFileId(fileId string) (*models.Sticker, error) {
	return models.Stickers(models.StickerWhere.FileID.EQ(fileId)).One(d.ctx, d.db)
}

func (d *Database) GetStickersFromKeyword(keyword string) (models.StickerSlice, error) {
	return models.Stickers(qm.InnerJoin("sticker_keyword on sticker_keyword.sticker_id = sticker.id"), qm.InnerJoin("keyword on keyword.id = sticker_keyword.keyword_id"), models.KeywordWhere.Keyword.EQ(keyword), qm.OrderBy("updated_at DESC"), qm.Limit(50)).All(d.ctx, d.db)
}

func (d *Database) SaveSticker(fileId, stickerType string, keywords []string) error {
	s := models.Sticker{
		FileID:      fileId,
		StickerType: stickerType,
	}
	err := s.Upsert(d.ctx, d.db, true, []string{models.StickerColumns.FileID}, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}
	for _, keyword := range keywords {
		k := models.Keyword{
			Keyword: keyword,
		}
		err = k.Upsert(d.ctx, d.db, true, []string{models.KeywordColumns.Keyword}, boil.Infer(), boil.Infer())
		if err != nil {
			return err
		}
		if s.ID == 0 || k.ID == 0 {
			continue
		}
		sk := models.StickerKeyword{
			StickerID: s.ID,
			KeywordID: k.ID,
		}
		err = sk.Upsert(d.ctx, d.db, true, []string{models.StickerKeywordColumns.StickerID, models.StickerKeywordColumns.KeywordID}, boil.Infer(), boil.Infer())
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) RemoveSticker(fileId string) error {
	// todo: remove any keywords that are no longer used
	_, err := models.Stickers(models.StickerWhere.FileID.EQ(fileId)).DeleteAll(d.ctx, d.db)
	return err
}

func (d *Database) UpdateStickerUsage(id int64) error {
	sticker, err := models.Stickers(models.StickerWhere.ID.EQ(id)).One(d.ctx, d.db)
	if err != nil {
		return err
	}
	_, err = sticker.Update(d.ctx, d.db, boil.Infer())
	return err
}
