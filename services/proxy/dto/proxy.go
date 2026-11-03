package dto

import (
	"github.com/nastyp/documents/services/proxy/domain"
)

type (
	CreateProxyHeaderRequest struct {
		OrganizationId int    `json:"organization_id"`
		CustomerId     int    `json:"customer_id"`
		EmployeeId     int    `json:"employee_id"`
		DateOfIssue    string `json:"date_of_issue"`
		IsValidUntil   string `json:"is_valid_until"`
	}

	CreateProxyHeaderResponse struct {
		ProxyHeader *domain.ProxyHeader `json:"proxy_header,omitempty"`
		Err         string              `json:"error,omitempty"`
	}

	GetProxiesRequest struct {
	}

	GetProxiesResponse struct {
		Proxies []domain.ProxyHeader `json:"proxies,omitempty"`
		Err     string               `json:"error,omitempty"`
	}

	ProxyByIdRequest struct {
		Id int `json:"id"`
	}

	ProxyByIdResponse struct {
		Proxy *domain.Proxy `json:"proxy,omitempty"`
		Err   string        `json:"error,omitempty"`
	}

	UpdateProxyHeaderRequest struct {
		ProxyHeader domain.ProxyHeader `json:"proxy_header"`
	}

	UpdateProxyHeaderResponse struct {
		Proxy *domain.Proxy `json:"proxy,omitempty"`
		Err   string        `json:"error,omitempty"`
	}

	DeleteProxyRequest struct {
		Id int `json:"id"`
	}

	DeleteProxyResponse struct {
		Err string `json:"error,omitempty"`
	}

	CreateProxyBodyItemRequest struct {
		ProductId     int `json:"product_id"`
		ProxyId       int `json:"proxy_id"`
		ProductAmount int `json:"product_amount"`
	}

	CreateProxyBodyItemResponse struct {
		Proxy *domain.Proxy `json:"proxy,omitempty"`
		Err   string        `json:"error,omitempty"`
	}

	DeleteProxyBodyItemRequest struct {
		ItemId int `json:"item_id"`
	}

	DeleteProxyBodyItemResponse struct {
		Err string `json:"error,omitempty"`
	}
)