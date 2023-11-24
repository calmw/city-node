package models

type BlockStore struct {
	Id         int   `gorm:"column:id;primaryKey"` // 1 用户定位事件处理，2 城市先锋奖励事件，3增加充值事件，4获取新增质押事件，5获取奖励领取事件
	StartBlock int64 `json:"start_block" gorm:"column:start_block"`
}
