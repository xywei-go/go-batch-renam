package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// step1. 选择操作类型：D-批量命名目录下的文件，F-只命名文件，Q-退出
// step2. 输入路径，
func main() {
	fmt.Println("start rename app")
	fmt.Println("please select operation type:")
	fmt.Println("D-dir, F-file, Q-exit")
	var optType string
	fmt.Scanln(&optType)
	if !strings.EqualFold(optType, "D") && !strings.EqualFold(optType, "F") && !strings.EqualFold(optType, "Q") {
		fmt.Println("input ivalid! please check your input!")
		return
	}

	fmt.Println("you select: ", optType)
	fmt.Println("please input the absolute path: ")
	var pathInput string
	fmt.Scanln(&pathInput)
	if len(pathInput) == 0 {
		fmt.Println("no path input, exit now......")
		return
	}
	fmt.Printf("input path is: %s\n", pathInput)
	var replaceWord string
	fmt.Println("please input the word to be deleted name: ")
	fmt.Scanln(&replaceWord)
	fmt.Println("you input desName: ", replaceWord)
	targetFile, err := os.Stat(pathInput)
	// 检查路径是否存在
	if err != nil {
		fmt.Println("path invalid!", err)
		return
	}
	// 是文件直接重命名
	if !targetFile.IsDir() && strings.EqualFold(optType, "F") {
		renameFile(pathInput, replaceWord)
		return
	}
	// 是目录就遍历当前目录的文件，然后重命名
	if targetFile.IsDir() && strings.EqualFold(optType, "D") {
		// 遍历目录
		// renameFile(pathInput, replaceWord)
		fis, err2 := ioutil.ReadDir(pathInput)
		if err2 != nil {
			fmt.Println(fis)
			return
		}
		for _, v := range fis {
			absPath := pathInput + string(os.PathSeparator) + v.Name()
			fi, _ := os.Stat(absPath)
			if !fi.IsDir() {
				fmt.Printf("v: %v\n", absPath)
				renameFile(absPath, replaceWord)
			}

		}
		fmt.Println("Done!")
		return
	}
	// 结束
	fmt.Println("do nothind!")
}

func renameFile(originName, replaceWord string) {
	// 获取文件后缀
	extName := path.Ext(originName)
	// 获取文件名，不包括后缀名
	fileNameOld := strings.TrimSuffix(path.Base(originName), extName)
	// 新文件名，不包括后缀名
	fileNameNew := strings.ReplaceAll(fileNameOld, replaceWord, "")
	// 新文件全路径
	absNewFile := fileNameNew + extName
	err2 := os.Rename(originName, absNewFile)
	if err2 != nil {
		fmt.Println("error! ", err2)
		return
	}
	fmt.Println(originName, " ---rename to--- ", absNewFile, " ---ok--- ")
}
