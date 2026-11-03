package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/nastyp/documents/services/base"
	"github.com/nastyp/documents/services/proxy/backend/implementation"
	"github.com/nastyp/documents/services/proxy/backend/transport"
	httptransport "github.com/nastyp/documents/services/proxy/backend/transport/http"
	kithttp "github.com/go-kit/kit/transport/http"
)

func main() {
	// Подключение к БД
	db, err := base.ConnectToDb()
	if err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
	defer db.Close()

	// Репозиторий доверенностей
	proxyRepo := implementation.NewProxyRepo(db)

	// Создание сервиса
	svc := implementation.NewProxyService(proxyRepo)

	// Создание эндпоинтов
	endpoints := transport.MakeEndpoints(svc)

	// Создание маршрутизатора
	h := httptransport.NewService(endpoints, []kithttp.ServerOption{})

	// Запуск сервиса на порту 7076
	fmt.Println("Listening on 7076...")
	if err := http.ListenAndServe(":7076", h); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
}