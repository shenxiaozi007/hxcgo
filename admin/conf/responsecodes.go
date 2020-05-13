package conf

var ResponseCodes = map[string]int{
	"success":          0,    //成功
	"err_request":      4000, //请求失败
	"err_param":        4100, //签名错误
	"err_invalid_id":   4101, //无效ID
	"err_invalid_name": 4102, //无效的名称

	"err_sign":                 4201, //参数错误
	"err_captcha":              4202, //验证码错误
	"err_username_or_password": 4203, //用户名或密码错误
	"err_restricted_admin":     4204, //限制的用户

	"err_upload":     4211, //上传失败
	"err_limit_size": 4212, //文件大小超出最大限制
	"err_mime_type":  4213, //无法识别MIMIE类型

	"err_role":            4221, //请选择角色
	"err_role_not_exists": 4222, //角色不存在
	"err_invalid_role_id": 4223, //无效角色ID

	"err_group":            4231, //请选择分组
	"err_group_not_exists": 4232, //分组不存在

	"err_admin_name":               4241, //帐号至少5个字符，不能单纯数字
	"err_admin_name_exists":        4242, //帐号已经存在
	"err_admin_name_can_not_email": 4243, //账户名称不能是邮箱格式

	"err_mobile":           4251, //手机号码格式错误
	"err_mobile_exists":    4252, //手机号码已经存在
	"err_email":            4261, //邮箱格式错误
	"err_email_exists":     4262, //邮箱已经存在
	"err_password":         4271, //密码错误，密码至少8个字符，不能纯数字或字母
	"err_confirm_password": 4272, //密码错误，两次输入密码不一致

	"err_invalid_privilege_id": 4281, //无效的权限ID

	"err_category_id":     4291, //请选择分类
	"err_wholesale_price": 4292, //请填写批发价
	"err_title":           4293, //请填写标题
	"err_cover_image":     4294, //请选择封面图
	"err_product_image":   4295, //请上传商品图片

	"err_response": 4999, //错误的响应
}
