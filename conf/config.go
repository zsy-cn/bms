// Package conf 使用viper读取环境变量中的配置并赋值给Config结构
package conf

var (
	CameraServiceAddr      = "camera-serv"
	CameraServicePort      = ":1201"
	CameraServiceYS7Addr   = ""
	CameraServiceYS7AppKey = ""
	CameraServiceYS7Secret = ""

	MediaCoreServiceAddr = "116.62.162.253"
	MediaCoreServicePort = ":6003"

	ParserHubServiceAddr    = "parserhub-serv"
	ParserHubServicePort    = ":1202"
	ManagerServiceAddr      = "manager-serv"
	ManagerServicePort      = ":1203"
	CoreServiceAddr         = "core-serv"
	CoreServicePort         = ":1204"
	DeviceHubServiceAddr    = "devicehub-serv"
	DeviceHubServicePort    = ":1205"
	DeviceSensorServiceAddr = "device-sensor-serv"
	DeviceSensorServicePort = ":1206"
	CustomerServiceAddr     = "customer-serv"
	CustomerServicePort     = ":1207"
	ContactServiceAddr      = "contact-serv"
	ContactServicePort      = ":1208"
	LoraclientServiceAddr   = "loraclient-serv"
	LoraclientServicePort   = ":1209"

	LoraAppServerAddr = "appserver-serv"
	LoraAppServerPort = ":8080"

	RedisAddr            = "redis-serv"
	RedisPort            = "6379"
	RedisPassword        = ""
	RedisCameraServiceDB = 1

	TrashcanHeight = float64(125) // 单位: 米

	ParserEnvYingfeichi01BindAddr    = ":1210"
	ParserEnvYingfeichi01ServiceAddr = "envyingfeichi-serv"
	ParserEnvYingfeichi01ServicePort = "1210"

	ParserGeomagneticWeichuan01BindAddr    = ":1211"
	ParserGeomagneticWeichuan01ServiceAddr = "geomagneticweichuan01-serv"
	ParserGeomagneticWeichuan01ServicePort = "1211"

	ParserTrashcanLierda01BindAddr    = ":1212"
	ParserTrashcanLierda01ServiceAddr = "trashcanlierda01-serv"
	ParserTrashcanLierda01ServicePort = "1212"

	ParserTemperatureWeichuan01BindAddr    = ":1213"
	ParserTemperatureWeichuan01ServiceAddr = "temperatureweichuan01-serv"
	ParserTemperatureWeichuan01ServicePort = "1213"

	ParserManholeCoverGti01BindAddr    = ":1214"
	ParserManholeCoverGti01ServiceAddr = "manholecovergti01-serv"
	ParserManholeCoverGti01ServicePort = "1214"

	ParserGeomagneticVchuan01BindAddr    = ":1215"
	ParserGeomagneticVchuan01ServiceAddr = "geomagneticvchuan01-serv"
	ParserGeomagneticVchuan01ServicePort = "1215"

	ParserManholeCoverNanpeng01BindAddr    = ":1216"
	ParserManholeCoverNanpeng01ServiceAddr = "manholecovernanpeng01-serv"
	ParserManholeCoverNanpeng01ServicePort = "1216"
)
