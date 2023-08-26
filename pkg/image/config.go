package image

import (
	"encoding/json"
	"os"
)

const ImgJsonPath = "configs/classification.json"

// Cate 分类
type Cate struct {
	ImgClass string
	ImgPath  []string
}

// Domain 域名
type Domain struct {
	Domain string `json:"domain"`
}

// GetDomain 获取json文件中的域名信息
func GetDomain() string {
	data, err := os.ReadFile("configs/site.json")
	if err != nil {
		panic(err)
	}

	var domain Domain
	err = json.Unmarshal(data, &domain)
	if err != nil {
		panic(err)
	}

	return domain.Domain
}
