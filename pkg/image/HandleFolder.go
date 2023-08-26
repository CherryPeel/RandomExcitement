package image

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// GetImageFromFolder 从文件夹中获取图片
func GetImageFromFolder(folderPath string) ([]byte, error) {
	imgClassMap := make(map[string][]string)

	// 从根文件夹中获取图片
	filepath.Walk(folderPath, func(_path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isImageFile(_path) {
			imgPath := _path
			imgClassMap["unknown"] = append(imgClassMap["unknown"], imgPath)
		} else if info.IsDir() && _path != folderPath {
			imgPaths, err := getValidImagePaths(_path)
			if err != nil {
				return err
			}

			imgClassMap[info.Name()] = imgPaths

			// 遍历子文件夹时阻止进入
			return filepath.SkipDir
		}
		return nil
	})

	var imgCates []Cate
	for imgClass, imgPaths := range imgClassMap {
		imgCates = append(imgCates, Cate{
			ImgClass: imgClass,
			ImgPath:  imgPaths,
		})
	}

	// 将分类信息转换为 JSON 格式
	jsonData, err := json.Marshal(imgCates)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

// GenerateClassifyJsonFile 生成分类信息的 JSON 文件
func GenerateClassifyJsonFile(folderPath string) error {
	jsonData, err := GetImageFromFolder(folderPath)
	if err != nil {
		return err
	}

	// 创建 config 文件夹
	configFolderPath := filepath.Join("", "configs")
	if err := os.MkdirAll(configFolderPath, os.ModePerm); err != nil {
		return err
	}

	// 将 JSON 数据写入测试文件
	jsonFilePath := filepath.Join(configFolderPath, "classification.json")
	if err := os.WriteFile(jsonFilePath, jsonData, 0644); err != nil {
		return err
	}
	return nil
}

// getValidImagePaths 获取文件夹中的图片路径
func getValidImagePaths(folderPath string) ([]string, error) {
	var imgPaths []string

	filepath.WalkDir(folderPath, func(_path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isImageFile(_path) {
			imgPaths = append(imgPaths, _path)
		}
		return nil
	})

	return imgPaths, nil
}

// isImageFile 判断文件是否为图片
func isImageFile(filename string) bool {
	ext := filepath.Ext(filename)
	return ext == ".jpg" || ext == ".png" || ext == ".jpeg"
}
