package model

import "time"

type FuseConfig struct {
	Id                    uint64    `json:"id" gorm:"primary_key"`
	Name                  string    `json:"name"`
	ForceBreakerSwitch    uint8     `json:"force_breaker_switch"`
	BreakerStrategy       uint8     `json:"breaker_strategy"`
	DegradeStrategy       uint8     `json:"degrade_strategy"`
	DegradeStrategyDetail string    `json:"degrade_strategy_detail"`
	BreakerStrategyDetail string    `json:"breaker_strategy_detail"`
	IsDelete              uint8     `json:"is_delete"`
	CreateTime            time.Time `json:"create_time"`
	UpdateTime            time.Time `json:"update_time"`
}

func (FuseConfig) TableName() string {
	return "fuse_config"
}

type BreakerStrategy struct {
	FailCount            uint8 `json:"fail_count"`
	ConsecutiveFailCount uint8 `json:"consecutive_fail_count"`
	FailRate             uint8 `json:"fail_rate"`
	MinCalls             uint8 `json:"min_calls"`
}
