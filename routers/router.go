package routers

import (
	controllersAdmin "beego-admin/controllers/admin"
	"beego-admin/middleware"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dchest/captcha"
	"net/http"
)

func init() {
	//授权登录中间件
	middleware.AuthMiddle()

	web.Get("/", func(ctx *context.Context) {
		ctx.Redirect(http.StatusFound, "/admin/index/index")
	})

	//admin模块路由
	admin := web.NewNamespace("/admin",
		//UEditor控制器
		web.NSRouter("/editor/server", &controllersAdmin.EditorController{}, "get,post:Server"),

		//登录页
		web.NSRouter("/auth/login", &controllersAdmin.AuthController{}, "get:Login"),
		//退出登录
		web.NSRouter("/auth/logout", &controllersAdmin.AuthController{}, "get:Logout"),
		//二维码图片输出
		web.NSHandler("/auth/captcha/*.png", captcha.Server(240, 80)),
		//登录认证
		web.NSRouter("/auth/check_login", &controllersAdmin.AuthController{}, "post:CheckLogin"),
		//刷新验证码
		web.NSRouter("/auth/refresh_captcha", &controllersAdmin.AuthController{}, "post:RefreshCaptcha"),

		//首页
		web.NSRouter("/index/index", &controllersAdmin.IndexController{}, "get:Index"),

		//用户管理
		web.NSRouter("/admin_user/index", &controllersAdmin.AdminUserController{}, "get:Index"),

		//操作日志
		web.NSRouter("/admin_log/index", &controllersAdmin.AdminLogController{}, "get:Index"),

		//菜单管理
		web.NSRouter("/admin_menu/index", &controllersAdmin.AdminMenuController{}, "get:Index"),
		//菜单管理-添加菜单-界面
		web.NSRouter("/admin_menu/add", &controllersAdmin.AdminMenuController{}, "get:Add"),
		//菜单管理-添加菜单-创建
		web.NSRouter("/admin_menu/create", &controllersAdmin.AdminMenuController{}, "post:Create"),
		//菜单管理-修改菜单-界面
		web.NSRouter("/admin_menu/edit", &controllersAdmin.AdminMenuController{}, "get:Edit"),
		//菜单管理-更新菜单
		web.NSRouter("/admin_menu/update", &controllersAdmin.AdminMenuController{}, "post:Update"),
		//菜单管理-删除菜单
		web.NSRouter("/admin_menu/del", &controllersAdmin.AdminMenuController{}, "post:Del"),

		//系统管理-个人资料
		web.NSRouter("/admin_user/profile", &controllersAdmin.AdminUserController{}, "get:Profile"),
		//系统管理-个人资料-修改昵称
		web.NSRouter("/admin_user/update_nickname", &controllersAdmin.AdminUserController{}, "post:UpdateNickName"),
		//系统管理-个人资料-修改密码
		web.NSRouter("/admin_user/update_password", &controllersAdmin.AdminUserController{}, "post:UpdatePassword"),
		//系统管理-个人资料-修改头像
		web.NSRouter("/admin_user/update_avatar", &controllersAdmin.AdminUserController{}, "post:UpdateAvatar"),
		//系统管理-用户管理-添加界面
		web.NSRouter("/admin_user/add", &controllersAdmin.AdminUserController{}, "get:Add"),
		//系统管理-用户管理-添加
		web.NSRouter("/admin_user/create", &controllersAdmin.AdminUserController{}, "post:Create"),
		//系统管理-用户管理-修改界面
		web.NSRouter("/admin_user/edit", &controllersAdmin.AdminUserController{}, "get:Edit"),
		//系统管理-用户管理-修改
		web.NSRouter("/admin_user/update", &controllersAdmin.AdminUserController{}, "post:Update"),
		//系统管理-用户管理-启用
		web.NSRouter("/admin_user/enable", &controllersAdmin.AdminUserController{}, "post:Enable"),
		//系统管理-用户管理-禁用
		web.NSRouter("/admin_user/disable", &controllersAdmin.AdminUserController{}, "post:Disable"),
		//系统管理-用户管理-删除
		web.NSRouter("/admin_user/del", &controllersAdmin.AdminUserController{}, "post:Del"),

		//系统管理-角色管理
		web.NSRouter("/admin_role/index", &controllersAdmin.AdminRoleController{}, "get:Index"),
		//系统管理-角色管理-添加界面
		web.NSRouter("/admin_role/add", &controllersAdmin.AdminRoleController{}, "get:Add"),
		//系统管理-角色管理-添加
		web.NSRouter("/admin_role/create", &controllersAdmin.AdminRoleController{}, "post:Create"),
		//菜单管理-角色管理-修改界面
		web.NSRouter("/admin_role/edit", &controllersAdmin.AdminRoleController{}, "get:Edit"),
		//菜单管理-角色管理-修改
		web.NSRouter("/admin_role/update", &controllersAdmin.AdminRoleController{}, "post:Update"),
		//菜单管理-角色管理-删除
		web.NSRouter("/admin_role/del", &controllersAdmin.AdminRoleController{}, "post:Del"),
		//菜单管理-角色管理-启用角色
		web.NSRouter("/admin_role/enable", &controllersAdmin.AdminRoleController{}, "post:Enable"),
		//菜单管理-角色管理-禁用角色
		web.NSRouter("/admin_role/disable", &controllersAdmin.AdminRoleController{}, "post:Disable"),
		//菜单管理-角色管理-角色授权界面
		web.NSRouter("/admin_role/access", &controllersAdmin.AdminRoleController{}, "get:Access"),
		//菜单管理-角色管理-角色授权
		web.NSRouter("/admin_role/access_operate", &controllersAdmin.AdminRoleController{}, "post:AccessOperate"),

		//设置中心-后台设置
		web.NSRouter("/setting/admin", &controllersAdmin.SettingController{}, "get:Admin"),
		//设置中心-更新设置
		web.NSRouter("/setting/update", &controllersAdmin.SettingController{}, "post:Update"),

		//系统管理-开发管理-数据维护
		web.NSRouter("/database/table", &controllersAdmin.DatabaseController{}, "get:Table"),
		//系统管理-开发管理-数据维护-优化表
		web.NSRouter("/database/optimize", &controllersAdmin.DatabaseController{}, "post:Optimize"),
		//系统管理-开发管理-数据维护-修复表
		web.NSRouter("/database/repair", &controllersAdmin.DatabaseController{}, "post:Repair"),
		//系统管理-开发管理-数据维护-查看详情
		web.NSRouter("/database/view", &controllersAdmin.DatabaseController{}, "get,post:View"),

		//用户等级管理
		web.NSRouter("/user_level/index", &controllersAdmin.UserLevelController{}, "get:Index"),
		//用户等级管理-添加界面
		web.NSRouter("/user_level/add", &controllersAdmin.UserLevelController{}, "get:Add"),
		//用户等级管理-添加
		web.NSRouter("/user_level/create", &controllersAdmin.UserLevelController{}, "post:Create"),
		//用户等级管理-修改界面
		web.NSRouter("/user_level/edit", &controllersAdmin.UserLevelController{}, "get:Edit"),
		//用户等级管理-修改
		web.NSRouter("/user_level/update", &controllersAdmin.UserLevelController{}, "post:Update"),
		//用户等级管理-启用
		web.NSRouter("/user_level/enable", &controllersAdmin.UserLevelController{}, "post:Enable"),
		//用户等级管理-禁用
		web.NSRouter("/user_level/disable", &controllersAdmin.UserLevelController{}, "post:Disable"),
		//用户等级管理-删除
		web.NSRouter("/user_level/del", &controllersAdmin.UserLevelController{}, "post:Del"),
		//用户等级管理-导出
		web.NSRouter("/user_level/export", &controllersAdmin.UserLevelController{}, "get:Export"),

		//用户管理
		web.NSRouter("/user/index", &controllersAdmin.UserController{}, "get:Index"),
		//用户管理-添加界面
		web.NSRouter("/user/add", &controllersAdmin.UserController{}, "get:Add"),
		//用户管理-添加
		web.NSRouter("/user/create", &controllersAdmin.UserController{}, "post:Create"),
		//用户管理-修改界面
		web.NSRouter("/user/edit", &controllersAdmin.UserController{}, "get:Edit"),
		//用户管理-修改
		web.NSRouter("/user/update", &controllersAdmin.UserController{}, "post:Update"),
		//用户管理-启用
		web.NSRouter("/user/enable", &controllersAdmin.UserController{}, "post:Enable"),
		//用户管理-禁用
		web.NSRouter("/user/disable", &controllersAdmin.UserController{}, "post:Disable"),
		//用户管理-删除
		web.NSRouter("/user/del", &controllersAdmin.UserController{}, "post:Del"),
		//用户管理-导出
		web.NSRouter("/user/export", &controllersAdmin.UserController{}, "get:Export"),
	)

	web.AddNamespace(admin)
}
