package e

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
	ERROR_EXIST_GAME : "已存在该游戏",
	ERROR_EXIST_TEAM : "已存在该队伍",
	ERROR_NOT_EXIST_GAME : "该游戏不存在",
	ERROR_NOT_EXIST_TEAM : "该队伍不存在",
	ERROR_NOT_EXIST_LEAGUE : "该联赛不存在",
	ERROR_NOT_EXIST_SERIES : "该系列赛不存在",

	ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
	ERROR_AUTH_TOKEN : "Token生成失败",
	ERROR_AUTH : "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}