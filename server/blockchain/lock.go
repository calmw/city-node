package blockchain

import "sync"

// InsertUserLocationLock 用户定位事件日志入库，锁
var InsertUserLocationLock = sync.RWMutex{}

// InsertDailyRewardLock 先锋每日奖励事件日志入库，锁
var InsertDailyRewardLock = sync.RWMutex{}

// InsertRechargeRecordEventLock 充值事件日志入库，锁
var InsertRechargeRecordEventLock = sync.RWMutex{}

// InsertRechargeWeightLock 充值权重入库，锁
var InsertRechargeWeightLock = sync.RWMutex{}

// InsertDelegateLock 新增质押入库，锁
var InsertDelegateLock = sync.RWMutex{}
