package admin

import (
	"beego-admin/services"
)

// AdminLogController struct.
type AdminLogController struct {
	baseController
}

// Index index.
func (alc *AdminLogController) Index() {
	var (
		adminLogService  services.AdminLogService
		adminUserService services.AdminUserService
	)
	data, pagination := adminLogService.GetPaginateData(admin["per_page"].(int), gQueryParams)

	alc.Data["admin_user_list"] = adminUserService.GetAllAdminUser()

	alc.Data["data"] = data
	alc.Data["paginate"] = pagination
	alc.Layout = "admin/public/base.html"
	alc.TplName = "admin/admin_log/index.html"
}
