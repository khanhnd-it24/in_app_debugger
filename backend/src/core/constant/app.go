package constant

import "backend/src/common/configs"

const (
	AppEnvDev  = "dev"
	AppEnvProd = "prod"
)

func IsProdEnv() bool {
	return configs.Get().Mode == AppEnvProd
}
