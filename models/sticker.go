// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Sticker is an object representing the database table.
type Sticker struct {
	ID          int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	FileID      string    `boil:"file_id" json:"file_id" toml:"file_id" yaml:"file_id"`
	StickerType string    `boil:"sticker_type" json:"sticker_type" toml:"sticker_type" yaml:"sticker_type"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *stickerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stickerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StickerColumns = struct {
	ID          string
	FileID      string
	StickerType string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	FileID:      "file_id",
	StickerType: "sticker_type",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var StickerTableColumns = struct {
	ID          string
	FileID      string
	StickerType string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "sticker.id",
	FileID:      "sticker.file_id",
	StickerType: "sticker.sticker_type",
	CreatedAt:   "sticker.created_at",
	UpdatedAt:   "sticker.updated_at",
}

// Generated where

var StickerWhere = struct {
	ID          whereHelperint64
	FileID      whereHelperstring
	StickerType whereHelperstring
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
}{
	ID:          whereHelperint64{field: "\"sticker\".\"id\""},
	FileID:      whereHelperstring{field: "\"sticker\".\"file_id\""},
	StickerType: whereHelperstring{field: "\"sticker\".\"sticker_type\""},
	CreatedAt:   whereHelpertime_Time{field: "\"sticker\".\"created_at\""},
	UpdatedAt:   whereHelpertime_Time{field: "\"sticker\".\"updated_at\""},
}

// StickerRels is where relationship names are stored.
var StickerRels = struct {
	StickerKeywords string
}{
	StickerKeywords: "StickerKeywords",
}

// stickerR is where relationships are stored.
type stickerR struct {
	StickerKeywords StickerKeywordSlice `boil:"StickerKeywords" json:"StickerKeywords" toml:"StickerKeywords" yaml:"StickerKeywords"`
}

// NewStruct creates a new relationship struct
func (*stickerR) NewStruct() *stickerR {
	return &stickerR{}
}

func (r *stickerR) GetStickerKeywords() StickerKeywordSlice {
	if r == nil {
		return nil
	}
	return r.StickerKeywords
}

// stickerL is where Load methods for each relationship are stored.
type stickerL struct{}

var (
	stickerAllColumns            = []string{"id", "file_id", "sticker_type", "created_at", "updated_at"}
	stickerColumnsWithoutDefault = []string{"file_id", "sticker_type", "created_at", "updated_at"}
	stickerColumnsWithDefault    = []string{"id"}
	stickerPrimaryKeyColumns     = []string{"id"}
	stickerGeneratedColumns      = []string{"id"}
)

type (
	// StickerSlice is an alias for a slice of pointers to Sticker.
	// This should almost always be used instead of []Sticker.
	StickerSlice []*Sticker

	stickerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stickerType                 = reflect.TypeOf(&Sticker{})
	stickerMapping              = queries.MakeStructMapping(stickerType)
	stickerPrimaryKeyMapping, _ = queries.BindMapping(stickerType, stickerMapping, stickerPrimaryKeyColumns)
	stickerInsertCacheMut       sync.RWMutex
	stickerInsertCache          = make(map[string]insertCache)
	stickerUpdateCacheMut       sync.RWMutex
	stickerUpdateCache          = make(map[string]updateCache)
	stickerUpsertCacheMut       sync.RWMutex
	stickerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single sticker record from the query.
func (q stickerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Sticker, error) {
	o := &Sticker{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for sticker")
	}

	return o, nil
}

// All returns all Sticker records from the query.
func (q stickerQuery) All(ctx context.Context, exec boil.ContextExecutor) (StickerSlice, error) {
	var o []*Sticker

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Sticker slice")
	}

	return o, nil
}

// Count returns the count of all Sticker records in the query.
func (q stickerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count sticker rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q stickerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if sticker exists")
	}

	return count > 0, nil
}

// StickerKeywords retrieves all the sticker_keyword's StickerKeywords with an executor.
func (o *Sticker) StickerKeywords(mods ...qm.QueryMod) stickerKeywordQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"sticker_keyword\".\"sticker_id\"=?", o.ID),
	)

	return StickerKeywords(queryMods...)
}

// LoadStickerKeywords allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (stickerL) LoadStickerKeywords(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSticker interface{}, mods queries.Applicator) error {
	var slice []*Sticker
	var object *Sticker

	if singular {
		var ok bool
		object, ok = maybeSticker.(*Sticker)
		if !ok {
			object = new(Sticker)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSticker)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSticker))
			}
		}
	} else {
		s, ok := maybeSticker.(*[]*Sticker)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSticker)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSticker))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stickerR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stickerR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`sticker_keyword`),
		qm.WhereIn(`sticker_keyword.sticker_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load sticker_keyword")
	}

	var resultSlice []*StickerKeyword
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice sticker_keyword")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on sticker_keyword")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for sticker_keyword")
	}

	if singular {
		object.R.StickerKeywords = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &stickerKeywordR{}
			}
			foreign.R.Sticker = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.StickerID {
				local.R.StickerKeywords = append(local.R.StickerKeywords, foreign)
				if foreign.R == nil {
					foreign.R = &stickerKeywordR{}
				}
				foreign.R.Sticker = local
				break
			}
		}
	}

	return nil
}

