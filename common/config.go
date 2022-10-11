package common

import(
	"github.com/spf13/viper"
	"os"
)

// InitConfig 初始化配置
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("faild to read config file, err: "+ err.Error())
	}
}
