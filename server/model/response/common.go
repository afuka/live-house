package response

// Response 基础序列化器
type Response struct {
	Code int         `json:"error_code"`
	Msg  string      `json:"error_msg"`
	Data interface{} `json:"data"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

const (
	// Success 成功
	Success int = 0
	// AuthErr 没有权限
	AuthErr int = 10000
	// ServiceErr 服务内部错误
	ServiceErr int = 20000
	// ParamsErr 参数错误
	ParamsErr int = 30000
	// OthersErr 其他错误
	OthersErr int = 40000
	// PermissionErr 权限错误
	PermissionErr int = 50000
)