// AddStickerKeywords adds the given related objects to the existing relationships
// of the sticker, optionally inserting them as new records.
// Appends related to o.R.StickerKeywords.
// Sets related.R.Sticker appropriately.
func (o *Sticker) AddStickerKeywords(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*StickerKeyword) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.StickerID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"sticker_keyword\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"sticker_id"}),
				strmangle.WhereClause("\"", "\"", 0, stickerKeywordPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.StickerID, rel.KeywordID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.StickerID = o.ID
		}
	}

	if o.R == nil {
		o.R = &stickerR{
			StickerKeywords: related,
		}
	} else {
		o.R.StickerKeywords = append(o.R.StickerKeywords, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &stickerKeywordR{
				Sticker: o,
			}
		} else {
			rel.R.Sticker = o
		}
	}
	return nil
}

// Stickers retrieves all the records using an executor.
func Stickers(mods ...qm.QueryMod) stickerQuery {
	mods = append(mods, qm.From("\"sticker\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"sticker\".*"})
	}

	return stickerQuery{q}
}

// FindSticker retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSticker(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Sticker, error) {
	stickerObj := &Sticker{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"sticker\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, stickerObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from sticker")
	}

	return stickerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Sticker) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sticker provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(stickerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	stickerInsertCacheMut.RLock()
	cache, cached := stickerInsertCache[key]
	stickerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			stickerAllColumns,
			stickerColumnsWithDefault,
			stickerColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, stickerGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(stickerType, stickerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stickerType, stickerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"sticker\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"sticker\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into sticker")
	}

	if !cached {
		stickerInsertCacheMut.Lock()
		stickerInsertCache[key] = cache
		stickerInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Sticker.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Sticker) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	stickerUpdateCacheMut.RLock()
	cache, cached := stickerUpdateCache[key]
	stickerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			stickerAllColumns,
			stickerPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, stickerGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update sticker, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"sticker\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, stickerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stickerType, stickerMapping, append(wl, stickerPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update sticker row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for sticker")
	}

	if !cached {
		stickerUpdateCacheMut.Lock()
		stickerUpdateCache[key] = cache
		stickerUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q stickerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for sticker")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for sticker")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StickerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stickerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"sticker\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, stickerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in sticker slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all sticker")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Sticker) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sticker provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(stickerColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	stickerUpsertCacheMut.RLock()
	cache, cached := stickerUpsertCache[key]
	stickerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			stickerAllColumns,
			stickerColumnsWithDefault,
			stickerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			stickerAllColumns,
			stickerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert sticker, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stickerPrimaryKeyColumns))
			copy(conflict, stickerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"sticker\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(stickerType, stickerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stickerType, stickerMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert sticker")
	}

	if !cached {
		stickerUpsertCacheMut.Lock()
		stickerUpsertCache[key] = cache
		stickerUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Sticker record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Sticker) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Sticker provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stickerPrimaryKeyMapping)
	sql := "DELETE FROM \"sticker\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from sticker")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for sticker")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q stickerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no stickerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sticker")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sticker")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StickerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stickerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"sticker\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, stickerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sticker slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sticker")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Sticker) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSticker(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StickerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StickerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stickerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"sticker\".* FROM \"sticker\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, stickerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StickerSlice")
	}

	*o = slice

	return nil
}

// StickerExists checks if the Sticker row exists.
func StickerExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"sticker\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if sticker exists")
	}

	return exists, nil
}

// Exists checks if the Sticker row exists.
func (o *Sticker) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return StickerExists(ctx, exec, o.ID)
}
