/**
 * @Time    :2022/2/25 9:02
 * @Author  :ZhangXiaoyu
 */

package casbin

import (
	"errors"
	"github.com/casbin/casbin/v2"
)

// AddRoleForUser
/**
 *  @Description: 为用户添加角色
 *  @param e
 *  @param user
 *  @param role
 *  @return err
 */
func AddRoleForUser(e *casbin.Enforcer, user, role string) (err error) {
	_, err = e.AddRoleForUser(user, role)
	if err != nil {
		return errors.New("为用户添加角色失败:" + err.Error())
	}
	return
}

// AddRolesForUser
/**
 *  @Description: 为用户添加多个角色
 *  @param e
 *  @param user
 *  @param roles
 *  @return err
 */
func AddRolesForUser(e *casbin.Enforcer, user string, roles []string) (err error) {
	_, err = e.AddRolesForUser(user, roles)
	if err != nil {
		return errors.New("为用户批量添加角色失败:" + err.Error())
	}
	return
}

// DeleteRoleForUser
/**
 *  @Description: 删除用户的一个角色
 *  @param e
 *  @param user
 *  @param role
 *  @return err
 */
func DeleteRoleForUser(e *casbin.Enforcer, user, role string) (err error) {
	_, err = e.DeleteRoleForUser(user, role)
	if err != nil {
		return errors.New("删除用户角色失败:" + err.Error())
	}
	return
}

// DeleteAllRolesForUser
/**
 *  @Description: 删除用户的所有角色
 *  @param e
 *  @param user
 *  @return err
 */
func DeleteAllRolesForUser(e *casbin.Enforcer, user string) (err error) {
	_, err = e.DeleteRolesForUser(user)
	if err != nil {
		return errors.New("删除用户所有角色失败:" + err.Error())
	}
	return
}

// GetRolesForUser
/**
 *  @Description: 查询用户所有角色
 *  @param e
 *  @param user
 *  @return roles
 *  @return err
 */
func GetRolesForUser(e *casbin.Enforcer, user string) (roles []string, err error) {
	roles, err = e.GetRolesForUser(user)
	if err != nil {
		return nil, errors.New("查询用户角色失败:" + err.Error())
	}
	return
}

// GetUsersForRole
/**
 *  @Description: 查询该角色下的用户
 *  @param e
 *  @param role
 *  @return users
 *  @return err
 */
func GetUsersForRole(e *casbin.Enforcer, role string) (users []string, err error) {
	users, err = e.GetUsersForRole(role)
	if err != nil {
		return nil, errors.New("查询角色的用户失败:" + err.Error())
	}
	return
}

// HasRoleForUser
/**
 *  @Description: 查询用户是否具有角色
 *  @param e
 *  @param user
 *  @param role
 *  @return has
 *  @return err
 */
func HasRoleForUser(e *casbin.Enforcer, user, role string) (has bool, err error) {
	has, err = e.HasRoleForUser(user, role)
	if err != nil {
		return false, errors.New("查询用户是否具有角色失败:" + err.Error())
	}
	return
}

// UpdateRolesForUser
/**
 *  @Description: 更新用户的角色列表
 *  @param e
 *  @param user
 *  @param roles
 *  @return err
 */
func UpdateRolesForUser(e *casbin.Enforcer, user string, roles []string) (err error) {
	// 删除用户角色
	err = DeleteAllRolesForUser(e, user)
	if err != nil {
		return err
	}
	// 如果用户角色不为空，则为用户添加角色
	if roles != nil {
		err = AddRolesForUser(e, user, roles)
		if err != nil {
			return err
		}
	}
	return
}
