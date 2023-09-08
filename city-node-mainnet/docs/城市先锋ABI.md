## 城市先锋

#### 查询城市保证金、城市等级，城市先锋地址

- 查询所有先锋城市
    - 查询先锋城市总数量
        - getPioneerCityNumber()
    - 根据总数量遍历查询成城市ID
        - pioneerCityIds，参数为整数索引
- 根据城市ID查询
    - 查看城市等级
        - cityLevel
    - 查看城市先锋
        - cityPioneer
    - 查看城市保证金
        - earnestMoney
- 根据地址查询先锋对应的城市
    - pioneerCity
- 查询城市保证金数量
    - 根据城市ID查询城市所需保证金数量
        - earnestMoney

#### 查询质押权重

- 根据城市ID查询城市新增质押（累计新增质押）量
    - cityDelegate

#### 查询先锋信息

- 根据先锋地址查询先锋信息,先锋字段备注

``` solidity
    struct Pioneer {
        address pioneerAddress;
        uint256 ctime; // 成为城市节点的时间戳
        bool firstMonthReturnEarnest; // 第1个月是否退了(发放到可提现余额)保证金
        bool secondMonthReturnEarnest; // 第2个月是否退了(发放到可提现余额)保证金
        bool thirdMonthReturnEarnest; // 第3个月是否退了(发放到可提现余额)保证金
        uint256 cityLevel; // 所在城市等级
        bool assessmentMonthStatus; // 按月考核状态
        bool assessmentStatus; // 最终考核状态
        bool returnSuretyStatus; // 保证金退还状态
        bool returnSuretyRate; // 保证金退还比例
    }
```

#### 查询先锋城市考核标准

- assessmentCriteria 参数1 为城市等级，2 为月份，均为整数

#### 查看收益领取状态

- 根据先锋地址查询用户领取状态，如果可领取余额为0，就为不可领取状态，否则查询状态
    - benefitPackageRewardStatus; // 用户福袋奖励提取状态
    - fundsRewardStatus; // 用户社交基金奖励提取状态
    - delegateRewardStatus; // 用户新增质押奖励提取状态