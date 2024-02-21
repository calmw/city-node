## 生成绑定文件

- 生成绑定go文件命令

#### 生成userLocation合约代码

- abigen --abi ../abi/IntoUserLocation.json --pkg intoCityNode --type UserLocation --out ./binding/userLocation.go

#### 生成cityPioneer合约代码

- abigen --abi ../abi/IntoCityPioneer.json --pkg intoCityNode --type CityPioneer --out ./pkg/binding/intoCityNode/cityPioneer.go

#### 生成city合约代码

- abigen --abi ../abi/IntoCity.json --pkg intoCityNode --type City --out ./pkg/binding/intoCityNode/city.go

#### 生成appraise合约代码

- abigen --abi ../abi/IntoAppraise.json --pkg intoCityNode --type Appraise --out ./pkg/binding/intoCityNode/appraise.go

#### 生成cityPioneerData合约代码

- abigen --abi ../abi/IntoCityPioneerData.json --pkg intoCityNode --type CityPioneerData --out ./pkg/binding/intoCityNode/cityPioneerData.go
