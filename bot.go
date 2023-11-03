package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/choseninlineresult"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/inlinequery"
	"github.com/pkg/errors"
)

type StickerBot struct {
	config *Config
	bot    *gotgbot.Bot
	db     *Database
}

func NewBot(config *Config, db *Database) (*StickerBot, error) {
	bot, err := gotgbot.NewBot(config.BotToken, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create bot")
	}

	sb := &StickerBot{
		config: config,
		bot:    bot,
		db:     db,
	}

	updater := ext.NewUpdater(&ext.UpdaterOpts{
		Dispatcher: ext.NewDispatcher(&ext.DispatcherOpts{
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				log.Println("an error occurred while handling update: ", err.Error())
				return ext.DispatcherActionNoop
			},
			MaxRoutines: ext.DefaultMaxRoutines,
		}),
	})
	dispatcher := updater.Dispatcher
	dispatcher.AddHandler(handlers.NewCommand("save", sb.commandSave))
	dispatcher.AddHandler(handlers.NewCommand("remove", sb.commandRemove))
	dispatcher.AddHandler(handlers.NewInlineQuery(inlinequery.All, sb.inlineQuery))
	dispatcher.AddHandler(handlers.NewChosenInlineResult(choseninlineresult.All, sb.choosenInlineResult))

	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout:        60,
			AllowedUpdates: []string{"message", "inline_query", "chosen_inline_result"},
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 60,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	updater.Idle()

	return sb, err
}

func (sb *StickerBot) commandSave(b *gotgbot.Bot, ctx *ext.Context) error {
	if !ctx.EffectiveSender.IsUser() {
		return nil
	}
	userId := ctx.EffectiveSender.User.Id
	if userId != sb.config.OwnerId {
		return nil
	}
	keywords := strings.Split(ctx.EffectiveMessage.Text, " ")
	if len(keywords) <= 0 {
		_, err := ctx.EffectiveMessage.Reply(b, "Please provide keywords", nil)
		return err
	}
	keywords = keywords[1:]
	msg := ctx.EffectiveMessage.ReplyToMessage
	if msg == nil {
		_, err := ctx.EffectiveMessage.Reply(b, "Please reply to a sticker", nil)
		return err
	}
	var fileId string
	var stickerType string
	if msg.Sticker != nil {
		fileId = msg.Sticker.FileId
		stickerType = "sticker"
	} else if msg.Animation != nil {
		fileId = msg.Animation.FileId
		stickerType = "mpeg4gif"
	} else if msg.Photo != nil {
		fmt.Println(msg.Photo)
		fileId = msg.Photo[0].FileId
		stickerType = "photo"
	} else if msg.Video != nil {
		fileId = msg.Video.FileId
		stickerType = "video"
	}
	err := sb.db.SaveSticker(fileId, stickerType, keywords)
	if err != nil {
		log.Println(err)
		_, err := ctx.EffectiveMessage.Reply(b, "Failed to save sticker", nil)
		return err
	}
	_, err = ctx.EffectiveMessage.Reply(b, "Sticker saved", nil)
	return err
}

func (sb *StickerBot) commandRemove(b *gotgbot.Bot, ctx *ext.Context) error {
	if !ctx.EffectiveSender.IsUser() {
		return nil
	}
	userId := ctx.EffectiveSender.User.Id
	if userId != sb.config.OwnerId {
		return nil
	}
	msg := ctx.EffectiveMessage.ReplyToMessage
	if msg == nil {
		_, err := ctx.EffectiveMessage.Reply(b, "Please reply to a sticker", nil)
		return err
	}
	var fileId string
	if msg.Sticker != nil {
		fileId = msg.Sticker.FileId
	} else if msg.Animation != nil {
		fileId = msg.Animation.FileId
	} else if msg.Photo != nil {
		fmt.Println(msg.Photo)
		fileId = msg.Photo[0].FileId
	} else if msg.Video != nil {
		fileId = msg.Video.FileId
	}
	err := sb.db.RemoveSticker(fileId)
	if err != nil {
		log.Println(err)
		_, err := ctx.EffectiveMessage.Reply(b, "Failed to remove sticker", nil)
		return err
	}
	_, err = ctx.EffectiveMessage.Reply(b, "Sticker removed", nil)
	return err
}

func (sb *StickerBot) inlineQuery(b *gotgbot.Bot, ctx *ext.Context) error {
	keyword := ctx.InlineQuery.Query
	if keyword == "" {
		ok, err := ctx.InlineQuery.Answer(b, nil, nil)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("answer inline failed")
		}
		return nil
	}
	stickers, err := sb.db.GetStickersFromKeyword(keyword)
	if err != nil {
		return err
	}
	var answers []gotgbot.InlineQueryResult
	for _, sticker := range stickers {
		log.Println(sticker)
		switch sticker.StickerType {
		case "gif":
			answers = append(answers, gotgbot.InlineQueryResultCachedGif{
				Id:        strconv.FormatInt(sticker.ID, 10),
				GifFileId: sticker.FileID,
			})
		case "mpeg4gif":
			answers = append(answers, gotgbot.InlineQueryResultCachedMpeg4Gif{
				Id:          strconv.FormatInt(sticker.ID, 10),
				Mpeg4FileId: sticker.FileID,
			})
		case "photo":
			answers = append(answers, gotgbot.InlineQueryResultCachedPhoto{
				Id:          strconv.FormatInt(sticker.ID, 10),
				PhotoFileId: sticker.FileID,
			})
		case "sticker":
			answers = append(answers, gotgbot.InlineQueryResultCachedSticker{
				Id:            strconv.FormatInt(sticker.ID, 10),
				StickerFileId: sticker.FileID,
			})
		case "video":
			answers = append(answers, gotgbot.InlineQueryResultCachedVideo{
				Id:          strconv.FormatInt(sticker.ID, 10),
				VideoFileId: sticker.FileID,
			})
		}
	}
	ok, err := ctx.InlineQuery.Answer(b, answers, &gotgbot.AnswerInlineQueryOpts{
		IsPersonal: true,
	})
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("answer inline failed")
	}
	return nil
}

func (sb *StickerBot) choosenInlineResult(b *gotgbot.Bot, ctx *ext.Context) error {
	resultId := ctx.ChosenInlineResult.ResultId
	id, err := strconv.ParseInt(resultId, 10, 64)
	if err != nil {
		return err
	}
	return sb.db.UpdateStickerUsage(id)
}
