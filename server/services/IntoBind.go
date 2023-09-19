package services

import (
	"city-node-server/blockchain"
	"city-node-server/db"
	"city-node-server/models"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

func IntoBind() {
	var user []models.User
	var total int64

	db.Mysql.Model(&models.User{}).Find(&user)
	fmt.Println(total, 999)
	for i, u := range user {
		address := common.HexToAddress(u.UserAddress)
		err, res := blockchain.GetUidsWithAddress([]common.Address{address})
		fmt.Println(i, res, err)
		if res[0] != "" {
			db.Mysql.Model(&models.User{}).Where("user_address=?", u.UserAddress).Update("phone_number", res[0])
		} else {
			db.Mysql.Model(&models.User{}).Where("user_address=?", u.UserAddress).Update("phone_number", "")
		}
	}

}
