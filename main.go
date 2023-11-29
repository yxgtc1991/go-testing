package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
}

/*
1、return 最先执行，负责将结果写入返回值
2、接着 defer 执行收尾工作
3、最后函数携带当前返回值退出
*/
func getNum(i int) int {
	return i
	defer fmt.Println(i)
	return i
}

/*
1、return 最先执行，负责将结果写入返回值
2、接着 defer 执行收尾工作
3、最后函数携带当前返回值退出
*/
func test() error {
	client, err := getClient()
	defer close(client)
	if err != nil {
		return err
	}
	return nil
}

func getClient() (string, error) {
	return "", errors.New("test")
}

func close(client string) {
	fmt.Println("ok")
}
