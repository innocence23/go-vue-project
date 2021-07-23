package system

type ServiceGroup struct {
	JwtService
	ApiService
	AuthorityService
	AutoCodeService
	AutoCodeHistoryService
	BaseMenuService
	CasbinService
	DictionaryService
	DictionaryDetailService
	EmailService
	MenuService
	OperationRecordService
	SystemConfigService
	UserService
}
