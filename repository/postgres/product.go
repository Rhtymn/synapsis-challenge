package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Rhtymn/synapsis-challenge/apperror"
	"github.com/Rhtymn/synapsis-challenge/constants"
	"github.com/Rhtymn/synapsis-challenge/domain"
	"github.com/Rhtymn/synapsis-challenge/util"
	"github.com/jackc/pgx/v5"
)

type productRepositoryPostgres struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepositoryPostgres {
	return &productRepositoryPostgres{
		db: db,
	}
}

func (r *productRepositoryPostgres) GetByID(ctx context.Context, id int64) (domain.Product, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	p := domain.Product{}
	args := pgx.NamedArgs{
		"id": id,
	}
	query := `
		SELECT ` + constants.ProductJoinedShopColumns + `
			FROM products p INNER JOIN shops s ON p.id_shop = s.id
		WHERE p.id = @id AND p.deleted_at IS NULL
	`

	err := queryRunner.
		QueryRowContext(ctx, query, args).
		Scan(&p.ID,
			&p.Name,
			&p.Slug,
			&p.PhotoURL,
			&p.Price,
			&p.Description,
			&p.Stock,
			&p.Shop.ID,
			&p.Shop.ShopName,
			&p.Shop.Slug,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, apperror.NewNotFound(err, "product not found")
		}
		return p, apperror.Wrap(err)
	}

	return p, nil
}

func (r *productRepositoryPostgres) GetAll(ctx context.Context, query domain.ProductQuery) ([]domain.Product, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	sb := strings.Builder{}
	sb.WriteString(`
		SELECT ` + constants.ProductJoinedShopColumns + `
		FROM products p INNER JOIN categories c ON p.id_category = c.id
			INNER JOIN shops s ON p.id_shop = s.id
		WHERE p.deleted_at IS NULL AND s.is_active = TRUE
	`)
	args := pgx.NamedArgs{}

	if query.CategorySlug != "" {
		sb.WriteString(` AND c.slug = @categorySlug `)
		args["categorySlug"] = query.CategorySlug
	}

	if query.Search != "" {
		sb.WriteString(` AND p.name ILIKE '%' || @search || '%' `)
		args["search"] = query.Search
	}

	fmt.Fprintf(&sb, " ORDER BY p.%s %s ", query.SortBy, query.SortType)

	offset := (query.Page - 1) * query.Limit
	if query.Limit != 0 {
		fmt.Fprintf(&sb, " OFFSET %d LIMIT %d ", offset, query.Limit)
	}

	rows, err := queryRunner.QueryContext(ctx, sb.String(), args)
	if err != nil {
		return nil, apperror.Wrap(err)
	}
	defer rows.Close()

	products := []domain.Product{}
	for rows.Next() {
		p := domain.Product{}
		err := rows.Scan(&p.ID,
			&p.Name,
			&p.Slug,
			&p.PhotoURL,
			&p.Price,
			&p.Description,
			&p.Stock,
			&p.Shop.ID,
			&p.Shop.ShopName,
			&p.Shop.Slug,
		)
		if err != nil {
			return nil, apperror.Wrap(err)
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *productRepositoryPostgres) GetPageInfo(ctx context.Context, query domain.ProductQuery) (domain.PageInfo, error) {
	queryRunner := util.GetQueryRunner(ctx, r.db)
	sb := strings.Builder{}
	sb.WriteString(`
		SELECT COUNT(*) as total_data
		FROM products p INNER JOIN categories c ON p.id_category = c.id
			INNER JOIN shops s ON p.id_shop = s.id
		WHERE p.deleted_at IS NULL
	`)
	args := pgx.NamedArgs{}

	if query.CategorySlug != "" {
		sb.WriteString(` AND c.slug = @categorySlug `)
		args["categorySlug"] = query.CategorySlug
	}

	if query.Search != "" {
		sb.WriteString(` AND p.name ILIKE '%' || @search || '%' `)
		args["search"] = query.Search
	}

	var totalData int64
	err := queryRunner.QueryRowContext(ctx, sb.String(), args).Scan(&totalData)
	if err != nil {
		return domain.PageInfo{}, apperror.Wrap(err)
	}

	return domain.PageInfo{
		CurrentPage: int(query.Page),
		ItemCount:   totalData,
	}, nil
}
