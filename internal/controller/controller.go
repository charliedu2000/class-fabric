package controller

import "github.com/google/wire"

// ControllerSet 控制器 DI
var ControllerSet = wire.NewSet(
	LoginSet,
	ConsumerSet,
	AdministratorSet,
	DistributorSet,
	ManufacturerSet,
	ContractControllerSet,
)
