package golib

import (
	"fmt"
	"github.com/tubzby/beego/orm"
)

type MysqlInterface struct {
	query orm.QuerySeter
	inErr error
	sp    string
	dStr  string
}

type DBParams map[string]interface{}
type DBParamsList []interface{}

//var ormer orm.ormer
var DB MysqlInterface

//define global database name
var GCommentMeta = new(TBCommentMeta).TableName()
var GComments = new(TBComments).TableName()
var GLinks = new(TBLinks).TableName()
var GOptions = new(TBOptions).TableName()
var GPostMeta = new(TBPostMeta).TableName()
var GPost = new(TBPost).TableName()
var GTermMeta = new(TBTermMeta).TableName()
var GTermRelationShips = new(TBTermRelationShips).TableName()
var GTermTaxonomy = new(TBTermTaxonomy).TableName()
var GUserMeta = new(TBUserMeta).TableName()
var GTBUser = new(TBUser).TableName()

func (this *MysqlInterface) init(mysql MysqlConf) {
	//register model
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", mysql.User, mysql.Password, mysql.IP, mysql.Port, mysql.DB)
	orm.RegisterModel(getAllTable()...)
	orm.RegisterDataBase("default", "mysql", connStr)
	orm.SetMaxIdleConns("default", mysql.MaxIdle)
	orm.SetMaxOpenConns("default", mysql.MaxActive)

	logger.Debug("init mysql conns maxIdle:%v , maxActive:%v\n", mysql.MaxIdle, mysql.MaxActive)
}
