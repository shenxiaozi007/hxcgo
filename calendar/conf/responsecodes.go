package conf

var ResponseCodes = map[string]int{
	"success":        0,    //成功
	"err_request":    4000, //请求失败
	"err_param":      4100, //签名错误
	"err_sign":       4201, //参数错误
	"err_code":       4202, //无效code值
	"err_product_id": 4210, //错误产品ID

	"err_response": 4999, //错误的响应
}
