package user

import (
	"github.com/jettjia/igo-pkg/pkg/util"
	"gorm.io/gorm"
)

type SysUser struct {
	Ulid       string `gorm:"column:ulid;primaryKey;type:varchar(128);comment:ulid;" json:"ulid"`                  // ulid
	CreatedAt  int64  `gorm:"column:created_at;autoCreateTime:milli;type:bigint;comment:创建时间;" json:"created_at"`  // 创建时间
	UpdatedAt  int64  `gorm:"column:updated_at;autoUpdateTime:milli;type:bigint;comment:修改时间;" json:"updated_at"`  // 修改时间
	DeletedAt  int64  `gorm:"column:deleted_at;autoDeletedTime:milli;type:bigint;comment:删除时间;" json:"deleted_at"` // 删除时间
	CreatedBy  string `gorm:"column:created_by;type:varchar(128);comment:创建者;" json:"created_by"`                  // 创建者
	UpdatedBy  string `gorm:"column:updated_by;type:varchar(128);comment:修改者;" json:"updated_by"`                  // 修改者
	DeletedBy  string `gorm:"column:deleted_by;type:varchar(128);comment:删除者;" json:"deleted_by"`                  // 删除者
	MemberCode string `gorm:"column:member_code;type:varchar(32);comment:会员号;" json:"member_code"`                 // 会员号
	Phone      string `gorm:"column:phone;type:varchar(20);comment:手机号码;" json:"phone"`                            // 手机号码
	LevelId    string `gorm:"column:level_id;type:varchar(128);default: 0;comment:会员等级id;" json:"level_id"`        // 会员等级id
	NickName   string `gorm:"column:nick_name;type:varchar(64);comment:昵称;" json:"nick_name"`                      // 昵称，不可更改，判断唯一
	TrueName   string `gorm:"column:true_name;type:varchar(64);comment:真实姓名;" json:"true_name"`                    // 真实姓名，可重复
	State      uint   `gorm:"column:state;type:int;default: 1;comment:1显示,2否;" json:"state"`                       // 1显示,2否
	Email      string `gorm:"column:email;type:varchar(255);comment:邮箱;" json:"email"`                             // 邮箱
	Password   string `gorm:"column:password;type:varchar(2000);comment:密码;" json:"password"`                      // 密码
	AdminLevel uint   `gorm:"column:admin_level;type:int;default: 0;comment:超管级别;" json:"admin_level"`             // 1是admin超管
	DepId      string `gorm:"column:dep_id;type:varchar(128);comment:部门id;" json:"dep_id"`                         // 部门id
	JobId      string `gorm:"column:job_id;type:varchar(128);comment:职位id;" json:"job_id"`                         // 职位id
	RoleId     string `gorm:"column:role_id;type:varchar(128);comment:角色id;" json:"role_id"`                       // 角色id
}

func (po *SysUser) BeforeCreate(tx *gorm.DB) (err error) {
	po.Ulid = util.Ulid()
	return
}

func (po *SysUser) TableName() string {
	return "sys_user"
}
