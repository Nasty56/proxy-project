package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/levigross/grequests"
	"github.com/nastyp/documents/services/base"
	"github.com/nastyp/documents/services/proxy/domain"
	"github.com/nastyp/documents/services/proxy/dto"
	"github.com/nastyp/documents/services/proxy/frontend/utils"
	organizationDomain "github.com/nastyp/documents/services/organization/domain"
	employeeDomain "github.com/nastyp/documents/services/employee/domain"
	customerDomain "github.com/nastyp/documents/services/customer/domain"
	accountDto "github.com/nastyp/documents/services/account/dto"
)

// Список всех доверенностей
func ProxiesHandler(w http.ResponseWriter, r *http.Request) {
	proxiesResp, _ := utils.GetProxies()
	if proxiesResp.Err != "" {
		http.Error(w, proxiesResp.Err, http.StatusInternalServerError)
		return
	}

	type proxyTemplate struct {
		Id           int
		Organization string
		DateOfIssue  string
		IsValidUntil string
		Employee     string
	}

	templateItems := []proxyTemplate{}
	for _, proxy := range proxiesResp.Proxies {
		orgResp, _ := utils.GetOrganizationById(proxy.OrganizationId)
		orgName := ""
		if orgResp.Err == "" && orgResp.Organization != nil {
			orgName = orgResp.Organization.Name
		}

		empResp, _ := utils.GetEmployeeById(proxy.EmployeeId)
		empName := ""
		if empResp.Err == "" && empResp.Employee != nil {
			empName = utils.Fullname(empResp.Employee.LastName, empResp.Employee.FirstName, empResp.Employee.MiddleName)
		}

		templateItems = append(templateItems, proxyTemplate{
			Id:           proxy.Id,
			Organization: orgName,
			DateOfIssue:  proxy.DateOfIssue,
			IsValidUntil: proxy.IsValidUntil,
			Employee:     empName,
		})
	}

	tmpl, _ := template.ParseFiles("../static/html/proxies.html")
	tmpl.Execute(w, templateItems)
}

// Страница создания доверенности (GET)
func CreateProxyGetHandler(w http.ResponseWriter, r *http.Request) {
	type proxyTemplate struct {
		Organizations []organizationDomain.Organization
		Employees     []employeeDomain.Employee
		Customers     []customerDomain.Customer
	}

	orgResp, _ := utils.GetOrganizations()
	empResp, _ := utils.GetEmployees()
	custResp, _ := utils.GetCustomers()

	data := proxyTemplate{}
	if orgResp.Err == "" {
		data.Organizations = orgResp.Organizations
	}
	if empResp.Err == "" {
		data.Employees = empResp.Employees
	}
	if custResp.Err == "" {
		data.Customers = custResp.Customers
	}

	tmpl, _ := template.ParseFiles("../static/html/create-proxy.html")
	tmpl.Execute(w, data)
}

