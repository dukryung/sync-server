module github.com/dukryung/sync-server

go 1.17

require (
	github.com/Workiva/go-datastructures v1.0.53 // indirect
	github.com/cosmos/cosmos-sdk v0.44.4
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/lib/pq v1.10.4 // indirect
	github.com/rs/cors v1.8.0 // indirect
	github.com/spf13/viper v1.9.0 // indirect
	github.com/tendermint/tm-db v0.6.6 // indirect
	golang.org/x/net v0.0.0-20211208012354-db4efeb81f4b // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	google.golang.org/grpc v1.42.0
)

replace github.com/hessegg/klaatoo => ../klaatoo

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
