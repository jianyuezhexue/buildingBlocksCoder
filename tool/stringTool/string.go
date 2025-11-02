package stringTool

import (
	"strings"
	"unicode"
)

// 将字符串转大驼峰 | 入参可能是蛇形写法，大驼峰，小驼峰
func ToUpperCamelCase(input string) string {
	res := ""

	// 判空
	if input == "" {
		return res
	}

	// 使用空格替换所有非字母数字字符，然后按空格分割
	parts := strings.FieldsFunc(input, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	for _, part := range parts {
		if len(part) > 0 {
			// 首字母大写，其余保持原样
			runes := []rune(part)
			res += strings.ToUpper(string(runes[0])) + string(runes[1:])
		}
	}

	return res
}

// 将字符串转小驼峰 | 入参可能是蛇形写法，大驼峰，小驼峰
func ToLowerCamelCase(input string) string {
	res := ""

	// 判空
	if input == "" {
		return res
	}

	// 使用空格替换所有非字母数字字符，然后按空格分割
	parts := strings.FieldsFunc(input, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	for i, part := range parts {
		if len(part) > 0 {
			// 首字母大写，其余保持原样
			runes := []rune(part)
			if i == 0 {
				res += strings.ToLower(string(runes[0])) + string(runes[1:])
			} else {
				res += strings.ToUpper(string(runes[0])) + string(runes[1:])
			}
		}
	}

	return res
}