// Создание доверенности (POST)
func CreateProxyPostHandler(w http.ResponseWriter, r *http.Request) {
	organizationId, _ := strconv.Atoi(r.FormValue("organization_id"))
	customerId, _ := strconv.Atoi(r.FormValue("customer_id"))
	employeeId, _ := strconv.Atoi(r.FormValue("employee_id"))
	dateOfIssue := r.FormValue("date_of_issue")
	isValidUntil := r.FormValue("is_valid_until")

	if organizationId == 0 || customerId == 0 || employeeId == 0 {
		http.Error(w, "fields validation error", http.StatusUnprocessableEntity)
		return
	}

	request := dto.CreateProxyHeaderRequest{
		OrganizationId: organizationId,
		CustomerId:     customerId,
		EmployeeId:     employeeId,
		DateOfIssue:    dateOfIssue,
		IsValidUntil:   isValidUntil,
	}

	url := fmt.Sprintf("%s/proxy", utils.GetProxiesAddress())
	resp, _ := grequests.Post(url, &grequests.RequestOptions{JSON: request})

	var data dto.CreateProxyHeaderResponse
	resp.JSON(&data)

	if data.Err != "" {
		http.Error(w, data.Err, http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/documents/proxies", http.StatusSeeOther)
	}
}

// Страница обновления доверенности (GET)
func UpdateProxyGetHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	proxy, _ := utils.GetProxyById(id)
	if proxy.Err != "" {
		http.Error(w, proxy.Err, http.StatusInternalServerError)
		return
	}

	type proxyBodyTemplate struct {
		Id            int
		Name          string
		Price         int
		ProductAmount int
	}

	type updateProxyTemplate struct {
		Proxy          domain.Proxy
		Organizations  []organizationDomain.Organization
		Employees      []employeeDomain.Employee
		Customers      []customerDomain.Customer
		ProxyBodyItems []proxyBodyTemplate
	}

	templateData := updateProxyTemplate{
		Proxy: *proxy.Proxy,
	}

	proxyBodyItems := []proxyBodyTemplate{}
	for _, bodyItem := range proxy.Proxy.ProxyBodyItems {
		product, _ := utils.GetProductById(bodyItem.ProductId)
		if product.Err != "" {
			base.LogError(errors.New(product.Err))
		} else {
			proxyBodyItems = append(proxyBodyItems, proxyBodyTemplate{
				Id:            product.Product.Id,
				Name:          product.Product.Name,
				Price:         product.Product.Price,
				ProductAmount: bodyItem.ProductAmount,
			})
		}
	}
	templateData.ProxyBodyItems = proxyBodyItems

	orgResp, _ := utils.GetOrganizations()
	if orgResp.Err == "" {
		templateData.Organizations = orgResp.Organizations
	}

	empResp, _ := utils.GetEmployees()
	if empResp.Err == "" {
		templateData.Employees = empResp.Employees
	}

	custResp, _ := utils.GetCustomers()
	if custResp.Err == "" {
		templateData.Customers = custResp.Customers
	}

	tmpl, _ := template.ParseFiles("../static/html/update-proxy.html")
	tmpl.Execute(w, templateData)
}

// Обновление доверенности (POST)
func UpdateProxyPostHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	organizationId, _ := strconv.Atoi(r.FormValue("organization_id"))
	customerId, _ := strconv.Atoi(r.FormValue("customer_id"))
	employeeId, _ := strconv.Atoi(r.FormValue("employee_id"))
	dateOfIssue := r.FormValue("date_of_issue")
	isValidUntil := r.FormValue("is_valid_until")

	request := dto.UpdateProxyHeaderRequest{
		ProxyHeader: domain.ProxyHeader{
			Id:             id,
			OrganizationId: organizationId,
			CustomerId:     customerId,
			EmployeeId:     employeeId,
			DateOfIssue:    dateOfIssue,
			IsValidUntil:   isValidUntil,
		},
	}

	url := fmt.Sprintf("%s/proxy", utils.GetProxiesAddress())
	resp, _ := grequests.Put(url, &grequests.RequestOptions{JSON: request})

	var data dto.UpdateProxyHeaderResponse
	resp.JSON(&data)

	if data.Err != "" {
		http.Error(w, data.Err, http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/documents/proxies", http.StatusSeeOther)
	}
}

