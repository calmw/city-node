package models

type Config struct {
	Id         int32  `gorm:"column:id;primaryKey"`
	RpcMainnet string `json:"rpc_mainnet" gorm:"column:rpc_mainnet"`
	RpcTestnet string `json:"rpc_testnet" gorm:"column:rpc_testnet"`
}
