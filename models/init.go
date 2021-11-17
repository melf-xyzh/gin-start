/**
 * @Time    :2021/11/9 15:17
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:init.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package models

import (
	"errors"
	conn "gin-start/connection"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

// DBAutoMigrate
/**
 * @Description: 自动迁移
 * @param DB 数据库连接池
 */
func DBAutoMigrate() (err error) {
	// 自动迁移
	err = conn.DB.AutoMigrate(
		&User{}, // 用户
		&Role{}, // 角色
	)
	if err != nil {
		return errors.New("自动迁移失败：" + err.Error())
	}
	return nil
}

// InitModel
/**
 * @Description: 初始化数据表
 */
func InitModel() {
	// Gorm适配器
	adapter, err1 := gormadapter.NewAdapterByDBWithCustomTable(conn.DB,&CasbinRule{})
	if err1 != nil {
		log.Fatalln("Casbin Gorm适配器错误：" + err1.Error())
	}
	log.Println("导入适配器")

	//// Gorm适配器
	//adapter, err1 := gormadapter.NewAdapterByDB(conn.DB)
	//if err1 != nil {
	//	log.Fatalln("Casbin Gorm适配器错误：" + err1.Error())
	//}
	// 通过ORM新建一个执行者
	Enforcer, err2 := casbin.NewEnforcer("config\\rbac_model.conf", adapter)
	if err2 != nil {
		log.Fatalln("新建Casbin执行者异常：" + err2.Error())
	}
	// 导入访问策略
	err3 := Enforcer.LoadPolicy()
	if err3 != nil {
		log.Fatalln("导入访问策略异常：" + err3.Error())
	}

	//subject := "tom"
	//object := "/api/routers"
	//action := "POST"
	//
	//// 添加策略
	//_, err := Enforcer.AddPolicy(subject, object, action)
	//if err != nil {
	//	log.Fatalln(err.Error())
	//	return
	//}
	//log.Println("添加策略成功")
	//
	//// 为用户添加角色
	//_, err = Enforcer.AddRoleForUser("alice", "role1")
	//if err != nil {
	//	return
	//}
	//
	//// 为用户或角色添加权限
	//Enforcer.AddPermissionForUser("role2", "read1")

}