// Удаление доверенности
func DeleteProxyHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	url := fmt.Sprintf("%s/proxy", utils.GetProxiesAddress())
	resp, _ := grequests.Delete(url, &grequests.RequestOptions{
		JSON: dto.DeleteProxyRequest{Id: id},
	})

	var data dto.DeleteProxyResponse
	resp.JSON(&data)

	if data.Err != "" {
		http.Error(w, data.Err, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/documents/proxies", http.StatusSeeOther)
}

// Детальный просмотр доверенности
func ProxyDetailsHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	proxy, _ := utils.GetProxyById(id)
	if proxy.Err != "" {
		http.Error(w, proxy.Err, http.StatusInternalServerError)
		return
	}

	type proxyBodyTemplate struct {
		ProductId     int
		Name          string
		Price         int
		ProductAmount int
	}

	type proxyDetailTemplate struct {
		ProxyId            int
		Organization       string
		Address            string
		Account            string
		BankName           string
		BankIdentityNumber string
		Employee           string
		PassportSeries     string
		PassportNumber     string
		PassportIssuedBy   string
		PassportDateOfIssue string
		Customer           string
		Chief              string
		FinancialChief     string
		DateOfIssue        string
		IsValidUntil       string
		ProxyBodyItems     []proxyBodyTemplate
	}

	templateData := proxyDetailTemplate{
		ProxyId:      proxy.Proxy.ProxyHeader.Id,
		DateOfIssue:  proxy.Proxy.ProxyHeader.DateOfIssue,
		IsValidUntil: proxy.Proxy.ProxyHeader.IsValidUntil,
	}

	proxyBodyItems := []proxyBodyTemplate{}
	for _, bodyItem := range proxy.Proxy.ProxyBodyItems {
		product, _ := utils.GetProductById(bodyItem.ProductId)
		if product.Err != "" {
			base.LogError(errors.New(product.Err))
		} else {
			proxyBodyItems = append(proxyBodyItems, proxyBodyTemplate{
				ProductId:     product.Product.Id,
				Name:          product.Product.Name,
				Price:         product.Product.Price,
				ProductAmount: bodyItem.ProductAmount,
			})
		}
	}
	templateData.ProxyBodyItems = proxyBodyItems

	orgResp, _ := utils.GetOrganizationById(proxy.Proxy.ProxyHeader.OrganizationId)
	if orgResp.Err != "" {
		base.LogError(errors.New(orgResp.Err))
		templateData.Organization = fmt.Sprintf("Ошибка: %s", orgResp.Err)
	} else {
		templateData.Organization = orgResp.Organization.Name
		templateData.Address = orgResp.Organization.Address
		templateData.Chief = utils.NameInitials(orgResp.Organization.Chief)
		templateData.FinancialChief = utils.NameInitials(orgResp.Organization.FinancialChief)
	}

	if orgResp.Organization != nil {
		accountUrl := fmt.Sprintf("%s/accounts/%d", utils.GetAccountsAddress(), orgResp.Organization.AccountId)
		accountResp, _ := grequests.Get(accountUrl, &grequests.RequestOptions{
			JSON: accountDto.AccountByIdRequest{Id: orgResp.Organization.AccountId},
		})
		var account accountDto.AccountByIdResponse
		accountResp.JSON(&account)
		if account.Err != "" {
			templateData.Account = fmt.Sprintf("Ошибка: %s", account.Err)
		} else {
			templateData.Account = account.Account.Account
			templateData.BankName = account.Account.BankName
			templateData.BankIdentityNumber = account.Account.BankIdentityNumber
		}
	}

	empResp, _ := utils.GetEmployeeById(proxy.Proxy.ProxyHeader.EmployeeId)
	if empResp.Err != "" {
		base.LogError(errors.New(empResp.Err))
		templateData.Employee = fmt.Sprintf("Ошибка: %s", empResp.Err)
	} else {
		templateData.Employee = utils.Fullname(
			empResp.Employee.LastName,
			empResp.Employee.FirstName,
			empResp.Employee.MiddleName,
		)
		templateData.PassportSeries = empResp.Employee.PassportSeries
		templateData.PassportNumber = empResp.Employee.PassportNumber
		templateData.PassportDateOfIssue = empResp.Employee.PassportDateOfIssue
		templateData.PassportIssuedBy = empResp.Employee.PassportIssuedBy
	}

	custResp, _ := utils.GetCustomerById(proxy.Proxy.ProxyHeader.CustomerId)
	if custResp.Err != "" {
		base.LogError(errors.New(custResp.Err))
		templateData.Customer = fmt.Sprintf("Ошибка: %s", custResp.Err)
	} else {
		templateData.Customer = custResp.Customer.Name
	}

	tmpl, _ := template.ParseFiles("../static/html/proxy-details.html")
	tmpl.Execute(w, templateData)
}

