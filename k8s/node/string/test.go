package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"k8s.io/klog/v2"
)

func main() {
	diffNodeInfo := make([]string, 0)
	dir, _ := os.Getwd()
	// fmt.Println(dir)
	k8sNodePath := dir + "/basic/string/k8sNode.txt"
	k8sNodeInfo := readNodeInfo(k8sNodePath)
	totalNodePath := dir + "/basic/string/totalNode.txt"
	totalNodeInfo := readNodeInfo(totalNodePath)

	for key, _ := range totalNodeInfo {
		if _, ok := k8sNodeInfo[key]; !ok {
			diffNodeInfo = append(diffNodeInfo, key)
		}
	}
	klog.Infoln(len(diffNodeInfo))
	// fmt.Println(len(diffNodeInfo))
	klog.Infoln(diffNodeInfo)
	// fmt.Println(diffNodeInfo)

}

func readNodeInfo(path string) map[string]int {
	nodeInfo := make(map[string]int)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("无法打开文件:", err)
	}
	defer file.Close()

	// 创建Scanner对象
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		line = strings.ToUpper(line)
		nodeInfo[line] = 1
	}

	klog.Infoln(len(nodeInfo))
	// fmt.Println(len(nodeInfo))
	// fmt.Println(nodeInfo)

	return nodeInfo
}
