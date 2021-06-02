package routers

import (
	adminControllers "beego-admin/controllers/admin"
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
		web.NSRouter("/editor/server", &adminControllers.EditorController{}, "get,post:Server"),

		//登录页
		web.NSRouter("/auth/login", &adminControllers.AuthController{}, "get:Login"),
		//退出登录
		web.NSRouter("/auth/logout", &adminControllers.AuthController{}, "get:Logout"),
		//二维码图片输出
		web.NSHandler("/auth/captcha/*.png", captcha.Server(240, 80)),
		//登录认证
		web.NSRouter("/auth/check_login", &adminControllers.AuthController{}, "post:CheckLogin"),
		//刷新验证码
		web.NSRouter("/auth/refresh_captcha", &adminControllers.AuthController{}, "post:RefreshCaptcha"),

		//首页
		web.NSRouter("/index/index", &adminControllers.IndexController{}, "get:Index"),

		//用户管理
		web.NSRouter("/admin_user/index", &adminControllers.AdminUserController{}, "get:Index"),

		//操作日志
		web.NSRouter("/admin_log/index", &adminControllers.AdminLogController{}, "get:Index"),

		//菜单管理
		web.NSRouter("/admin_menu/index", &adminControllers.AdminMenuController{}, "get:Index"),
		//菜单管理-添加菜单-界面
		web.NSRouter("/admin_menu/add", &adminControllers.AdminMenuController{}, "get:Add"),
		//菜单管理-添加菜单-创建
		web.NSRouter("/admin_menu/create", &adminControllers.AdminMenuController{}, "post:Create"),
		//菜单管理-修改菜单-界面
		web.NSRouter("/admin_menu/edit", &adminControllers.AdminMenuController{}, "get:Edit"),
		//菜单管理-更新菜单
		web.NSRouter("/admin_menu/update", &adminControllers.AdminMenuController{}, "post:Update"),
		//菜单管理-删除菜单
		web.NSRouter("/admin_menu/del", &adminControllers.AdminMenuController{}, "post:Del"),

		//系统管理-个人资料
		web.NSRouter("/admin_user/profile", &adminControllers.AdminUserController{}, "get:Profile"),
		//系统管理-个人资料-修改昵称
		web.NSRouter("/admin_user/update_nickname", &adminControllers.AdminUserController{}, "post:UpdateNickName"),
		//系统管理-个人资料-修改密码
		web.NSRouter("/admin_user/update_password", &adminControllers.AdminUserController{}, "post:UpdatePassword"),
		//系统管理-个人资料-修改头像
		web.NSRouter("/admin_user/update_avatar", &adminControllers.AdminUserController{}, "post:UpdateAvatar"),
		//系统管理-用户管理-添加界面
		web.NSRouter("/admin_user/add", &adminControllers.AdminUserController{}, "get:Add"),
		//系统管理-用户管理-添加
		web.NSRouter("/admin_user/create", &adminControllers.AdminUserController{}, "post:Create"),
		//系统管理-用户管理-修改界面
		web.NSRouter("/admin_user/edit", &adminControllers.AdminUserController{}, "get:Edit"),
		//系统管理-用户管理-修改
		web.NSRouter("/admin_user/update", &adminControllers.AdminUserController{}, "post:Update"),
		//系统管理-用户管理-启用
		web.NSRouter("/admin_user/enable", &adminControllers.AdminUserController{}, "post:Enable"),
		//系统管理-用户管理-禁用
		web.NSRouter("/admin_user/disable", &adminControllers.AdminUserController{}, "post:Disable"),
		//系统管理-用户管理-删除
		web.NSRouter("/admin_user/del", &adminControllers.AdminUserController{}, "post:Del"),

		//系统管理-角色管理
		web.NSRouter("/admin_role/index", &adminControllers.AdminRoleController{}, "get:Index"),
		//系统管理-角色管理-添加界面
		web.NSRouter("/admin_role/add", &adminControllers.AdminRoleController{}, "get:Add"),
		//系统管理-角色管理-添加
		web.NSRouter("/admin_role/create", &adminControllers.AdminRoleController{}, "post:Create"),
		//菜单管理-角色管理-修改界面
		web.NSRouter("/admin_role/edit", &adminControllers.AdminRoleController{}, "get:Edit"),
		//菜单管理-角色管理-修改
		web.NSRouter("/admin_role/update", &adminControllers.AdminRoleController{}, "post:Update"),
		//菜单管理-角色管理-删除
		web.NSRouter("/admin_role/del", &adminControllers.AdminRoleController{}, "post:Del"),
		//菜单管理-角色管理-启用角色
		web.NSRouter("/admin_role/enable", &adminControllers.AdminRoleController{}, "post:Enable"),
		//菜单管理-角色管理-禁用角色
		web.NSRouter("/admin_role/disable", &adminControllers.AdminRoleController{}, "post:Disable"),
		//菜单管理-角色管理-角色授权界面
		web.NSRouter("/admin_role/access", &adminControllers.AdminRoleController{}, "get:Access"),
		//菜单管理-角色管理-角色授权
		web.NSRouter("/admin_role/access_operate", &adminControllers.AdminRoleController{}, "post:AccessOperate"),

		//设置中心-后台设置
		web.NSRouter("/setting/admin", &adminControllers.SettingController{}, "get:Admin"),
		//设置中心-更新设置
		web.NSRouter("/setting/update", &adminControllers.SettingController{}, "post:Update"),

		//系统管理-开发管理-数据维护
		web.NSRouter("/database/table", &adminControllers.DatabaseController{}, "get:Table"),
		//系统管理-开发管理-数据维护-优化表
		web.NSRouter("/database/optimize", &adminControllers.DatabaseController{}, "post:Optimize"),
		//系统管理-开发管理-数据维护-修复表
		web.NSRouter("/database/repair", &adminControllers.DatabaseController{}, "post:Repair"),
		//系统管理-开发管理-数据维护-查看详情
		web.NSRouter("/database/view", &adminControllers.DatabaseController{}, "get,post:View"),

		//用户等级管理
		web.NSRouter("/user_level/index", &adminControllers.UserLevelController{}, "get:Index"),
		//用户等级管理-添加界面
		web.NSRouter("/user_level/add", &adminControllers.UserLevelController{}, "get:Add"),
		//用户等级管理-添加
		web.NSRouter("/user_level/create", &adminControllers.UserLevelController{}, "post:Create"),
		//用户等级管理-修改界面
		web.NSRouter("/user_level/edit", &adminControllers.UserLevelController{}, "get:Edit"),
		//用户等级管理-修改
		web.NSRouter("/user_level/update", &adminControllers.UserLevelController{}, "post:Update"),
		//用户等级管理-启用
		web.NSRouter("/user_level/enable", &adminControllers.UserLevelController{}, "post:Enable"),
		//用户等级管理-禁用
		web.NSRouter("/user_level/disable", &adminControllers.UserLevelController{}, "post:Disable"),
		//用户等级管理-删除
		web.NSRouter("/user_level/del", &adminControllers.UserLevelController{}, "post:Del"),
		//用户等级管理-导出
		web.NSRouter("/user_level/export", &adminControllers.UserLevelController{}, "get:Export"),

		//用户管理
		web.NSRouter("/user/index", &adminControllers.UserController{}, "get:Index"),
		//用户管理-添加界面
		web.NSRouter("/user/add", &adminControllers.UserController{}, "get:Add"),
		//用户管理-添加
		web.NSRouter("/user/create", &adminControllers.UserController{}, "post:Create"),
		//用户管理-修改界面
		web.NSRouter("/user/edit", &adminControllers.UserController{}, "get:Edit"),
		//用户管理-修改
		web.NSRouter("/user/update", &adminControllers.UserController{}, "post:Update"),
		//用户管理-启用
		web.NSRouter("/user/enable", &adminControllers.UserController{}, "post:Enable"),
		//用户管理-禁用
		web.NSRouter("/user/disable", &adminControllers.UserController{}, "post:Disable"),
		//用户管理-删除
		web.NSRouter("/user/del", &adminControllers.UserController{}, "post:Del"),
		//用户管理-导出
		web.NSRouter("/user/export", &adminControllers.UserController{}, "get:Export"),
	)

	web.AddNamespace(admin)
}
