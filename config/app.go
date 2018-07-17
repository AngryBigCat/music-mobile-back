package config

const (
	DBTYPE = "mysql"
	DBHOST = "59.110.20.181"
	DBNAME = "music"
	DBUSER = "root"
	DBPASS = "root"
	DBPORT = "3306"

	LISTEN_PORT = "8080"

	REDIS_HOST = "localhost:6379"
	REDIS_PASS = ""
	REDIS_DB = 0
	// 短信验证码时效300秒
	REDIS_SENDCODE_EX = 300
)