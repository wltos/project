package utils

import "github.com/axgle/mahonia"

// ConvertToString 转码工具
func ConvertToString(srcText, srcFormat, tagFormat string) string {

	srcCoder := mahonia.NewDecoder(srcFormat)
	srcResult := srcCoder.ConvertString(srcText)

	tagCoder := mahonia.NewDecoder(tagFormat)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)

	tagText := string(cdata)
	return tagText
}
