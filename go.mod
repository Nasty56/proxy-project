module github.com/nastyp/documents

go 1.22

replace github.com/nastyp/documents => ./

require (
	github.com/go-kit/kit v0.13.0
	github.com/go-kit/log v0.2.1
	github.com/gorilla/mux v1.8.1
	github.com/levigross/grequests v0.0.0-20250606031859-3f3c12e4e704
	github.com/lib/pq v1.12.3
	github.com/rs/cors v1.11.1
)

require github.com/go-logfmt/logfmt v0.5.1 // indirect
