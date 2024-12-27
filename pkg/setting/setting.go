package setting

import "time"

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath   string
	LogSaveName   string
	LogFileExt    string
	TimeFormat    string
	FileMaxSize   int
	FileAllowExts []string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

func Setup() {
	AppSetting.JwtSecret = "NozfkxSLpMQACMgKwm#rfcFEfSMytD2B@&FJT!XQ6qdf5#uGN5"
	AppSetting.PageSize = 10
	AppSetting.PrefixUrl = "http://127.0.0.1:3966"

	// AppSetting.RuntimeRootPath = "runtime/"

	AppSetting.LogSavePath = "logs/"
	AppSetting.LogSaveName = "log"
	AppSetting.LogFileExt = "log"
	AppSetting.TimeFormat = "20060102"

	ServerSetting.RunMode = "debug"
	ServerSetting.HttpPort = 3000
	ServerSetting.ReadTimeout = 60
	ServerSetting.WriteTimeout = 60

	DatabaseSetting.Type = "sqlserver"
	DatabaseSetting.User = "sa"
	DatabaseSetting.Password = "12345"
	DatabaseSetting.Host = "KHOI:1433"
	DatabaseSetting.Name = "Blueprint"
	DatabaseSetting.TablePrefix = ""

	// AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	AppSetting.ImageMaxSize = 50 * 1024 * 1024 // 50MB
	AppSetting.ImageSavePath = "upload/files/"
	AppSetting.FileMaxSize = 50 * 1024 * 1024 // 50MB
	AppSetting.FileAllowExts = []string{".jpg", ".jpeg", ".png", ".pdf", ".txt"}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
}
