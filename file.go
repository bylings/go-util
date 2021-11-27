package util

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // 时间戳,纳秒
	// time.Now().Unix() 		// unix 时间戳，秒
	// time.Now().UnixMilli()	// unix 时间戳，毫秒
	// time.Now().UnixMicro()	// unix 时间戳，微秒
	// time.Now().UnixNano()	// unix 时间戳，纳秒
	//fmt.Println("%v", time.Now().Unix())
	//fmt.Println("%v", time.Now().UnixMilli())
	//fmt.Println("%v", time.Now().UnixMicro())
	//fmt.Println("%v", time.Now().UnixNano())
	/**
	%v 1632618739
	%v 1632618739036
	%v 1632618739036173
	%v 1632618739036176000
	*/
	//time.Sleep(1e9)	//表示延迟1秒  1e9 = time.Second
	//time.Sleep(time.Second)		// 表示延迟1秒
}

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
			//fmt.Println("文件路径创建失败", err)
			return false, errors.New("文件路径创建失败")
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
		//fmt.Println("文件写入失败", err)
		return false, errors.New("文件写入失败")
	}
	// 刷新缓冲区
	errs := outputWriter.Flush()
	if errs != nil {
		//fmt.Println("文件刷新失败", errs)
		return false, errors.New("文件刷新失败")
	}
	return true, nil
}

// 根据文件获取文件大小，单位/字节
func GetFileSize(filename string) (int64, error) {
	// 1、打开文件
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// 添加一个延迟执行释放文件句柄
	defer f.Close()
	fi, err := f.Stat() // 获得文件的元数据信息
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil // 单位，字节
}

// 获取文件内容
func GetFileContent(filename string) (string, error) {
	f, err := os.Open(filename)
	// 判断是否是一个文件
	if err != nil {
		return "", err
	}
	defer f.Close()
	bytes, err := ioutil.ReadFile(filename)
	return string(bytes), nil
}
