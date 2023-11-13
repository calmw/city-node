package statecode

const (
	// LangZh language
	LangZh   = 111
	LangEn   = 112
	LangZhTw = 113

	// CommonSuccess common
	CommonSuccess      = 0
	CommonErrServerErr = 1000
	ParameterEmptyErr  = 1001

	UserEmpty    = 1201 //p_name empty
	PioneerEmpty = 1202 //p_name empty

	SyncingData = 1301 // sync data

)

var Msg = map[int]map[int]string{

	0: {
		LangZh:   "成功",
		LangZhTw: "成功",
		LangEn:   "success",
	},
	1000: {
		LangZh:   "服务器繁忙，请稍后重试",
		LangZhTw: "服務器繁忙，請稍後重試",
		LangEn:   "server is busy, please try again later",
	},
	1001: {
		LangZh:   "参数不能为空",
		LangZhTw: "参数不能為空",
		LangEn:   "parameter is empty",
	},
	1101: {
		LangZh:   "token 不能为空",
		LangZhTw: "token 不能為空",
		LangEn:   "token required",
	},
	1102: {
		LangZh:   "token错误",
		LangZhTw: "token錯誤",
		LangEn:   "token invalid",
	},
	1201: {
		LangZh:   "user 不能为空",
		LangZhTw: "user 不能為空",
		LangEn:   "user required",
	},
	1202: {
		LangZh:   "pioneer 不能为空",
		LangZhTw: "pioneer 不能為空",
		LangEn:   "pioneer required",
	},
	1203: {
		LangZh:   "chain_id 不能为空",
		LangZhTw: "chain_id 不能為空",
		LangEn:   "chain_id required",
	},
	1301: {
		LangZh:   "数据获取中，请稍后再试",
		LangZhTw: "数据获取中，请稍后再试",
		LangEn:   "syncing data, please waite a minute",
	},
}

func GetMsg(c int, lang int) string {
	_, ok := Msg[c]
	if ok {
		msg, ok := Msg[c][lang]
		if ok {
			return msg
		}
		return Msg[CommonErrServerErr][lang]
	}
	return Msg[CommonErrServerErr][lang]
}
