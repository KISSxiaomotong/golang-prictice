//字符串
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string = "This is hello world!"

	//判断字符串以this开头
	var prefix = strings.HasPrefix(str, "this")

	fmt.Printf("str 是以this开头的")
	fmt.Printf("%t\n", prefix)

	//判断字符串以！结尾
	var end = strings.HasSuffix(str, "!")
	fmt.Printf("str 是以!结尾的")
	fmt.Printf("%t\n", end)

	//字符串包含关系
	var contains = strings.Contains(str, "world")
	fmt.Printf("str 是否包含world")
	fmt.Printf("%t\n", contains)

	//查询字符在字符串中第一次出现的位置
	var index = strings.Index(str, "e")
	fmt.Printf("e在字符串str中第一次出现的位置是")
	fmt.Printf("%d\n", index)

	//查询字符在字符串中最后一次出现的位置
	var lastIndex = strings.LastIndex(str, "h")
	fmt.Printf("h在字符串str中最后一次出现的位置是")
	fmt.Printf("%d\n", lastIndex)

	//字符串替换
	var newStr = strings.Replace(str, "!", ".", 10)
	fmt.Printf("替换后的新字符串为")
	fmt.Printf("%s\n", newStr)

	//统计字符在字符串中出现的次数
	var count = strings.Count(str, "h")
	fmt.Printf("h在字符串str中出现的次数为")
	fmt.Printf("%d\n", count)

	//重复输出字符串
	var repeatString = strings.Repeat(str, 2)
	fmt.Printf("%s\n", repeatString)

	//字符串转换大写
	var lowStr = strings.ToLower(str)
	fmt.Printf("字符串str转换为小写")
	fmt.Printf("%s\n", lowStr)

	//字符串转换大写
	var upStr = strings.ToUpper(str)
	fmt.Printf("字符串str转换为大写")
	fmt.Printf("%s\n", upStr)

	//分割字符串
	var strSlice = strings.Fields(str)
	var strSplit = strings.Split(str, "o")
	fmt.Printf("分割字符串str: %v\n", strSlice)
	fmt.Printf("分割字符串str: %v\n", strSplit)

	//组合字符串
	var strJoin = strings.Join(strSplit, "-")
	fmt.Printf("组合字符串str: %s\n", strJoin)

	//字符串转换为浮点数
	var numberStr string = "43.96"
	//多返回值（第一个为返回的结果，第二个为有可能抛出的错误）
	var num, _ = strconv.ParseFloat(numberStr, 64)
	fmt.Printf("字符串转换为数字: %f\n", num)
}
