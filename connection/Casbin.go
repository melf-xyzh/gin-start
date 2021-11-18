/**
 * @Time    :2021/11/17 17:42
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:Casbin.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package connection

import (
	"errors"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

var (
	Enforcer *casbin.Enforcer
)

func InitCasbin() (err error) {
	// Gorm适配器
	adapter, err := gormadapter.NewAdapterByDB(DB)
	if err != nil {
		return errors.New("Casbin Gorm适配器错误：" + err.Error())
	}
	log.Println("导入适配器")

	// 通过ORM新建一个执行者
	Enforcer, err = casbin.NewEnforcer("config\\rbac_model.conf", adapter)
	if err != nil {
		return errors.New("新建Casbin执行者异常：" + err.Error())
	}
	// 导入访问策略
	err = Enforcer.LoadPolicy()
	if err != nil {
		return errors.New("导入访问策略异常：" + err.Error())
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
	return err
}