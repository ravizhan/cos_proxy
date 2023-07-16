# cos_proxy

## 功能
- 在腾讯云服务器上反代COS对象存储，实现免流
- 支持私有读
- 支持数据处理

## 使用方法
### 下载预构建文件 or 源码编译
- 预构建文件可前往 https://github.com/ravizhan/cos_proxy/releases 下载


- 源码编译
```shell
git clone github.com/ravizhan/cos_proxy
cd cos_proxy
go build
```
### 运行
第一次运行会生成配置文件，请将每一项填写完整后再运行

| 字段        | 释义                                                     |
|-----------|--------------------------------------------------------|
| BucketUrl | 访问域名 格式: <BucketName-APPID>.cos.\<Region>.myqcloud.com |
| SecretId  | 用户的 SecretId,建议使用子账号密钥                                 |
| SecretKey | 用户的 SecretKey,建议使用子账号密钥                                |
| Suffix    | 文件参数 例: /image.jpg\<Suffix>                            |
| Port      | 监听端口                                                   |

## 后记
一时心血来潮写的。

正好最近在学GO，浅浅试个水。

本来是自己用的，但想了想还是开源了，虽然可能没什么人用就是了。

嗯，就这样