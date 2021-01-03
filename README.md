# goSslNotAfter
使用golang检测ssl过期时间

# 使用acme.sh生成免费证书,每3个月自动更新一次(实际2月就会更新)
# https://github.com/acmesh-official/acme.sh

功能:
1. 检测指定url的证书有效期是否小于10天
2. 如小于10天则邮件报警

使用方法:
 
