package online

/**
 * AddOnlineListRequest
 * @Description: 添加用户到在线列表请求数据
 */
type AddOnlineListRequest struct {
	Prefix   string            `json:"hash_prefix"`
	SetName  string            `json:"set_name"`
	Uid      int               `json:"uid"`
	DateTime int64             `json:"date_time"`
	Data     map[string]string `json:"data"`
}

/**
 * AddOnlineListResponse
 * @Description: 添加用户到在线列表返回数据
 */
type AddOnlineListResponse struct {
	Code    int               `json:"code" :"code"`
	Message string            `json:"message" :"message"`
	Data    map[string]string `json:"data" :"data"`
}

/**
 * @Name DelOnlinListRequest
 * @Description: 删除用户请求数据
 */
type DelOnlinListRequest struct {
	Prefix  string `json:"hash_prefix"`
	SetName string `json:"set_name"`
	Uid     int    `json:"uid"`
}

/**
 * @Name GetOnlineListRequest
 * @Description: 获取onlineList请求
 */
type GetOnlineListRequest struct {
	Prefix  string `json:"hash_prefix"`
	SetName string `json:"set_name"`
	Page    int    `json:"page"`
	Size    int    `json:"size"`
}

type GetOnlineListResponse struct {
	Uid      string `json:"uid"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	NickName string `json:"nick_name"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}
