package generator

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/jinzhu/inflection"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

func Generate(c *cli.Context) error{
	dbSns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.String("username"),
		c.String("password"),
		c.String("host"),
		c.String("port"),
		c.String("database"),
	)

	db := GetDB(dbSns)

	if c.String("t") == "ALL"{
		tables := db.GetDataBySql("show tables")
		for _, table := range tables{
			tableName := table["Tables_in_"  + c.String("d")]
			columns := db.GetDataBySql("show full columns from " + tableName)
			generateModel(tableName, columns, c.String("dir"))
		}
	} else {
		columns := db.GetDataBySql("show full columns from " + c.String("t"))
		generateModel(c.String("t"), columns, c.String("dir"))
	}

	return nil
}

func generateModel(tableName string, columns []map[string]string, dir string) {
	var codes []jen.Code
	for _, col := range columns{
		t := col["Type"]
		column := col["Field"]
		var st *jen.Statement
		if column == "id" {
			st = jen.Id("ID").Uint().Tag(map[string]string{"json": "id","gorm": "column:"+column,"db": column,})
		} else {
			st = jen.Id(SnakeCase2CamelCase(column, true))
			getCol(st, t)
			st.Tag(map[string]string{
				"json": column,
				"gorm": "column:"+column,
				"db": column,
			})

			st.Comment(col["Comment"])
		}
		codes = append(codes, st)
	}
	f := jen.NewFilePath(dir)
	//定义结构体
	f.Type().Id(SnakeCase2CamelCase(inflection.Singular(tableName), true)).Struct(codes...)

	//获取表名
	f.Func().Params(jen.Id(SnakeCase2CamelCase(inflection.Singular(tableName), true))).Id("TableName").Params().String().Block(
		jen.Return().Lit(tableName),
		)

	_ = os.MkdirAll(dir, os.ModePerm)
	fileName := dir + "/" + inflection.Singular(tableName) + ".go"
	err := f.Save(fileName)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(fileName)
}

func getCol(st *jen.Statement, t string) {
	prefix := strings.Split(t, "(")[0]
	switch prefix {
	case "int", "tinyint", "smallint", "bigint", "mediumint":
		st.Int()
	case "float":
		st.Float32()
	case "varchar":
		st.String()
	case "decimal":
		st.Float32()
	case "date", "time", "timestamp", "year", "datetime":
		st.Qual("time", "Time")
	default:
		st.String()
	}
}

func SnakeCase2CamelCase(input string, pascal bool) string {
	names := strings.Split(input, "_")
	var n string
	for k, name := range names {
		if name == "id" {
			n += "ID"
		} else {
			if k == 0 && !pascal {
				n += name
			} else {
				n += strings.Title(name)
			}
		}
	}
	return n
}
