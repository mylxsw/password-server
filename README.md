# password-server

一个用于生成随机密码的 web 服务器。

只有两个接口

- "/": 生成随机密码，支持参数 length-密码长度，digit-数字个数，symbol-符号个数
- "/-"：生成连字符密码，不支持任何参数，生成密码格式为 4字符-4字符-3字符
