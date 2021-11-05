package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ReadFileFun func(a interface{})

// 读取json文件
func ReadJson(filepath string, a interface{}) error {
	// 读取文件
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("读取json文件失败", err)
		return err
	}
	// 解析数据
	err = json.Unmarshal(bytes, a)
	if err != nil {
		fmt.Println("解析数据失败", err)
		return err
	}
	return nil
}

// 判断文件是否存在
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	return false, nil
}

// 写入文件
func WriteFile(path, file, data string) (bool, error) {
	// 1、判断路径是否存在，不存在则创建
	if exist, _ := PathExist(path); !exist {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Println("文件路径创建失败", err)
			return false, nil
		}
	}
	// 2、打开文件资源,设置可写并创建
	outputFile, outputError := os.OpenFile(path+file, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		return false, outputError
	}
	// 释放文件资源
	defer outputFile.Close()
	// 创建写缓冲区
	outputWriter := bufio.NewWriter(outputFile)

	// 写入文件
	_, err := outputWriter.WriteString(data)
	if err != nil {
		fmt.Println("文件写入失败", err)
		return false, nil
	}
	// 刷新缓冲区
	errs := outputWriter.Flush()
	if errs != nil {
		fmt.Println("文件刷新失败", errs)
		return false, errs
	}
	return true, nil
}
