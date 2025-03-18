package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Config)

// Config represents the configuration structure for the application.
type Config struct {
	Name        string       `mapstructure:"name"`
	Mode        string       `mapstructure:"mode"`
	Port        int          `mapstructure:"port"`
	StartTime   string       `mapstructure:"start_time"`
	MachineID   int64        `mapstructure:"machine_id"`
	LogConfig   *LogConfig   `mapstructure:"log"`
	MysqlConfig *MySQLConfig `mapstructure:"mysql"`
	RedisConfig *RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {            // 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed,err: #%v\n", err)
		return
	}
	err = viper.Unmarshal(Conf)
	if err != nil {
		return err
	}
	// 监控配置文件变化
	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		// 配置文件发生变化，重新读取配置信息
		fmt.Println("配置文件发生变化")
		err = viper.Unmarshal(Conf)
		if err != nil {
			return
		}
	})
	return
}
