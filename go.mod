module github.com/pewt1403/nameservice

go 1.14

require (
	github.com/cosmos/cosmos-sdk v0.38.4
	github.com/gorilla/mux v1.7.4
	github.com/spf13/cobra v0.0.7
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.3
	github.com/tendermint/tm-db v0.5.1
)

//replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
