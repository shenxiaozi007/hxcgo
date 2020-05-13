package reg

import "regexp"

//验证手机
func Mobile(s string) bool {
	r := regexp.MustCompile(`^(0|86|17951)?(13[0-9]|15[012356789]|1[78][0-9]|14[57])[0-9]{8}$`)
	return r.MatchString(s)
}

//验证数字
func Number(s string) bool {
	r := regexp.MustCompile(`^[0-9]+$`)
	return r.MatchString(s)
}

//验证字母
func Alphabet(s string) bool {
	r := regexp.MustCompile(`^[A-Za-z]+$`)
	return r.MatchString(s)
}

//验证年份 格式：yyyy
func Year(s string) bool {
	r := regexp.MustCompile(`^(\d{4})$`)
	return r.MatchString(s)
}

//验证月份 格式：mm
func Month(s string) bool {
	r := regexp.MustCompile(`^0?([1-9])$|^(1[0-2])$`)
	return r.MatchString(s)
}

//验证日期 格式：yyyy-mm-dd
func Date(s string) bool {
	r := regexp.MustCompile(`^(\d{4})-(0?\d{1}|1[0-2])-(0?\d{1}|[12]\d{1}|3[01])$`)
	return r.MatchString(s)
}

//日期时间 格式：yyyy-mm-dd hh:ii:ss
func DateTime(s string) bool {
	r := regexp.MustCompile(`^(\d{4})-(0?\d{1}|1[0-2])-(0?\d{1}|[12]\d{1}|3[01])\s(0\d{1}|1\d{1}|2[0-3]):[0-5]\d{1}:([0-5]\d{1})$`)
	return r.MatchString(s)
}

//验证邮箱
func Email(s string) bool {
	r := regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	return r.MatchString(s)
}

//验证邮编
func Postcode(s string) bool {
	r := regexp.MustCompile(`[1-9]\d{5}(?!\d)`)
	return r.MatchString(s)
}

//验证URL
func URL(s string) bool {
	r := regexp.MustCompile(`\b(([\w-]+:\/\/?|www[.])[^\s()<>]+(?:\([\w\d]+\)|([^[:punct:]\s]|\/)))`)
	return r.MatchString(s)
}

//验证身份证
func Identify(s string) bool {
	r := regexp.MustCompile(`(^\d{15}$)|(^\d{17}([0-9]|X)$)`)
	return r.MatchString(s)
}

func IPv4(s string) bool {
	r := regexp.MustCompile(`^(((\d{1,2})|(1\d{2})|(2[0-4]\d)|(25[0-5]))\.){3}((\d{1,2})|(1\d{2})|(2[0-4]\d)|(25[0-5]))$`)
	return r.MatchString(s)
}

func IPv6(s string) bool {
	r := regexp.MustCompile(`^([\da-fA-F]{1,4}:){7}[\da-fA-F]{1,4}$`)
	return r.MatchString(s)
}
