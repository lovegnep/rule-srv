package constants

const(
	// 请病假
	SickLeave = iota + 1
	// 审批
	Approve
	// 病重
	CriticallyIll
	// 去医院
	GoToDoctor
)

// 事件的几种状态
const(
	// 初始
	EventStatusInit = iota + 1
	// 审批了
	EventStatusApproved
	// 病重了
	EventStatusIll
)
