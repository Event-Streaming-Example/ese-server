module ese.server/redis

go 1.21.5

replace ese.server/models => ../models

require github.com/redis/go-redis/v9 v9.5.1

require (
	ese.server/models v0.0.0-00010101000000-000000000000
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
)
