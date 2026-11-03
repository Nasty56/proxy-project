package repo

import "github.com/nastyp/documents/services/proxy/domain"

type ProxyRepo interface {
	CreateProxyHeader(domain.ProxyHeader) (*domain.ProxyHeader, error)
	CreateProxyBodyItem(domain.ProxyBodyItem) (*domain.Proxy, error)
	GetProxies() ([]domain.ProxyHeader, error)
	ProxyById(id int) (*domain.Proxy, error)
	UpdateProxyHeader(domain.ProxyHeader) (*domain.Proxy, error)
	DeleteProxy(id int) error
	DeleteProxyBodyItem(itemId int) error
}