import base64
import pyDes
import sha3
import sys

country_code = sys.argv[1]
city_code = sys.argv[2]
a = '0	中国	277	宿迁市'

country_code = int(country_code)

city_code = int(city_code)

message = "{},{}".format(country_code, city_code).encode()

k = sha3.keccak_256()
k.update(message)
print('不可逆：', '0x' + k.hexdigest())

keys = 'gZAnoptLJm6GYXdClPhIMfo6'
k = pyDes.triple_des(keys, pyDes.ECB, "\0\0\0\0\0\0\0\0", pad=None, padmode=pyDes.PAD_PKCS5)
d = k.encrypt(message)
print("可逆：", base64.b64encode(d).decode())
