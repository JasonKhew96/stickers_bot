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

// Keyword is an object representing the database table.
type Keyword struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Keyword   string    `boil:"keyword" json:"keyword" toml:"keyword" yaml:"keyword"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *keywordR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L keywordL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var KeywordColumns = struct {
	ID        string
	Keyword   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Keyword:   "keyword",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var KeywordTableColumns = struct {
	ID        string
	Keyword   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "keyword.id",
	Keyword:   "keyword.keyword",
	CreatedAt: "keyword.created_at",
	UpdatedAt: "keyword.updated_at",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod  { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var KeywordWhere = struct {
	ID        whereHelperint64
	Keyword   whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "\"keyword\".\"id\""},
	Keyword:   whereHelperstring{field: "\"keyword\".\"keyword\""},
	CreatedAt: whereHelpertime_Time{field: "\"keyword\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"keyword\".\"updated_at\""},
}

// KeywordRels is where relationship names are stored.
var KeywordRels = struct {
	StickerKeywords string
}{
	StickerKeywords: "StickerKeywords",
}

// keywordR is where relationships are stored.
type keywordR struct {
	StickerKeywords StickerKeywordSlice `boil:"StickerKeywords" json:"StickerKeywords" toml:"StickerKeywords" yaml:"StickerKeywords"`
}

// NewStruct creates a new relationship struct
func (*keywordR) NewStruct() *keywordR {
	return &keywordR{}
}

func (r *keywordR) GetStickerKeywords() StickerKeywordSlice {
	if r == nil {
		return nil
	}
	return r.StickerKeywords
}

// keywordL is where Load methods for each relationship are stored.
type keywordL struct{}

var (
	keywordAllColumns            = []string{"id", "keyword", "created_at", "updated_at"}
	keywordColumnsWithoutDefault = []string{"keyword", "created_at", "updated_at"}
	keywordColumnsWithDefault    = []string{"id"}
	keywordPrimaryKeyColumns     = []string{"id"}
	keywordGeneratedColumns      = []string{"id"}
)

type (
	// KeywordSlice is an alias for a slice of pointers to Keyword.
	// This should almost always be used instead of []Keyword.
	KeywordSlice []*Keyword

	keywordQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	keywordType                 = reflect.TypeOf(&Keyword{})
	keywordMapping              = queries.MakeStructMapping(keywordType)
	keywordPrimaryKeyMapping, _ = queries.BindMapping(keywordType, keywordMapping, keywordPrimaryKeyColumns)
	keywordInsertCacheMut       sync.RWMutex
	keywordInsertCache          = make(map[string]insertCache)
	keywordUpdateCacheMut       sync.RWMutex
	keywordUpdateCache          = make(map[string]updateCache)
	keywordUpsertCacheMut       sync.RWMutex
	keywordUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single keyword record from the query.
func (q keywordQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Keyword, error) {
	o := &Keyword{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for keyword")
	}

	return o, nil
}

// All returns all Keyword records from the query.
func (q keywordQuery) All(ctx context.Context, exec boil.ContextExecutor) (KeywordSlice, error) {
	var o []*Keyword

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Keyword slice")
	}

	return o, nil
}

// Count returns the count of all Keyword records in the query.
func (q keywordQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count keyword rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q keywordQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if keyword exists")
	}

	return count > 0, nil
}

// StickerKeywords retrieves all the sticker_keyword's StickerKeywords with an executor.
func (o *Keyword) StickerKeywords(mods ...qm.QueryMod) stickerKeywordQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"sticker_keyword\".\"keyword_id\"=?", o.ID),
	)

	return StickerKeywords(queryMods...)
}

// LoadStickerKeywords allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (keywordL) LoadStickerKeywords(ctx context.Context, e boil.ContextExecutor, singular bool, maybeKeyword interface{}, mods queries.Applicator) error {
	var slice []*Keyword
	var object *Keyword

	if singular {
		var ok bool
		object, ok = maybeKeyword.(*Keyword)
		if !ok {
			object = new(Keyword)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeKeyword)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeKeyword))
			}
		}
	} else {
		s, ok := maybeKeyword.(*[]*Keyword)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeKeyword)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeKeyword))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &keywordR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &keywordR{}
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
		qm.WhereIn(`sticker_keyword.keyword_id in ?`, args...),
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
			foreign.R.Keyword = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.KeywordID {
				local.R.StickerKeywords = append(local.R.StickerKeywords, foreign)
				if foreign.R == nil {
					foreign.R = &stickerKeywordR{}
				}
				foreign.R.Keyword = local
				break
			}
		}
	}

	return nil
}

// AddStickerKeywords adds the given related objects to the existing relationships
// of the keyword, optionally inserting them as new records.
// Appends related to o.R.StickerKeywords.
// Sets related.R.Keyword appropriately.
func (o *Keyword) AddStickerKeywords(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*StickerKeyword) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.KeywordID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"sticker_keyword\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"keyword_id"}),
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

			rel.KeywordID = o.ID
		}
	}

	if o.R == nil {
		o.R = &keywordR{
			StickerKeywords: related,
		}
	} else {
		o.R.StickerKeywords = append(o.R.StickerKeywords, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &stickerKeywordR{
				Keyword: o,
			}
		} else {
			rel.R.Keyword = o
		}
	}
	return nil
}

// Keywords retrieves all the records using an executor.
func Keywords(mods ...qm.QueryMod) keywordQuery {
	mods = append(mods, qm.From("\"keyword\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"keyword\".*"})
	}

	return keywordQuery{q}
}

// FindKeyword retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindKeyword(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Keyword, error) {
	keywordObj := &Keyword{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"keyword\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, keywordObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from keyword")
	}

	return keywordObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Keyword) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no keyword provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(keywordColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	keywordInsertCacheMut.RLock()
	cache, cached := keywordInsertCache[key]
	keywordInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			keywordAllColumns,
			keywordColumnsWithDefault,
			keywordColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, keywordGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(keywordType, keywordMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(keywordType, keywordMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"keyword\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"keyword\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into keyword")
	}

	if !cached {
		keywordInsertCacheMut.Lock()
		keywordInsertCache[key] = cache
		keywordInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Keyword.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Keyword) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	keywordUpdateCacheMut.RLock()
	cache, cached := keywordUpdateCache[key]
	keywordUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			keywordAllColumns,
			keywordPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, keywordGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update keyword, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"keyword\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, keywordPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(keywordType, keywordMapping, append(wl, keywordPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update keyword row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for keyword")
	}

	if !cached {
		keywordUpdateCacheMut.Lock()
		keywordUpdateCache[key] = cache
		keywordUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q keywordQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for keyword")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for keyword")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o KeywordSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), keywordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"keyword\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, keywordPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in keyword slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all keyword")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Keyword) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no keyword provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(keywordColumnsWithDefault, o)

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

	keywordUpsertCacheMut.RLock()
	cache, cached := keywordUpsertCache[key]
	keywordUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			keywordAllColumns,
			keywordColumnsWithDefault,
			keywordColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			keywordAllColumns,
			keywordPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert keyword, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(keywordPrimaryKeyColumns))
			copy(conflict, keywordPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"keyword\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(keywordType, keywordMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(keywordType, keywordMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert keyword")
	}

	if !cached {
		keywordUpsertCacheMut.Lock()
		keywordUpsertCache[key] = cache
		keywordUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Keyword record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Keyword) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Keyword provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), keywordPrimaryKeyMapping)
	sql := "DELETE FROM \"keyword\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from keyword")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for keyword")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q keywordQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no keywordQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from keyword")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for keyword")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o KeywordSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), keywordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"keyword\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, keywordPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from keyword slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for keyword")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Keyword) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindKeyword(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *KeywordSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := KeywordSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), keywordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"keyword\".* FROM \"keyword\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, keywordPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in KeywordSlice")
	}

	*o = slice

	return nil
}

// KeywordExists checks if the Keyword row exists.
func KeywordExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"keyword\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if keyword exists")
	}

	return exists, nil
}

// Exists checks if the Keyword row exists.
func (o *Keyword) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return KeywordExists(ctx, exec, o.ID)
}