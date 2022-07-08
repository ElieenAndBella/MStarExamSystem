/*
 * Created Date: Friday, July 8th 2022, 12:10:07 pm
 * Author: ZhanG
 */
package constant

type (
	// MongoDB 配置
	MongoConfig struct {
		User string
		Pass string
		Addr string
		Port string
	}
	// Redis 配置
	RedisConfig struct {
		Addr string
		Port string
		Pass string
		DB   int
	}
)

var (
	MC MongoConfig
	RC RedisConfig
)

func init() {
	// MongoDB初始化配置
	MC = MongoConfig{
		User: "",
		Pass: "",
		Addr: "192.168.50.1",
		Port: "27017",
	}
	// Redis初始化配置
	RC = RedisConfig{
		Addr: "192.168.50.1",
		Port: "6379",
		Pass: "",
		DB:   2,
	}
}
