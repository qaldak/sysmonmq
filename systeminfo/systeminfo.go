package systeminfo

type SystemInfo struct {
	CPU01          float64 `json:"CPU01"`
	CPU05          float64 `json:"CPU05"`
	CPU15          float64 `json:"CPU15"`
	CPU_temp       float64 `json:"CPU_temp"`
	RAM_total      uint64  `json:"RAM_total"`
	RAM_free       uint64  `json:"RAM_free"`
	RAM_avlbl      uint64  `json:"RAM_avlbl"`
	RAM_used       uint64  `json:"RAM_used"`
	Disk_total     uint64  `json:"Disk_total"`
	Disk_free      uint64  `json:"Disk_free"`
	Disk_used      uint64  `json:"Disk_used"`
	Sys_Uptime     uint64  `json:"Sys_Uptime"`
	LastLogin_date string  `json:"LastLogin_date"`
	LastLogin_user string  `json:"LastLogin_user"`
	LastLogin_from string  `json:"LastLogin_from"`
}
