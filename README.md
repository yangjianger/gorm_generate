# gorm model生成工具
    go install github.com/yangjianger/gorm_generate

### 使用
    $GOPATH/bin/gorm_gen -host=host -port=port -u=username -p=password -d=database -t=ALL -dir=./model
    Options:
        -host string 链接地址
        -port string 端口
        -u string 用户名
        -p string  密码
        -d string 数据库
        -t string 表名  (default "ALL")
        -dir string 保存路径 (default "./models")

### 生成效果
    package model

    type Category struct {
        ID            uint   `db:"id" gorm:"column:id" json:"id"`
        CategoryID    int    `db:"category_id" gorm:"column:category_id" json:"category_id"`             // 分类id
        Category      string `db:"category" gorm:"column:category" json:"category"`                      // 分类名称
        Status        int    `db:"status" gorm:"column:status" json:"status"`                            // 状态 1正常 -1删除
    }
    
    func (Category) TableName() string {
        return "category"
    }


