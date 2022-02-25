/**
 * @Time    :2022/2/25 9:29
 * @Author  :ZhangXiaoyu
 */

package casbin

import (
	"errors"
	"github.com/casbin/casbin/v2"
)

type Api struct {
	Url    string
	Method string
}

// AddApiForRole
/**
 *  @Description: 为角色添加Api权限
 *  @param e
 *  @param role
 *  @param api
 *  @return err
 */
func AddApiForRole(e *casbin.Enforcer, role string, api Api) (err error) {
	_, err = e.AddPolicy(role, api.Url, api.Method)
	if err != nil {
		return errors.New("为角色添加API权限失败:" + err.Error())
	}
	return
}

// AddApisForRole
/**
 *  @Description: 为角色添加多个Api权限
 *  @param e
 *  @param role
 *  @param apis
 *  @return err
 */
func AddApisForRole(e *casbin.Enforcer, role string, apis []Api) (err error) {
	if apis == nil {
		return
	}
	rules := make([][]string, 0)
	for _, api := range apis {
		rules = append(rules, []string{role, api.Url, api.Method})
	}
	_, err = e.AddPolicies(rules)
	if err != nil {
		return errors.New("为角色批量添加API权限失败:" + err.Error())
	}
	return
}

// DeleteApisForRole
/**
 *  @Description: 删除角色的API权限
 *  @param e
 *  @param role
 *  @return err
 */
func DeleteApisForRole(e *casbin.Enforcer, role string) (err error) {
	_, err = e.DeletePermissionsForUser(role)
	if err != nil {
		return errors.New("删除角色API权限失败")
	}
	return
}

// FindApisForRole
/**
 *  @Description: 查询角色Api列表
 *  @param e
 *  @param role
 *  @return apis
 */
func FindApisForRole(e *casbin.Enforcer, role string) (apis []Api) {
	apis = make([]Api, 0)
	result := e.GetPermissionsForUser(role)
	for _, str := range result {
		var api Api
		api.Url = str[1]
		api.Method = str[2]
		apis = append(apis, api)
	}
	return
}

// FindApisForUser
/**
 *  @Description: 查询用户的Api权限列表
 *  @param e
 *  @param user
 *  @return apis
 *  @return err
 */
func FindApisForUser(e *casbin.Enforcer, user string) (apis []Api, err error) {
	apis = make([]Api, 0)
	result, err := e.GetImplicitPermissionsForUser(user)
	for _, str := range result {
		var api Api
		api.Url = str[1]
		api.Method = str[2]
		apis = append(apis, api)
	}
	return
}

// Check
/**
 *  @Description: 权限校验
 *  @param e
 *  @param userOrRole
 *  @param url
 *  @param method
 *  @return ok
 *  @return err
 */
func Check(e *casbin.Enforcer, userOrRole, url, method string) (ok bool, err error) {
	ok, err = e.Enforce(userOrRole, url, method)
	if ok {
		return true, nil
	} else {
		if err != nil {
			return false, errors.New("权限校验异常:" + err.Error())
		} else {
			return false, nil
		}
	}
}
