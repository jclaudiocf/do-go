package cli

type DeviceDto struct {
	Id			string		`json:"id"`
	Nickname	string		`json:"nickname"`
	Owner		string		`json:"owner"`
	Ip			string		`json:"ip"`
}

type DevicePageable struct {
	Devices 	[]DeviceDto		`json:"devices"`
	Page 		int				`json:"page"`
	Size 		int				`json:"size"`
}
