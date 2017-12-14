package stackpointio

// this must go away

type mspec struct {
	cpu    int
	memory int
}

func getMachineTypeMemory(machineType string) int {
	return allSpecs[machineType].memory
}

func getMachineTypeCPU(machineType string) int {
	return allSpecs[machineType].cpu
}

var allSpecs = map[string]mspec{
	"n1-standard-1":  mspec{1, 3750},
	"n1-standard-2":  mspec{2, 7500},
	"n1-standard-4":  mspec{4, 15000},
	"n1-standard-8":  mspec{8, 30000},
	"n1-standard-16": mspec{16, 60000},
	"n1-standard-32": mspec{32, 12000},
	"n1-standard-64": mspec{64, 240000},

	"n1-highmem-2":  mspec{2, 13000},
	"n1-highmem-4":  mspec{4, 26000},
	"n1-highmem-8":  mspec{8, 52000},
	"n1-highmem-16": mspec{16, 104000},
	"n1-highmem-32": mspec{32, 208000},

	"n1-highcpu-2":  mspec{2, 1800},
	"n1-highcpu-4":  mspec{4, 3600},
	"n1-highcpu-8":  mspec{8, 7200},
	"n1-highcpu-16": mspec{16, 14400},
	"n1-highcpu-32": mspec{32, 28800},
	"n1-highcpu-64": mspec{64, 57600},

	"t2.medium":  mspec{2, 4000},
	"t2.large":   mspec{2, 8000},
	"t2.xlarge":  mspec{4, 16000},
	"t2.2xlarge": mspec{8, 32000},

	"c4.large":   mspec{2, 3750},
	"c4.xlarge":  mspec{4, 7500},
	"c4.2xlarge": mspec{8, 15000},
	"c4.4xlarge": mspec{16, 30000},
	"c4.8xlarge": mspec{36, 60000},

	"m3.large":   mspec{2, 7500},
	"m3.xlarge":  mspec{4, 15000},
	"m3.2xlarge": mspec{8, 30000},

	"m4.large":    mspec{2, 8000},
	"m4.xlarge":   mspec{4, 16000},
	"m4.2xlarge":  mspec{8, 32000},
	"m4.4xlarge":  mspec{16, 64000},
	"m4.10xlarge": mspec{40, 160000},
	"m4.16xlarge": mspec{64, 256000},

	"2gb": mspec{2, 2000},
	"4gb": mspec{2, 4000},
	"8gb": mspec{4, 8000},
}
