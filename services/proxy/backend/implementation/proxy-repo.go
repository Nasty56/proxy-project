package implementation

import (
	"database/sql"

	"github.com/nastyp/documents/services/base"
	"github.com/nastyp/documents/services/proxy/backend/repo"
	"github.com/nastyp/documents/services/proxy/domain"
)

type proxyRepo struct {
	db *sql.DB
}

func NewProxyRepo(db *sql.DB) repo.ProxyRepo {
	return &proxyRepo{
		db: db,
	}
}

// Самая важная функция из методички ЛР №6 — получение доверенности по id
func (r *proxyRepo) ProxyById(id int) (*domain.Proxy, error) {
	sql := `
		SELECT * FROM proxies WHERE id = $1
	`
	proxy := domain.Proxy{}
	if err := r.db.QueryRow(sql, id).Scan(
		&proxy.ProxyHeader.Id,
		&proxy.ProxyHeader.OrganizationId,
		&proxy.ProxyHeader.CustomerId,
		&proxy.ProxyHeader.EmployeeId,
		&proxy.ProxyHeader.DateOfIssue,
		&proxy.ProxyHeader.IsValidUntil,
	); err != nil {
		return nil, err
	}
	proxy.ProxyHeader.DateOfIssue, _ = base.ParseTime(proxy.ProxyHeader.DateOfIssue)
	proxy.ProxyHeader.IsValidUntil, _ = base.ParseTime(proxy.ProxyHeader.IsValidUntil)

	getBodySql := `SELECT * FROM proxy_bodies WHERE proxy_id = $1`
	rows, err := r.db.Query(getBodySql, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		bodyItem := domain.ProxyBodyItem{}
		if err := rows.Scan(
			&bodyItem.Id,
			&bodyItem.ProductId,
			&bodyItem.ProxyId,
			&bodyItem.ProductAmount,
		); err != nil {
			return nil, err
		}
		proxy.ProxyBodyItems = append(proxy.ProxyBodyItems, bodyItem)
	}
	return &proxy, nil
}