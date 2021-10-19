# go-project
基于Gin+Crontab的封装框架Demo
### Required

- Mysql
- Redis
### Ready

Create a **Test database** and import SQL
```
CREATE TABLE `test` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '名称',
  `ctime` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `mtime` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
### Run
```
$ cd $GOPATH/src/go-project

$ go run main.go 
```
## Features

- [x] Gin
- [x] Crontab
- [x] RESTful API
- [x] Gorm
- [x] Redis
- [x] logging
- [x] App configurable
- [ ] Middleware
- [ ] Swagger
- [ ] Jwt
- [ ] Docker