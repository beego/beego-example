module github.com/beego-dev/beego-example

go 1.14

require (
	github.com/beego/beego/v2 v2.0.2-0.20210415083014-29fd8719a8ab
	github.com/go-sql-driver/mysql v1.6.0
	github.com/lib/pq v1.10.2
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/opentracing-contrib/beego v0.0.0-20190807091843-a0eec2f7e67c
)

replace github.com/beego/beego/v2 => ../beego
