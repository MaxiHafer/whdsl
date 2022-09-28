package mariadb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/avast/retry-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oiime/logrusbun"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func NewInitializedBackendFromEnv(ctx context.Context) (*Backend, error) {
	c, cErr := NewConfigFromEnv()
	if cErr != nil {
		return nil, cErr
	}

	b, bErr := NewBackendFromConfig(c)
	if bErr != nil {
		return nil, bErr
	}

	return b, b.Init(ctx)
}

func NewBackendFromConfig(c *Config) (*Backend, error) {
	dsn := c.DSN()
	logrus.WithField("DSN", dsn).Debug("built DSN for MariaDB")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failure when trying to open mysql-db")
	}

	err = retry.Do(
		db.Ping,
		retry.Attempts(c.Retries),
		retry.Delay(time.Millisecond*500),
		retry.OnRetry(
			func(n uint, err error) {
				logrus.WithField("retries", fmt.Sprintf("%d/%d", n, c.Retries)).
					WithError(err).
					Warn("failed to ping mysql-db, retrying")
			},
		),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failure when trying to ping mysql-db")
	}

	bunDB := bun.NewDB(db, mysqldialect.New())
	bunDB.AddQueryHook(
		logrusbun.NewQueryHook(
			logrusbun.QueryHookOptions{
				LogSlow:         time.Second,
				Logger:          logrus.StandardLogger(),
				QueryLevel:      logrus.DebugLevel,
				ErrorLevel:      logrus.ErrorLevel,
				SlowLevel:       logrus.WarnLevel,
				MessageTemplate: "{{.Operation}}[{{.Duration}}]: {{.Query}}",
				ErrorTemplate:   "{{.Operation}}[{{.Duration}}]: {{.Query}}: {{.Error}}",
			},
		),
	)

	return &Backend{
		cfg:   c,
		bunDB: bunDB,
	}, nil
}

type Backend struct {
	cfg   *Config
	bunDB *bun.DB
}

func (b *Backend) List(ctx context.Context, slice *[]Model) error {
	return b.bunDB.NewSelect().Model(slice).Scan(ctx)
}

func (b *Backend) BindByID(ctx context.Context, id string, model Model) error {
	return b.bunDB.NewSelect().Model(model).WherePK().Scan(ctx, model)
}

func (b *Backend) InsertOrUpdate(ctx context.Context, model Model) error {
	exists, existsErr := b.bunDB.NewSelect().Model(model).WherePK().Exists(ctx)
	if existsErr != nil {
		return existsErr
	}

	if exists {
		_, err := b.bunDB.NewUpdate().Model(model).WherePK().Exec(ctx)
		return err
	}

	_, err := b.bunDB.NewInsert().Model(model).Exec(ctx)
	return err
}

func (b *Backend) Delete(ctx context.Context, model Model) error {
	_, err := b.bunDB.NewDelete().Model(model).WherePK().Exec(ctx)
	return err
}

func (b *Backend) Init(ctx context.Context) error {
	if err := b.initModels(ctx); err != nil {
		return err
	}

	return nil
}

func (b *Backend) Close() error {
	return b.bunDB.Close()
}
