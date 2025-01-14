package redis

import "time"

// redis键 数据相关
const (
	KeysUserIdNameMap    = "user_id_name_map"    // 用户名称map{id:name}
	KeysUserIdPhoneMap   = "user_id_phone_map"   // 用户名称map{id:phone}
	KeysItemIdMap        = "item_id_map"         // 检测项目信息map{id:item}
	KeysEventWorkerIdMap = "event_worker_id_map" // 事件任务信息map{id:event_worker}
	KeysBannerList       = "banner_list_"        // +parma banner列表
	KeysConfigInfo       = "config_info_"        // +name 配置信息
)

// redis键 操作相关
const (
	KeysCommonRepeatRequest = "common_repeat_request_" // +userId, 值为操作的路由；防止用户短时间内重复请求
	KeyLockerEventConsumer  = "locker_event_consumer"  // 锁，防止同时消费事件队列
)

// KeyExpireMap 键对应的过期时间
var KeyExpireMap = map[string]time.Duration{
	KeysCommonRepeatRequest: time.Second * 3,

	KeysUserIdNameMap:    time.Second * 3600,
	KeysUserIdPhoneMap:   time.Second * 3600,
	KeysItemIdMap:        time.Second * 3600,
	KeysEventWorkerIdMap: time.Second * 86400,
	KeysBannerList:       time.Second * 86400,
	KeysConfigInfo:       time.Second * 60,
}
