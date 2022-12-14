package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/config"
	"time"
)

func Interdict(ctx *gin.Context, conf config.MiddlewareConfig) {
	ctx.Set(conf.InterdictorTraceKey, true)
}

func NewInterdictor(conf config.MiddlewareConfig) func(ctx *gin.Context) {
	limiter := kv.NewSyncMap[kv.Pair[uint, time.Time]]()
	return func(ctx *gin.Context) {
		success, ip := GetIPTrace(ctx, conf)
		if !success {
			// 无法成功获取客户端IP，返回BadRequest
			ctx.AbortWithStatus(400)
			return
		}

		if limiter.Exists(ip) {
			if pair := limiter.Get(ip); pair.Key() > conf.InterdictorTriggerCount {
				if blockedTime := pair.Value(); !blockedTime.IsZero() && blockedTime.Add(time.Second*time.Duration(conf.InterdictorBlockedTime)).After(time.Now()) {
					// 失败次数过多且未到解封时间，返回PreconditionFailed
					ctx.AbortWithStatus(412)
					return
				} else {
					// 已到解封时间，重置封禁状态
					limiter.Remove(ip)
				}
			}
		}

		ctx.Next()

		if f, exist := ctx.Get(conf.InterdictorTraceKey); exist {
			if result, ok := f.(bool); ok && result {
				if limiter.Exists(ip) {
					// 存在标记且被封禁器记录了
					pair := limiter.Get(ip)
					pair.SetKey(pair.Key() + 1)

					if pair.Key() >= conf.InterdictorTriggerCount {
						// 达到封禁标准，封禁
						pair.SetValue(time.Now())
					}

					limiter.Set(ip, pair)
				} else {
					// 存在标记但没被封禁器记录，增加记录
					limiter.Set(ip, kv.NewPair(uint(1), time.Time{}))
				}
			} else if limiter.Exists(ip) {
				// 标记为假但以前被封禁器记录过，重置状态
				limiter.Remove(ip)
			}
		} else if limiter.Exists(ip) {
			// 没有标记但以前被封禁器记录过，重置状态
			limiter.Remove(ip)
		}
	}
}
