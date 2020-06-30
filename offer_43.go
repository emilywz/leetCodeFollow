package main

//输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。
//
//例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。
//
//示例 1：
//输入：n = 12
//输出：5
//示例 2：
//输入：n = 13
//输出：6

//设置当前位为current，高位为high，低位为low；
//可以总结出某位上出现1的次数的相关公式(i为10，因为当前位是十位)：
//如果current=0；则count = high * i
//如果current=1；则count = high * i + low + 1
//如果current>1；则count = (high+1) * i

func countDigitOne(n int) int {
	if n < 1 {
		return 0
	}
	base := 1
	high, current, low, result := n, 0, 0, 0
	for high > 0 {
		high /= 10
		current = (n / base) % 10
		low = n - n/base*base

		if current == 0 {
			result = result + high*base
		} else if current == 1 {
			result = result + high*base + low + 1
		} else {
			result = result + (high+1)*base
		}
		base *= 10
	}
	return result
}
