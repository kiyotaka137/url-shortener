package repository

import (
	"context"
	"fmt"
	"url-shortener/internal/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresRepo struct {
	pool *pgxpool.Pool
}

func NewRepository(ctx context.Context, config *config.Config) (*PostgresRepo, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.DBName,
		config.Database.SSLMode,
	)

	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	return &PostgresRepo{
		pool: pool,
	}, nil
}

func (r *PostgresRepo) Close() {
	if r.pool != nil {
		r.pool.Close()
	}
}

func (r *PostgresRepo) Create(ctx context.Context, urlToSave, alias string) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO urls(url,alias) VALUES ($1,$2)", urlToSave, alias)
	return err
}

func (r *PostgresRepo) Get(ctx context.Context, alias string) (string, error) {
	var url string
	err := r.pool.QueryRow(ctx,
		"SELECT url FROM urls WHERE alias=$1", alias).Scan(&url)
	return url, err
}

func (r *PostgresRepo) Delete(ctx context.Context, alias string) error {
	_, err := r.pool.Exec(ctx, "DELETE FROM urls WHERE alias=$1", alias)
	return err
}
