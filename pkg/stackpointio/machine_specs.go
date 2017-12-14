package stackpointio

// --> this must go away

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
	"n1-standard-32": mspec{32, 120000},
	"n1-standard-64": mspec{64, 240000},

	"n1-highmem-2":  mspec{2, 13000},
	"n1-highmem-4":  mspec{4, 26000},
	"n1-highmem-8":  mspec{8, 52000},
	"n1-highmem-16": mspec{16, 104000},
	"n1-highmem-32": mspec{32, 208000},
	"n1-highmem-64": mspec{64, 416000},

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
	"m4.large":   mspec{2, 8000},
	"m4.xlarge":  mspec{4, 16000},

	"m4.2xlarge":  mspec{8, 32000},
	"m4.4xlarge":  mspec{16, 64000},
	"m4.10xlarge": mspec{40, 160000},
	"m4.16xlarge": mspec{64, 256000},

	"2gb":  mspec{2, 2000},
	"4gb":  mspec{2, 4000},
	"8gb":  mspec{4, 8000},
	"16gb": mspec{8, 16000},
	"32gb": mspec{12, 32000},
	"48gb": mspec{16, 32000},
	"64gb": mspec{20, 64000},

	"c-2":     mspec{2, 3000},
	"c-4":     mspec{4, 6000},
	"c-8":     mspec{8, 12000},
	"c-16":    mspec{16, 32000},
	"c-32":    mspec{32, 64000},
	"m-16gb":  mspec{2, 16000},
	"m-32gb":  mspec{4, 32000},
	"m-64gb":  mspec{8, 64000},
	"m-128gb": mspec{16, 128000},
	"m-224gb": mspec{32, 224000},

	"M":   mspec{1, 2000},
	"L":   mspec{2, 2000},
	"XL":  mspec{2, 4000},
	"XXL": mspec{4, 8000},
	"3XL": mspec{8, 16000},
	"4XL": mspec{12, 32000},

	"2,INTEL_XEON,4096,50,SSD": mspec{2, 4000},

	"standard_a2": mspec{2, 3500},
	"standard_a3": mspec{4, 7000},
	"standard_a4": mspec{8, 14000},
	"standard_a5": mspec{2, 14000},
	"standard_a6": mspec{4, 28000},
	"standard_a7": mspec{8, 56000},
	"standard_a8": mspec{8, 56000},
	"standard_a9": mspec{16, 112000},

	"standard_a10": mspec{8, 56000},
	"standard_a11": mspec{16, 112000},

	"basic_a2": mspec{2, 3500},
	"basic_a3": mspec{4, 7000},
	"basic_a4": mspec{8, 14000},

	"standard_a2_v2":  mspec{2, 4000},
	"standard_a4_v2":  mspec{4, 8000},
	"standard_a8_v2":  mspec{8, 16000},
	"standard_a2m_v2": mspec{2, 16000},
	"standard_a4m_v2": mspec{4, 32000},
	"standard_a8m_v2": mspec{8, 64000},

	"standard_d1": mspec{1, 3500},
	"standard_d2": mspec{2, 7000},
	"standard_d3": mspec{4, 14000},
	"standard_d4": mspec{8, 28000},

	"standard_d11": mspec{2, 14000},
	"standard_d12": mspec{4, 28000},
	"standard_d13": mspec{8, 56000},
	"standard_d14": mspec{16, 112000},

	"standard_ds1": mspec{1, 3500},
	"standard_ds2": mspec{2, 7000},
	"standard_ds3": mspec{4, 14000},
	"standard_ds4": mspec{8, 28000},

	"standard_d1_v2": mspec{1, 3500},
	"standard_d2_v2": mspec{2, 7000},
	"standard_d3_v2": mspec{4, 14000},
	"standard_d4_v2": mspec{8, 28000},
	"standard_d5_v2": mspec{16, 56000},

	"standard_d11_v2": mspec{2, 14000},
	"standard_d12_v2": mspec{4, 28000},
	"standard_d13_v2": mspec{8, 56000},
	"standard_d14_v2": mspec{16, 112000},
	"standard_d15_v2": mspec{20, 140000},

	"standard_ds1_v2": mspec{1, 3500},
	"standard_ds2_v2": mspec{2, 7000},
	"standard_ds3_v2": mspec{4, 14000},
	"standard_ds4_v2": mspec{8, 28000},
	"standard_ds5_v2": mspec{16, 56000},

	"standard_ds11_v2": mspec{2, 14000},
	"standard_ds12_v2": mspec{4, 28000},
	"standard_ds13_v2": mspec{8, 56000},
	"standard_ds14_v2": mspec{16, 112000},
	"standard_ds15_v2": mspec{20, 140000},

	"standard_f1":  mspec{1, 2000},
	"standard_f2":  mspec{2, 4000},
	"standard_f4":  mspec{4, 8000},
	"standard_f8":  mspec{8, 16000},
	"standard_f16": mspec{16, 32000},

	"standard_f1s":  mspec{1, 2000},
	"standard_f2s":  mspec{2, 4000},
	"standard_f4s":  mspec{4, 8000},
	"standard_f8s":  mspec{8, 16000},
	"standard_f16s": mspec{16, 32000},

	"standard_g1": mspec{2, 28000},
	"standard_g2": mspec{4, 56000},
	"standard_g3": mspec{8, 112000},
	"standard_g4": mspec{16, 224000},
	"standard_g5": mspec{32, 448000},

	"standard_gs1": mspec{2, 28000},
	"standard_gs2": mspec{4, 56000},
	"standard_gs3": mspec{8, 112000},
	"standard_gs4": mspec{16, 224000},
	"standard_gs5": mspec{32, 448000},
}
