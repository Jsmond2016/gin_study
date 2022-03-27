## 【GoWeb开发实战】Gin框架_数据库操作 CURD

- 课件：https://www.chaindesk.cn/witbook/19/352
- 安装/创建 mysql：[docker 创建 mysql](https://www.runoob.com/docker/docker-install-mysql.html)

步骤：

```bash
docker pull mysql:latest

docker images

docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql
```

使用mysql [可视化工具 heidisql](https://www.heidisql.com/download.php)，查看

mysql 连接的账号密码：

```
root/123456
```

## 建表

```sql
CREATE TABLE IF NOT EXISTS `user_info`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `user_name` VARCHAR(100) NOT NULL,
   `password` VARCHAR(40) NOT NULL,
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=UTF8;

USE test;
```



相关资料：

- https://www.liwenzhou.com/posts/Go/go_mysql/  Go语言操作MySQL——李文周的博客

- https://segmentfault.com/a/1190000021693989 golang 操作mysql示例（增、删、改、查、事务）
