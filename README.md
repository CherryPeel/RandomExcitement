# RandomExcitement

> v1.0.0

RandomExcitement是一个用go语言开发的随机图片的api，它可以让您轻松地获取不同大小和类别的随机图片，用于测试、设计或娱乐等目的。您只需要发送一个
简单的http请求，就可以得到一张高质量的随机图片，或者一个指向该图片的url。您可以自由地指定图片的宽度、高度、格式、颜色、模糊度等参数，或者让api
为您随机选择。您还可以从多种类别中选择图片，如动物、食物、风景、人物等，或者让api为您随机选择。 RandomExcitement是一个开源的项目，遵循
Apache 2.0许可协议，您可以在github上查看和下载源代码，并对项目进行修改和分发。如果您喜欢RandomExcitement，请给我们一个星星或者留下您的评论
和反馈，我们将不断改进和完善我们的项目。

## 功能

- [x] 随机图片
- [x] 指定图片类别
- [ ] 指定图片大小
- [ ] 图片裁剪

## 快速开始

1. 下载项目到本地

    ```gitexclude
    git clone https://github.com/CherryPeel/RandomExcitement.git
    ```

2. 进入项目目录

    ```shell
    cd RandomExcitement
    ```

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

