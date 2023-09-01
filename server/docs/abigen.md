## 生成绑定文件

- 生成绑定go文件命令

#### 生成userLocation合约代码

- abigen --abi ../abi/IntoUserLocation.json --pkg intoCityNode --type UserLocation --out ./binding/userLocation.go

#### 生成cityPioneer合约代码

- abigen --abi ../abi/IntoCityPioneer.json --pkg intoCityNode --type CityPioneer --out ./binding/cityPioneer.go