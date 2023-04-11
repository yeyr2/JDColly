package config

const LocalHost = "127.0.0.1"
const Host = "www.iyeyr2.top"
const User = "root"
const Pass = "root"
const SqlPort = "3306"

//var Host = "127.0.0.1:9090" // 获取ip地址
//func GetOutBoundIP() (ip string, err error) {
//	conn, err := net.Dial("udp", "8.8.8.8:53")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	localAddr := conn.LocalAddr().(*net.UDPAddr)
//	ip = strings.Split(localAddr.String(), ":")[0]
//	return
//}
//
//func init() {
//	ip, err := GetOutBoundIP()
//	if err != nil {
//		log.Fatalln("获取ip失败:", err)
//	}
//	Host = ip
//}