// Страница добавления товара в тело доверенности (GET)
func CreateProxyBodyGetHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	proxy, _ := utils.GetProxyById(id)
	if proxy.Err != "" {
		http.Error(w, proxy.Err, http.StatusInternalServerError)
		return
	}

	type proxyDetail struct {
		Id           int
		Organization string
		Employee     string
		Customer     string
		DateOfIssue  string
		IsValidUntil string
	}

	type createProxyBodyTemplate struct {
		Proxy    proxyDetail
		Products []domain.Product
	}

	orgResp, _ := utils.GetOrganizationById(proxy.Proxy.ProxyHeader.OrganizationId)
	empResp, _ := utils.GetEmployeeById(proxy.Proxy.ProxyHeader.EmployeeId)
	custResp, _ := utils.GetCustomerById(proxy.Proxy.ProxyHeader.CustomerId)
	productsResp, _ := utils.GetProducts()

	proxyDetail := proxyDetail{
		Id:           proxy.Proxy.ProxyHeader.Id,
		DateOfIssue:  proxy.Proxy.ProxyHeader.DateOfIssue,
		IsValidUntil: proxy.Proxy.ProxyHeader.IsValidUntil,
	}

	if orgResp.Err == "" && orgResp.Organization != nil {
		proxyDetail.Organization = orgResp.Organization.Name
	}
	if empResp.Err == "" && empResp.Employee != nil {
		proxyDetail.Employee = utils.Fullname(empResp.Employee.LastName, empResp.Employee.FirstName, empResp.Employee.MiddleName)
	}
	if custResp.Err == "" && custResp.Customer != nil {
		proxyDetail.Customer = custResp.Customer.Name
	}

	var products []domain.Product
	if productsResp.Err == "" {
		products = productsResp.Products
	}

	data := createProxyBodyTemplate{
		Proxy:    proxyDetail,
		Products: products,
	}

	tmpl, _ := template.ParseFiles("../static/html/create-proxy-body.html")
	tmpl.Execute(w, data)
}

// Добавление товара в тело доверенности (POST)
func CreateProxyBodyPostHandler(w http.ResponseWriter, r *http.Request) {
	proxyId, _ := strconv.Atoi(r.FormValue("proxy_id"))
	productId, _ := strconv.Atoi(r.FormValue("product_id"))
	productAmount, _ := strconv.Atoi(r.FormValue("product_amount"))

	if proxyId == 0 || productId == 0 || productAmount == 0 {
		http.Error(w, "fields validation error", http.StatusUnprocessableEntity)
		return
	}

	request := dto.CreateProxyBodyItemRequest{
		ProductId:     productId,
		ProxyId:       proxyId,
		ProductAmount: productAmount,
	}

	url := fmt.Sprintf("%s/proxy/body", utils.GetProxiesAddress())
	resp, _ := grequests.Post(url, &grequests.RequestOptions{JSON: request})

	var data dto.CreateProxyBodyItemResponse
	resp.JSON(&data)

	if data.Err != "" {
		http.Error(w, data.Err, http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, fmt.Sprintf("/documents/proxies/details/%d", proxyId), http.StatusSeeOther)
	}
}

// Удаление товара из тела доверенности
func DeleteProxyBodyHandler(w http.ResponseWriter, r *http.Request) {
	itemId, _ := strconv.Atoi(r.FormValue("item_id"))
	proxyId, _ := strconv.Atoi(r.FormValue("proxy_id"))

	url := fmt.Sprintf("%s/proxy/body", utils.GetProxiesAddress())
	resp, _ := grequests.Delete(url, &grequests.RequestOptions{
		JSON: dto.DeleteProxyBodyItemRequest{ItemId: itemId},
	})

	var data dto.DeleteProxyBodyItemResponse
	resp.JSON(&data)

	if data.Err != "" {
		http.Error(w, data.Err, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/documents/proxies/details/%d", proxyId), http.StatusSeeOther)
}