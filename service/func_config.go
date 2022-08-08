package service

import (
	"FuseMiddleware/global"
	"FuseMiddleware/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func AddFuseConfig(req *http.Request) int {
	errorCode := 0
	name := req.PostFormValue("name")
	forceBreakerSwitch, _ := strconv.ParseUint(req.PostFormValue("force_breaker_switch"), 10, 64)
	breakerStrategy, _ := strconv.ParseUint(req.PostFormValue("breaker_strategy"), 10, 64)
	degradeStrategy, _ := strconv.ParseUint(req.PostFormValue("degrade_strategy"), 10, 64)
	degradeStrategyDetail := req.PostFormValue("degrade_strategy_detail")
	breakerStrategyDetail := req.PostFormValue("breaker_strategy_detail")

	formatStrategy := model.BreakerStrategy{}
	if err := json.Unmarshal([]byte(degradeStrategyDetail), &formatStrategy); err != nil {
		return 1
	}

	fuseConfig := model.FuseConfig{}

	fuseConfig.Name = name
	fuseConfig.ForceBreakerSwitch = uint8(forceBreakerSwitch)
	fuseConfig.BreakerStrategy = uint8(breakerStrategy)
	fuseConfig.DegradeStrategy = uint8(degradeStrategy)
	fuseConfig.BreakerStrategyDetail = breakerStrategyDetail
	fuseConfig.DegradeStrategyDetail = degradeStrategyDetail
	fuseConfig.IsDelete = 0
	if err := global.GVA_DB.Omit("CreateTime", "UpdateTime").Create(&fuseConfig).Error; err != nil {
		return 1
	}

	return errorCode
}
