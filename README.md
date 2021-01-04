# goSslNotAfter
使用golang检测ssl过期时间

## 说明
```
一直使用acme.sh生成免费证书,每3个月自动更新一次(实际2月就会更新)
# https://github.com/acmesh-official/acme.sh

为避免自动更新失败或其它原因造成ssl证书过期,主动检测ssl过期时间;

功能:
1. 检测指定url的证书有效期是否小于10天
2. 如小于10天则邮件报警


1. 下载代码
   git clone https://github.com/WPFS/goSslNotAfter.git
   cd goSslNotAfter
2. 修改config.json文件
   a) 邮件服务器信息
   b) 发件人信息
   c) 收件人信息(多个收件人以";"分割)
3. 修改main.go中domain变量为要检测的网站
   如:domains := []string{"https://www.baidu.com", "https://www.aliyun.com"}
4. go build 
5. 添加至计划任务
```

## ssl邮件发送参考大神的写法
https://my.oschina.net/u/3768573/blog/1607327
