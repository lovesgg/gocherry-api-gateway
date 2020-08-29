package common_enum

import "time"

const (
	REDIS_EXPIRE_TIME_BY_FIVE_SECOND  = 5 * time.Second       //5秒
	REDIS_EXPIRE_TIME_BY_ONE_MINUTE   = 60 * time.Second      //1分钟
	REDIS_EXPIRE_TIME_BY_THREE_MINUTE = 180 * time.Second     //三分钟
	REDIS_EXPIRE_TIME_BY_FIVE_MINUTE  = 300 * time.Second     //五分钟
	REDIS_EXPIRE_TIME_BY_SIX_MINUTE   = 360 * time.Second     //六分钟
	REDIS_EXPIRE_TIME_BY_TEN_MINUTE   = 600 * time.Second     //十分钟
	REDIS_EXPIRE_TIME_BY_ONE_HOUR     = 3600 * time.Second    //1小时
	REDIS_EXPIRE_TIME_BY_TWO_HOUR     = 7200 * time.Second    //两小时
	REDIS_EXPIRE_TIME_BY_FOUR_HOUR    = 14400 * time.Second   //四小时
	REDIS_EXPIRE_TIME_BY_EIGHT_HOUR   = 28800 * time.Second   //八小时
	REDIS_EXPIRE_TIME_BY_TWELVE_HOURS = 43200 * time.Second   //12
	REDIS_EXPIRE_TIME_BY_ONE_DAY      = 86400 * time.Second   //一天
	REDIS_EXPIRE_TIME_BY_FIFTEEN_DAY  = 1296000 * time.Second //十五天
	REDIS_EXPIRE_TIME_BY_ONE_MONTH    = 2592000 * time.Second //一个月
	REDIS_EXPIRE_TIME_BY_SIXTEEN_HOUR = 57600 * time.Second   //16小时
)
