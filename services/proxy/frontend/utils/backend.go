package utils

import (
	"fmt"

	"github.com/nastyp/documents/services/base"
	"github.com/nastyp/documents/services/proxy/dto"
	"github.com/nastyp/documents/services/product/dto"
	"github.com/nastyp/documents/services/organization/dto"
	"github.com/nastyp/documents/services/employee/dto"
	"github.com/nastyp/documents/services/customer/dto"
	"github.com/levigross/grequests"
)

const (
	PROXIES_ENV         = "PROXY_BACKEND_ADDRESS"
	EMPLOYEES_ENV       = "EMPLOYEE_BACKEND_ADDRESS"
	PRODUCTS_ENV        = "PRODUCT_BACKEND_ADDRESS"
	ORGANIZATIONS_ENV   = "ORGANIZATION_BACKEND_ADDRESS"
	CUSTOMERS_ENV       = "CUSTOMER_BACKEND_ADDRESS"
	ACCOUNTS_ENV        = "ACCOUNT_BACKEND_ADDRESS"
)

func GetProxiesAddress() string {
	return base.GetEnv(PROXIES_ENV, "http://localhost:7076")
}

func GetEmployeesAddress() string {
	return base.GetEnv(EMPLOYEES_ENV, "http://localhost:7074")
}

func GetProductsAddress() string {
	return base.GetEnv(PRODUCTS_ENV, "http://localhost:7071")
}

func GetOrganizationsAddress() string {
	return base.GetEnv(ORGANIZATIONS_ENV, "http://localhost:7072")
}

func GetCustomersAddress() string {
	return base.GetEnv(CUSTOMERS_ENV, "http://localhost:7075")
}

func GetAccountsAddress() string {
	return base.GetEnv(ACCOUNTS_ENV, "http://localhost:7073")
}

func GetProxyById(id int) (*dto.ProxyByIdResponse, error) {
	url := fmt.Sprintf("%s/proxy/%d", GetProxiesAddress(), id)
	resp, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data dto.ProxyByIdResponse
	if err := resp.JSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetProxies() (*dto.GetProxiesResponse, error) {
	url := fmt.Sprintf("%s/proxy", GetProxiesAddress())
	resp, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data dto.GetProxiesResponse
	if err := resp.JSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetProducts() (*productDto.GetProductsResponse, error) {
	url := fmt.Sprintf("%s/products", GetProductsAddress())
	resp, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data productDto.GetProductsResponse
	if err := resp.JSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetProductById(id int) (*productDto.ProductByIdResponse, error) {
	url := fmt.Sprintf("%s/products/%d", GetProductsAddress(), id)
	resp, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data productDto.ProductByIdResponse
	if err := resp.JSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetOrganizations() (*organizationDto.GetOrganizationsResponse, error) {
	url := fmt.Sprintf("%s/organizations", GetOrganizationsAddress())
	resp, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data organizationDto.GetOrganizationsResponse
	if err := resp.JSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetOrganizationById(id int) (*organizationDto.OrganizationByIdResponse, error) {
	url := fmt.Sprintf("%s/organizations/%d", GetOrganizationsAddress(), id)
	resp, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data organizationDto.OrganizationByIdResponse
	if err := resp.JSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetCustomers() (*customerDto.GetCustomersResponse, error) {
	url := fmt.Sprintf("%s/customers", GetCustomersAddress())
	resp, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data customerDto.GetCustomersResponse
	if err := resp.JSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetCustomerById(id int) (*customerDto.CustomerByIdResponse, error) {
	url := fmt.Sprintf("%s/customers/%d", GetCustomersAddress(), id)
	resp, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data customerDto.CustomerByIdResponse
	if err := resp.JSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetEmployees() (*employeeDto.GetEmployeesResponse, error) {
	url := fmt.Sprintf("%s/employees", GetEmployeesAddress())
	resp, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data employeeDto.GetEmployeesResponse
	if err := resp.JSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetEmployeeById(id int) (*employeeDto.EmployeeByIdResponse, error) {
	url := fmt.Sprintf("%s/employees/%d", GetEmployeesAddress(), id)
	resp, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data employeeDto.EmployeeByIdResponse
	if err := resp.JSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}