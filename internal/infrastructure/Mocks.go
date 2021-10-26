package infrastructure

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type MockApplicationDatabase struct {
	callParams []interface{}
}

func (mdb *MockApplicationDatabase) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	mdb.callParams = []interface{}{sql}
	mdb.callParams = append(mdb.callParams, arguments...)

	return nil, nil
}

func (mdb *MockApplicationDatabase) Begin(ctx context.Context) (pgx.Tx, error) {
	return MockTx{}, nil
}

func (mdb *MockApplicationDatabase) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	mdb.callParams = []interface{}{sql}
	mdb.callParams = append(mdb.callParams, args...)

	return nil, nil
}

func (mdb *MockApplicationDatabase) CalledWith() []interface{} {
	return mdb.callParams
}

type MockTx struct{}

func (mtx MockTx) Begin(ctx context.Context) (pgx.Tx, error) {
	return nil, nil
}

func (mtx MockTx) BeginFunc(ctx context.Context, f func(pgx.Tx) error) (err error) {
	return nil
}

func (mtx MockTx) Commit(ctx context.Context) error {
	return nil
}

func (mtx MockTx) Rollback(ctx context.Context) error {
	return nil
}

func (mtx MockTx) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return 0, nil
}

func (mtx MockTx) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return nil
}

func (mtx MockTx) LargeObjects() pgx.LargeObjects {
	return pgx.LargeObjects{}
}

func (mtx MockTx) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
	return nil, nil
}

func (mtx MockTx) Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error) {
	return nil, nil
}

func (mtx MockTx) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return nil, nil
}

func (mtx MockTx) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return MockRow{}
}

func (mtx MockTx) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	return nil, nil
}

func (mtx MockTx) Conn() *pgx.Conn {
	return nil
}

type MockRow struct{}

func (mr MockRow) Scan(dest ...interface{}) error {
	return nil
}
