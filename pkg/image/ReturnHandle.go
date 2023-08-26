package image

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strings"
)

// GenerateAllClassify 获取所有分类
func GenerateAllClassify() ([]string, error) {
	cate, err := readJsonFile()
	if err != nil {
		return nil, err
	}
	var allClassify []string
	for _, v := range cate {
		allClassify = append(allClassify, v.ImgClass)

	}
	return allClassify, nil
}

// RandomImgByReq 根据请求返回一张图片
func RandomImgByReq(classify string) (string, error) {
	cate, err := readJsonFile()
	if err != nil {
		return "", err
	}

	// 判断是否存在该分类,存在则返回该分类下的随机一张图片
	var classifyPath []string
	if contains(cate, classify) {
		for _, v := range cate {
			if v.ImgClass == classify {
				classifyPath = append(classifyPath, v.ImgPath...)
				break
			}
		}
		reqClassifyLength := len(classifyPath)
		randomIndex := rand.Intn(reqClassifyLength)
		randomImage := classifyPath[randomIndex]
		randomImage = path.Clean(randomImage)
		randomImage = strings.ReplaceAll(randomImage, "\\", "/")

		//ImgUrl := "http://localhost:8080/" + randomImage
		domain := GetDomain()
		ImgUrl := domain + randomImage
		return ImgUrl, nil
	} else {
		// 不存在则报错
		//for _, v := range cate {
		//	classifyPath = append(classifyPath, v.ImgPath...)
		//}
		//reqClassifyLength := len(classifyPath)
		//randomIndex := rand.Intn(reqClassifyLength)
		//randomImage := classifyPath[randomIndex]
		//return randomImage, nil
		return "", fmt.Errorf("分类不存在")
	}
}

// RandomImg 随机返回一张图片
func RandomImg() (string, error) {
	cate, err := readJsonFile()
	if err != nil {
		return "", err
	}
	var allClassify []string
	for _, v := range cate {
		allClassify = append(allClassify, v.ImgClass)

	}
	reqClassifyLength := len(allClassify)
	randomIndex := rand.Intn(reqClassifyLength)
	randomClassify := allClassify[randomIndex]

	// 返回该分类下的随机一张图片
	var classifyPath []string
	for _, v := range cate {
		if v.ImgClass == randomClassify {
			classifyPath = append(classifyPath, v.ImgPath...)
			break
		}
	}

	reqClassifyLength = len(classifyPath)
	randomIndex = rand.Intn(reqClassifyLength)
	randomImage := classifyPath[randomIndex]
	randomImage = path.Clean(randomImage)
	randomImage = strings.ReplaceAll(randomImage, "\\", "/")

	domain := GetDomain()
	//ImgUrl := "http://localhost:8080/" + randomImage
	ImgUrl := domain + randomImage
	return ImgUrl, nil
}

// contains 判断所有分类中是否包含用户所请求的分类
func contains(s []Cate, str string) bool {
	for _, v := range s {
		if v.ImgClass == str {
			return true
		}
	}
	return false
}

// readJsonFile 读取 JSON 文件
func readJsonFile() ([]Cate, error) {
	jsonFile, err := os.ReadFile(ImgJsonPath)
	if err != nil {
		return nil, err
	}
	var cate []Cate
	err = json.Unmarshal(jsonFile, &cate)
	if err != nil {
		return nil, err
	}
	return cate, nil
}
