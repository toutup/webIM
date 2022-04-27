package online

import (
	"fmt"
	"xiliangzi_pro/pkg/redisConn"

	"github.com/gomodule/redigo/redis"
)

type Online interface {
	AddOnlineList(req AddOnlineListRequest) error
	DelOnlineList(req DelOnlinListRequest) error
	GetOnlineList(req GetOnlineListRequest) ([]GetOnlineListResponse, error)
}

func NewOnline() Online {
	return &onlineList{}
}

type onlineList struct{}

/**
 *  @Name AddOnlineList
 *  @Description: 添加用户到onliest
 *  @receiver o *onlineList
 *  @param req AddOnlineListRequest
 *  @return error
 */
func (o *onlineList) AddOnlineList(req AddOnlineListRequest) error {
	// hasName
	hasName := fmt.Sprintf("%s:%v", req.Prefix, req.Uid)
	// 获取redis连接
	conn, _ := redisConn.GetOnlineListConn()
	defer conn.Close()
	// 将用户信息存储在hash表中
	_, err := conn.Do("hmset", redis.Args{}.Add(hasName).AddFlat(req.Data)...)
	if err != nil {
		return err
	}
	// 将元素添加到有序集合
	zSetName := fmt.Sprintf("%s:%s", req.Prefix, req.SetName)
	_, err = conn.Do("ZADD", zSetName, req.DateTime, req.Uid)
	if err != nil {
		return err
	}
	return nil
}

/**
 *  @Name DelOnlineList
 *  @Description: 将用户从onlineList删除
 *  @receiver o *onlineList
 *  @param req DelOnlinListRequest
 *  @return error
 */
func (o *onlineList) DelOnlineList(req DelOnlinListRequest) error {
	// 获取redis连接
	conn, _ := redisConn.GetOnlineListConn()
	defer conn.Close()
	// 将用户从在线列表移除
	zSetName := fmt.Sprintf("%s:%s", req.Prefix, req.SetName)
	_, err := conn.Do("ZREM", zSetName, req.Uid)
	if err != nil {
		return err
	}
	// 将用户信息从hash表中删除
	hasName := fmt.Sprintf("%s:%v", req.Prefix, req.Uid)
	_, err = conn.Do("DEL", hasName)
	if err != nil {
		return err
	}
	return nil
}

/**
 *  @Name GetOnlineList
 *  @Description:
 *  @receiver *onlineList
 *  @param req GetOnlineListRequest
 *  @return []GetOnlineListResponse
 *  @return error
 */
func (o *onlineList) GetOnlineList(req GetOnlineListRequest) ([]GetOnlineListResponse, error) {
	// 获取redis连接
	conn, _ := redisConn.GetOnlineListConn()
	defer conn.Close()
	// 计算limitStart， limitEnd
	limitStart := (req.Page - 1) * req.Size
	limitEnd := (limitStart + req.Size) - 1
	// 利用redis有序集合zrange方法获取指定区间的值，返回带有score值的有序结果集
	zSetName := fmt.Sprintf("%s:%s", req.Prefix, req.SetName)
	res, err := redis.ByteSlices(conn.Do("ZRANGE", zSetName, limitStart, limitEnd, "withscores"))
	if err != nil {
		return []GetOnlineListResponse{}, err
	}
	// 获取用户消息
	var sUser = make([]GetOnlineListResponse, 0, req.Size)
	for _, uid := range res {
		hasName := fmt.Sprintf("%s:%v", req.Prefix, uid)
		userInfo, _ := redis.Values(conn.Do("hgetall", hasName))
		if len(userInfo) == 0 {
			continue
		}
		var user GetOnlineListResponse
		redis.ScanStruct(userInfo, &user)
		sUser = append(sUser, user)
	}

	return sUser, nil
}
