module tiktok/dal

replace tiktok/pkg => ../pkg

go 1.19

require (
	gorm.io/driver/mysql v1.5.1
	gorm.io/gen v0.3.23
	gorm.io/gorm v1.25.4
	gorm.io/plugin/dbresolver v1.4.7
	tiktok/pkg v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	gorm.io/datatypes v1.1.1-0.20230130040222-c43177d3cf8c // indirect
	gorm.io/hints v1.1.0 // indirect
)
