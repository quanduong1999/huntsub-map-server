package station

type SystemConfig struct {
	TempPath string
}

func (sc *SystemConfig) check() {
	if sc.TempPath == "" {
		sc.TempPath = "./static/temp"
	}
}
