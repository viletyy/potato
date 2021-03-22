package e

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",

	ERROR_EXIST_USER : "已存在该用户",
	ERROR_NOT_EXIST_USER : "该用户不存在",
	ERROR_EXIST_VENDOR : "已存在该系统厂商",
	ERROR_NOT_EXIST_VENDOR : "该系统厂商不存在",
	ERROR_EXIST_BUSINESS : "已存在该业务系统",
	ERROR_NOT_EXIST_BUSINESS : "该业务系统不存在",
	ERROR_EXIST_META_DATABASE : "已存在该数据源",
	ERROR_NOT_EXIST_META_DATABASE : "该数据源不存在",
	ERROR_EXIST_META_TABLE : "已存在该元数据",
	ERROR_NOT_EXIST_META_TABLE : "该元数据不存在",

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