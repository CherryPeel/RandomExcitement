# RandomExcitement

> v1.0.0

[English](./docs/README-EN.md) | 中文

RandomExcitement 是一个使用 Go 语言开发的随机图片 API，它为您提供便捷的方式获取不同尺寸和类别的随机图片，满足测试、设计和娱乐等多种需求。通过
简单的HTTP 请求，您即可获得高质量的随机图片，或者获取指向该图片的 URL。您可以灵活地设定图片的宽度、高度、格式、颜色、模糊程度等参数，或者让API
自动随机选择。

随机图片的类别可以自定义，您可以根据需求手动选择特定类别，或者由 API 自动随机挑选。RandomExcitement 以开源项目的形式发布，采用 Apache 2.0 
许可协议，您可以在 GitHub 上查阅和下载源代码，并进行自定义修改和分发。

我们诚挚地欢迎您为 RandomExcitement 提供支持，您可以给予我们一个星标、留下宝贵的评论和反馈。我们将持续不断地改进和优化项目，为用户提供更好的
体验。感谢您的关注和参与！

## 项目地址

[https://github.com/CherryPeel/RandomExcitement](https://github.com/CherryPeel/RandomExcitement)

## 功能

- [x] 随机图片
- [x] 指定图片类别
- [ ] 指定图片大小
- [ ] 图片裁剪

## 快速开始

### 个人服务器使用

1. 下载项目到本地

    ```gitexclude
    git clone https://github.com/CherryPeel/RandomExcitement.git
    ```

2. 进入项目目录

    ```shell
    cd RandomExcitement
    ```
   在`config/site.json`中配置你的站点域名，例如`http://localhost:8080/`

3. 运行项目
    - 快速运行

    ```shell
    go run cmd/main.go
    ```

    - 构建项目运行
        1. 构建项目

        ```shell
          go build -tags netgo -ldflags '-s -w' -o app cmd/main.go
        ```

        2. 运行

        ```shell
          ./app
        ```

### Render 部署


## 接口文档

或者运行此项目后访问`http://localhost:8080`查看接口文档

### 初始化

#### POST 刷新分类

POST /v1/refresh

一旦你整理好图片文件夹，执行这个接口，会在`configs`文件夹下生成一个名为`classification.json`的配置文件。该文件会记录`static/images`文件夹中的所有图片以及它们所属的分类信息。

> 返回示例

> 成功

```json
{
  "code": 200,
  "msg": "success",
  "data": null
}
```

##### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

##### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选 | 约束 | 中文名 | 说明 |
| ------ | ------- | ---- | ---- | ------ | ---- |
| » code | integer | true | none |        | none |
| » msg  | string  | true | none |        | none |
| » data | null    | true | none |        | none |

### 随机图片

#### GET 返回所有的图片类别

GET /v1/allClassifies

此接口将会返回你的所有图片类别

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": [
    "string"
  ],
  "msg": "string"
}
```

##### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

##### 返回数据结构

状态码 **200**

| 名称   | 类型     | 必选 | 约束 | 中文名 | 说明 |
| ------ | -------- | ---- | ---- | ------ | ---- |
| » code | integer  | true | none |        | none |
| » data | [string] | true | none |        | none |
| » msg  | string   | true | none |        | none |

#### GET 获取随机图

GET /v1/randomExcitement

从所有类别中随机重定向到一个图片链接

> 返回示例

> 200 Response

##### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

##### 返回数据结构

#### POST 获取随机图

POST /v1/randomExcitement

从所有类别中以json格式随机返回一个图片链接

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": "string",
  "msg": "string"
}
```

##### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

##### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选 | 约束 | 中文名 | 说明 |
| ------ | ------- | ---- | ---- | ------ | ---- |
| » code | integer | true | none |        | none |
| » data | string  | true | none |        | none |
| » msg  | string  | true | none |        | none |

#### POST 根据用户需求返回

POST /v1/random

将会生成根据用户所请求的类别以json格式返回对应的随机图链接

##### 请求参数

| 名称     | 位置  | 类型   | 必选 | 说明 |
| -------- | ----- | ------ | ---- | ---- |
| classify | query | string | 是   | none |

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": "string",
  "msg": "string"
}
```

##### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

##### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选 | 约束 | 中文名 | 说明 |
| ------ | ------- | ---- | ---- | ------ | ---- |
| » code | integer | true | none |        | none |
| » data | string  | true | none |        | none |
| » msg  | string  | true | none |        | none |

#### GET 根据用户需求返回

GET /v1/random

将会根据用户所请求的类别随机重定向到一个图片链接

##### 请求参数

| 名称     | 位置  | 类型   | 必选 | 说明 |
| -------- | ----- | ------ | ---- | ---- |
| classify | query | string | 是   | none |

> 返回示例

> 200 Response

##### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

##### 返回数据结构

## 鸣谢

<div style="display: flex; flex-direction: row;">
       <div style="margin-right: 10px;">
           <a href="https://img.shields.io/badge/gin-v1.6.3-green">
               <img src="https://img.shields.io/badge/Gin-v1.6.3-green" alt="Gin">
           </a>
       </div>
       <div style="margin-right: 10px;">
           <a href="https://img.shields.io/badge/GoLand-2020.2.3-green">
               <img src="https://img.shields.io/badge/GoLand-2020.2.3-green" alt="GoLand">
           </a>
       </div>
       <div style="margin-right: 10px;">
           <a href="https://img.shields.io/badge/Go-v1.20.0-green">
               <img src="https://img.shields.io/badge/Go-v1.20.0-green" alt="Go">
           </a>
       </div>
     <div style="margin-right: 10px;">
        <a href="https://img.shields.io/badge/Apifox-v2.3.11-green">
            <img src="https://img.shields.io/badge/Apifox-v2.3.11-green" alt="Apifox">
        </a>
      </div>
      <div style="margin-right: 10px;">
        <a href="https://img.shields.io/badge/license-Apache%202-blue">
            <img src="https://img.shields.io/badge/license-Apache%202-blue" alt="Apache 2.0">
        </a>
      </div>
</div>

