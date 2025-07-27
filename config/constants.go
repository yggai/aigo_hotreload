package config

// 服务器相关常量
const (
	DefaultPort = "8888"
	DefaultHost = "localhost"
)

// URL相关常量
var (
	BaseURL     = "http://" + DefaultHost + ":" + DefaultPort
	HealthURL   = BaseURL + "/health"
	APIBaseURL  = BaseURL + "/api/v1"
	UsersAPIURL = APIBaseURL + "/users"
)

// Air相关常量
const (
	AirInstallCmd = "go install github.com/air-verse/air@latest"
	AirCommand    = "air"
)

// 文件权限常量
const (
	DirPermission  = 0755
	FilePermission = 0644
)

// 状态消息常量
const (
	StatusSuccess = "success"
	StatusHealthy = "healthy"
)

// 时间格式常量
const (
	TimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
) 