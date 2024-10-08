package routers

import (
	"be-hoatieu/pkg/upload"
	v1 "be-hoatieu/routers/api/v1"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitRouter(r *fiber.App) {

	apiv1 := r.Group("/api/v1")

	// // dang ky
	// apiv1.Post("/signup", v1.SignUp)
	// // dang nhap
	// apiv1.Post("/signin", v1.SignIn)

	apiv1.Use(cors.New(cors.ConfigDefault))         // CORS
	apiv1.Use(compress.New(compress.ConfigDefault)) // Compress
	apiv1.Use(logger.New(logger.ConfigDefault))     // Logger
	apiv1.Use(requestid.New())
	fmt.Println(upload.GetImageFullPath())
	apiv1.Static("/upload/files", upload.GetImageFullPath(), fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})

	// JWT check
	// apiv1.Use(jwtCustom.Protected())
	// apiv1.Use(jwtCustom.Authorization())
	// Route
	apiv1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👌!")
	})

	// Home page
	home := apiv1.Group("/home")
	home.Get("/carousel", v1.GetCarousel_Component)
	home.Get("/introduction", v1.GetIntroduction_Component)
	home.Get("/product", v1.GetServiceProduct_Component)
	home.Get("/news", v1.GetNews_Component)

	// Category Navigator
	hoatieu := apiv1.Group("/hoatieu")
	hoatieu.Get("/", v1.GetAllNavigator_Component)

	// Category Maneuvering Draft
	maneuveringDraft := apiv1.Group("/maneuvering-draft")
	maneuveringDraft.Get("/", v1.GetManeuveringDraft_Component)

	// Category product Price
	productPrice := apiv1.Group("/product-price")
	productPrice.Get("/", v1.GetProductPrice_Component)

	// Ship list
	ship := apiv1.Group("/ship")
	ship.Get("/", v1.GetShip_Component)

	//Tide Calendar
	tideCalendar := apiv1.Group("/tide-calendar")
	tideCalendar.Get("/", v1.GetTideCalendar_Component)

	//Stream
	stream := apiv1.Group("/stream")
	stream.Get("/", v1.GetStream_Component)
}

// 	// User
// 	userRoute := apiv1.Group("/user")
// 	userRoute.Get("/", v1.GetAllUser_Router)
// 	userRoute.Get("/:id", v1.GetByIdUser_Router)
// 	userRoute.Get("/username/:username", v1.GetByUserName_Router)
// 	userRoute.Put("/:id", v1.PutUsers)
// 	userRoute.Get("/trangthai/status", v1.GetAllUserTrue_Router)
// 	userRoute.Put("/pass/:id", v1.PutPasswordUsers)
// 	userRoute.Put("/image/:id", v1.PutIamgeUsers)
// 	userRoute.Put("/image/certificate/:id", v1.PutImageCertificateUsers)
// 	userRoute.Put("/info/:id", v1.PutDetailUsers)
// 	userRoute.Get("/search/role", v1.SearcUserRole_Component)
// 	userRoute.Post("/daily", v1.CreateUserDaiLy)

// 	//Image
// 	apiv1.Post("/upload", api.UploadFileSingle)
// 	apiv1.Post("/upload/multiple", api.UploadFileMultiple)

// 	calamviec := apiv1.Group("/calamviec")
// 	calamviec.Post("/", v1.CreateCaLamViec_Component)
// 	calamviec.Get("/", v1.GetAllCaLamViec_Component)
// 	calamviec.Put("/:id", v1.UpdateCaLamViec_Component)
// 	calamviec.Delete("/:id", v1.DeleteCaLamViec_Component)

// 	thoigian2daumaytheokhuvuc := apiv1.Group("/thoigian2daumaytheokhuvuc")
// 	thoigian2daumaytheokhuvuc.Post("/", v1.CreateThoiGian2DauMayTheoKhuVuc_Component)
// 	thoigian2daumaytheokhuvuc.Get("/", v1.GetAllThoiGian2DauMayTheoKhuVuc_Component)
// 	thoigian2daumaytheokhuvuc.Put("/:id", v1.UpdateThoiGian2DauMayTheoKhuVuc_Component)

// 	catruc := apiv1.Group("/catruc")
// 	catruc.Post("/", v1.CreateCaTruc_Component)
// 	catruc.Get("/", v1.GetAllCaTruc_Component)
// 	catruc.Put("/:id", v1.UpdateCaTruc_Component)
// 	catruc.Delete("/:id", v1.DeleteCaTruc_Component)
// 	catruc.Get("/search", v1.SearchCaTruc_Component)
// 	catruc.Get("/trangthai/status", v1.GetAllCaTrucTrue_Component)

// 	truccadieuhanh := apiv1.Group("/truccadieuhanh")
// 	truccadieuhanh.Post("/", v1.CreateTrucCaDieuHanh_Component)
// 	truccadieuhanh.Get("/", v1.GetAllTrucCaDieuHanh_Component)
// 	truccadieuhanh.Put("/:id", v1.UpdateTrucCaDieuHanh_Component)
// 	truccadieuhanh.Delete("/:id", v1.DeleteTrucCaDieuHanh_Component)
// 	truccadieuhanh.Get("/search", v1.SearchTrucCaDieuHanh_Component)
// 	truccadieuhanh.Get("/trangthai/status", v1.GetAllTrucCaDieuHanhTrue_Component)
// 	truccadieuhanh.Get("/ngayexcel", v1.GetAllTrucCaDieuHanhExcelNgayTTrue_Component)

// 	truccadieudo := apiv1.Group("/truccadieudo")
// 	truccadieudo.Post("/", v1.CreateTrucCaDieuDo_Component)
// 	truccadieudo.Get("/", v1.GetAllTrucCaDieuDo_Component)
// 	truccadieudo.Put("/:id", v1.UpdateTrucCaDieuDo_Component)
// 	truccadieudo.Delete("/:id", v1.DeleteTrucCaDieuDo_Component)
// 	truccadieudo.Get("/search", v1.SearchTrucCaDieuDo_Component)
// 	truccadieudo.Get("/trangthai/status", v1.GetAllTrucCaDieuDoTrue_Component)
// 	truccadieudo.Get("/ngayexcel", v1.GetAllTrucCaDieuDoExcelNgayTTrue_Component)

// 	truccano := apiv1.Group("/truccano")
// 	truccano.Post("/", v1.CreateTrucCano_Component)
// 	truccano.Get("/", v1.GetAllTrucCano_Component)
// 	truccano.Put("/:id", v1.UpdateTrucCano_Component)
// 	truccano.Delete("/:id", v1.DeleteTrucCano_Component)
// 	truccano.Get("/search", v1.SearchTrucCano_Component)
// 	truccano.Get("/trangthai/status", v1.GetAllTrucCanoTrue_Component)
// 	truccano.Get("/ngayexcel", v1.GetAllTrucCanoExcelNgayTTrue_Component)

// 	cang := apiv1.Group("/cang")
// 	cang.Get("/", v1.GetAllCang_Component)
// 	cang.Post("/", v1.CreateCang_Component)
// 	cang.Put("/:id", v1.UpdateCang_Component)
// 	cang.Get("/search", v1.SearchCang_Component)
// 	cang.Get("/trangthai/status", v1.GetAllCangTrue_Component)

// 	cangtuyen := apiv1.Group("/cangtuyen")
// 	cangtuyen.Get("/", v1.GetAllCangTuyen_Component)
// 	cangtuyen.Post("/", v1.CreateCangTuyen_Component)
// 	cangtuyen.Put("/:id", v1.UpdateCangTuyen_Component)
// 	cangtuyen.Get("/search", v1.SearchCangTuyen_Component)
// 	cangtuyen.Get("/multiple/cangtuyen", v1.GetMultipleCangTuyenTrue_Component)
// 	cangtuyen.Get("/trangthai/status", v1.GetAllCangTuyenTrue_Component)

// 	hang := apiv1.Group("/hang")
// 	hang.Get("/", v1.GetAllHang_Component)
// 	hang.Post("/", v1.CreateHang_Component)
// 	hang.Put("/:id", v1.UpdateHang_Component)
// 	hang.Get("/trangthai/status", v1.GetAllHangTrue_Component)
// 	hang.Get("/tau/trangthai/status", v1.GetAllHangTrueTau_Component)

// 	phieudoica := apiv1.Group("/phieudoica")
// 	phieudoica.Get("/", v1.GetAllPhieuDoiCa_Component)
// 	phieudoica.Post("/", v1.CreatePhieuDoiCa_Component)
// 	phieudoica.Put("/:id", v1.UpdatePhieuDoiCa_Component)
// 	phieudoica.Put("/approval/:id", v1.UpdateApprovalPhieuDoiCa_Component)
// 	phieudoica.Get("/trangthai/status", v1.GetAllPhieuDoiCaTrue_Component)
// 	phieudoica.Get("/approval/", v1.GetAllPhieuDoiCaApproval_Component)
// 	phieudoica.Delete("/:id", v1.DeletePhieuDoiCa_Component)

// 	phieudoitau := apiv1.Group("/phieudoitau")
// 	phieudoitau.Get("/", v1.GetAllPhieuDoiTau_Component)
// 	phieudoitau.Post("/", v1.CreatePhieuDoiTau_Component)
// 	phieudoitau.Put("/:id", v1.UpdatePhieuDoiTau_Component)
// 	phieudoitau.Put("/approval/:id", v1.UpdateApprovalPhieuDoiTau_Component)
// 	phieudoitau.Get("/trangthai/status", v1.GetAllPhieuDoiTauTrue_Component)
// 	phieudoitau.Get("/approval/", v1.GetAllPhieuDoiTauApproval_Component)
// 	phieudoitau.Delete("/:id", v1.DeletePhieuDoiTau_Component)

// 	phieucongtac := apiv1.Group("/phieucongtac")
// 	phieucongtac.Get("/", v1.GetAllPhieuCongTac_Component)
// 	phieucongtac.Post("/", v1.CreatePhieuCongTac_Component)
// 	phieucongtac.Put("/:id", v1.UpdatePhieuCongTac_Component)
// 	phieucongtac.Put("/approval/:id", v1.UpdateApprovalPhieuCongTac_Component)
// 	phieucongtac.Get("/trangthai/status", v1.GetAllPhieuCongTacTheoDoi_Component)
// 	phieucongtac.Get("/approval/", v1.GetAllPhieuCongTacApproval_Component)
// 	phieucongtac.Delete("/:id", v1.DeletePhieuCongTac_Component)

// 	chungchihoatieu := apiv1.Group("/chungchihoatieu")
// 	chungchihoatieu.Get("/", v1.GetAllChungChiHoaTieu_Component)
// 	chungchihoatieu.Post("/", v1.CreateChungChiHoaTieu_Component)
// 	chungchihoatieu.Put("/:id", v1.UpdateChungChiHoaTieu_Component)
// 	chungchihoatieu.Get("/trangthai/status", v1.GetAllChungChiHoaTieuTrue_Component)

// 	chucvu := apiv1.Group("/chucvu")
// 	chucvu.Get("/", v1.GetAllChucVu_Component)
// 	chucvu.Post("/", v1.CreateChucVu_Component)
// 	chucvu.Put("/:id", v1.UpdateChucVu_Component)
// 	chucvu.Get("/trangthai/status", v1.GetAllChucVuTrue_Component)

// 	bpctac := apiv1.Group("/bpctac")
// 	bpctac.Get("/", v1.GetAllBPCTac_Component)
// 	bpctac.Post("/", v1.CreateBPCTac_Component)
// 	bpctac.Put("/:id", v1.UpdateBPCTac_Component)
// 	bpctac.Get("/trangthai/status", v1.GetAllBPCTacTrue_Component)

// 	hoatieu := apiv1.Group("/hoatieu")
// 	hoatieu.Get("/", v1.GetAllHoaTieu_Component)
// 	hoatieu.Get("/userid/:id", v1.GetAllHoaTieuUserID_Component)
// 	hoatieu.Get("/byhoatieuidu/:id", v1.GetByHoaTieuUserID_Component)
// 	hoatieu.Post("/", v1.CreateUserAndHoaTieu)
// 	hoatieu.Put("/:id", v1.PutUserAndHoaTieu)
// 	hoatieu.Get("/search/", v1.SearchHoaTieu_Component)
// 	hoatieu.Get("/search/update", v1.SearchUpDateHoaTieu_Component)
// 	hoatieu.Get("/searchht2/update", v1.SearchHT2UpDateHoaTieu_Component)
// 	hoatieu.Get("/searchhtbangkesanluong/update", v1.SearchHTBangKeSanLuongUpDateHoaTieu_Component)
// 	hoatieu.Get("/search/role", v1.SearchHoaTieuRole_Component)
// 	hoatieu.Get("/multiple/hoatieu", v1.GetMultipleHoaTieuTrue_Component)

// 	hohieu := apiv1.Group("/hohieu")
// 	hohieu.Get("/", v1.GetAllHoHieu_Component)
// 	hohieu.Post("/", v1.CreateHoHieu_Component)
// 	hohieu.Put("/:id", v1.UpdateHoHieu_Component)
// 	hohieu.Get("/:id", v1.GetIDHoHieu_Component)
// 	hohieu.Get("/trangthai/status", v1.GetAllHoHieuTrue_Component)

// 	daily := apiv1.Group("/daily")
// 	daily.Get("/", v1.GetAllDaiLy_Component)
// 	daily.Post("/", v1.CreateDaiLy_Component)
// 	daily.Put("/:id", v1.UpdateDaiLy_Component)
// 	daily.Get("/trangthai/status", v1.GetAllDaiLyTrue_Component)

// 	khachhang := apiv1.Group("/khachhang")
// 	khachhang.Get("/", v1.GetAllKhachHang_Component)
// 	khachhang.Post("/", v1.CreateKhachHang_Component)
// 	khachhang.Put("/:id", v1.UpdateKhachHang_Component)
// 	khachhang.Get("/search", v1.SearchKhachHang_Component)
// 	//khachhang.Get("/searchbyid", v1.SearchKhachHangById_Component)
// 	khachhang.Get("/:id", v1.GetIDKhachHang_Component)
// 	//khachhang.Get("/:id/daily",v1.GetOrderById_Component) // Route để lấy thông tin khách hàng và đại lý (join giữa khách hàng và đại lý)

// 	doituongxuathoadon := apiv1.Group("/doituongxuathoadon")
// 	doituongxuathoadon.Get("/", v1.GetAllDoiTuongXuatHoaDon_Component)
// 	doituongxuathoadon.Post("/", v1.CreateDoiTuongXuatHoaDon_Component)
// 	doituongxuathoadon.Put("/:id", v1.UpdateDoiTuongXuatHoaDonn_Component)
// 	doituongxuathoadon.Get("/search", v1.SearchDoiTuongXuatHoaDon_Component)
// 	doituongxuathoadon.Get("/:id", v1.GetIDDoiTuongXuatHoaDon_Component)

// 	kehoach := apiv1.Group("/kehoach")
// 	kehoach.Get("/theodoi", v1.GetAllKeHoachTheoDoi_Component)
// 	kehoach.Get("/approval", v1.GetAllKeHoachApproval_Component)
// 	kehoach.Post("/", v1.CreateKeHoach_Component)
// 	kehoach.Put("/:id", v1.UpdateKeHoach_Component)
// 	kehoach.Get("/:id", v1.GetIDKeHoach_Component)
// 	kehoach.Delete("/:id", v1.DeleteKeHoach_Component)

// 	loaikhachhang := apiv1.Group("/loaikhachhang")
// 	loaikhachhang.Get("/", v1.GetAllLoaiKhachHang_Component)
// 	loaikhachhang.Post("/", v1.CreateLoaiKhachHang_Component)
// 	loaikhachhang.Put("/:id", v1.UpdateLoaiKhachHang_Component)
// 	loaikhachhang.Get("/:id", v1.GetIDLoaiKhachHang_Component)
// 	loaikhachhang.Get("/trangthai/status", v1.GetAllLoaiKhachHangTrue_Component)

// 	khunghoatdonghoatieu := apiv1.Group("/khunghoatdonghoatieu")
// 	khunghoatdonghoatieu.Get("/", v1.GetAllKhungGioHoatDongHoaTieu_Component)
// 	khunghoatdonghoatieu.Post("/", v1.CreateKhungGioHoatDongHoaTieu_Component)
// 	khunghoatdonghoatieu.Put("/:id", v1.UpdateKhungGioHoatDongHoaTieu_Component)

// 	laixe := apiv1.Group("/laixe")
// 	laixe.Get("/", v1.GetAllLaiXe_Component)
// 	laixe.Post("/", v1.CreateLaiXe_Component)
// 	laixe.Put("/:id", v1.UpdateLaiXe_Component)
// 	laixe.Get("/trangthai/status", v1.GetAllLaiXeTrue_Component)

// 	loaihanghoa := apiv1.Group("/loaihanghoa")
// 	loaihanghoa.Get("/", v1.GetAllLoaiHangHoa_Component)
// 	loaihanghoa.Post("/", v1.CreateLoaiHangHoa_Component)
// 	loaihanghoa.Put("/:id", v1.UpdateLoaiHangHoa_Component)
// 	loaihanghoa.Get("/:id", v1.GetIDLoaiHangHoa_Component)
// 	loaihanghoa.Get("/trangthai/status", v1.GetAllLoaiHangHoaTrue_Component)

// 	monnuoc := apiv1.Group("/monnuoc")
// 	monnuoc.Get("/", v1.GetAllMonNuoc_Component)
// 	monnuoc.Post("/", v1.CreateMonNuoc_Component)
// 	monnuoc.Put("/:id", v1.UpdateMonNuoc_Component)
// 	monnuoc.Get("/:id", v1.GetIDMonNuoc_Component)

// 	nghiphep := apiv1.Group("/nghiphep")
// 	nghiphep.Get("/", v1.GetAllNghiPhep_Component)
// 	nghiphep.Post("/", v1.CreateNghiPhep_Component)
// 	nghiphep.Put("/:id", v1.UpdateNghiPhep_Component)

// 	phuongthucthanhtoan := apiv1.Group("/phuongthucthanhtoan")
// 	phuongthucthanhtoan.Get("/", v1.GetAllPhuongThucThanhToan_Component)
// 	phuongthucthanhtoan.Post("/", v1.CreatePhuongThucThanhToan_Component)
// 	phuongthucthanhtoan.Put("/:id", v1.UpdatePhuongThucThanhToan_Component)
// 	phuongthucthanhtoan.Get("/trangthai/status", v1.GetAllPhuongThucThanhToanTrue_Component)

// 	tau := apiv1.Group("/tau")
// 	tau.Get("/", v1.GetAllTau_Component)
// 	tau.Post("/", v1.CreateTau_Component)
// 	tau.Put("/:id", v1.UpdateTau_Component)
// 	tau.Get("/search", v1.SearchTau_Component)

// 	tuyen := apiv1.Group("/tuyen")
// 	tuyen.Get("/", v1.GetAllTuyen_Component)
// 	tuyen.Post("/", v1.CreateTuyen_Component)
// 	tuyen.Put("/:id", v1.UpdateTuyen_Component)
// 	tuyen.Get("/search", v1.SearchTuyen_Component)

// 	hangphancong := apiv1.Group("/hangphancong")
// 	hangphancong.Get("/", v1.GetAllPhanCongHangDan_Component)
// 	hangphancong.Post("/", v1.CreatePhanCongHangDan_Component)
// 	hangphancong.Put("/:id", v1.UpdatePhanCongHangDan_Component)
// 	hangphancong.Get("/search", v1.SearchPhanCongHangDan_Component)

// 	thoigian := apiv1.Group("/thoigian")
// 	thoigian.Get("/", v1.GetAllThoiGian_Component)
// 	thoigian.Post("/", v1.CreateThoiGian_Component)
// 	thoigian.Put("/:id", v1.UpdateThoiGian_Component)
// 	thoigian.Get("/search", v1.SearchThoiGian_Component)

// 	thongbao := apiv1.Group("/thongbao")
// 	thongbao.Get("/", v1.GetAllThongBao_Component)
// 	thongbao.Post("/", v1.CreateThongBao_Component)
// 	thongbao.Put("/:id", v1.UpdateThongBao_Component)

// 	trangthailaixe := apiv1.Group("/trangthailaixe")
// 	trangthailaixe.Get("/", v1.GetAllTrangThaiLaiXe_Component)
// 	trangthailaixe.Get("/userid/", v1.GetAllTrangThaiLaiXeUser_Component)
// 	trangthailaixe.Post("/", v1.CreateTrangThaiLaiXe_Component)
// 	trangthailaixe.Put("/:id", v1.UpdateTrangThaiLaiXe_Component)
// 	trangthailaixe.Put("/trangthaiduadon/:id", v1.UpdateTrangThaiDuaDon_Component)

// 	trangthaicano := apiv1.Group("/trangthaicano")
// 	trangthaicano.Get("/", v1.GetAllTrangThaiCaNo_Component)
// 	trangthaicano.Get("/userid/", v1.GetAllTrangThaiCaNoUser_Component)
// 	trangthaicano.Post("/", v1.CreateTrangThaiCaNo_Component)
// 	trangthaicano.Put("/:id", v1.UpdateTrangThaiCaNo_Component)
// 	trangthaicano.Put("/trangthaiduadon/:id", v1.UpdateTrangThaiCaNoDuaDon_Component)

// 	taulaicano := apiv1.Group("/taulaicano")
// 	taulaicano.Get("/", v1.GetAllTauLaiCano_Component)
// 	taulaicano.Post("/", v1.CreateTauLaiCano_Component)
// 	taulaicano.Get("/trangthai/status", v1.GetAllTauLaiCanoTrue_Component)
// 	taulaicano.Get("/loai/status", v1.GetAllTauLaiCanoTrueLoai_Component)
// 	taulaicano.Put("/:id", v1.UpdateTauLaiCano_Component)
// 	taulaicano.Get("/multiple/taulaicanodepluong", v1.GetMultipleTauLaiAndCanoDepLuongTrue_Component)

// 	notification := apiv1.Group("/notification")
// 	notification.Get("/:id", v1.GetAllNotification_Component)
// 	notification.Post("/", v1.CreateNotification_Component)
// 	notification.Put("/:id", v1.UpdateNotification_Component)

// 	tuahoatieu := apiv1.Group("/tuahoatieu")
// 	tuahoatieu.Get("/", v1.GetAllTuaHoaTieu_Component)
// 	tuahoatieu.Post("/", v1.CreateTuaHoaTieu_Component)
// 	tuahoatieu.Put("/:id", v1.UpdateTuaHoaTieu_Component)
// 	tuahoatieu.Get("/trangthai/status", v1.GetAllTuaHoaTieuTrue_Component)
// 	tuahoatieu.Get("/search/", v1.SearchTuaHoaTieu_Component)

// 	tuahoatieu_detail := apiv1.Group("/tuahoatieudetail")
// 	tuahoatieu_detail.Get("/", v1.GetAllTuaHoaTieu_Details_Component)
// 	tuahoatieu_detail.Post("/", v1.CreateTuaHoaTieu_Details_Component)
// 	tuahoatieu_detail.Put("/:id", v1.UpdateTuaHoaTieu_Details_Component)
// 	//tuahoatieu_detail.Get("/trangthai/status", v1.GetAllTuaHoaTieu_DetailTrue_Component)

// 	nhomtour := apiv1.Group("/nhomtour")
// 	nhomtour.Get("/", v1.GetAllNhomTour_Component)
// 	nhomtour.Post("/", v1.CreateNhomTour_Component)
// 	nhomtour.Put("/:id", v1.UpdateNhomTour_Component)
// 	nhomtour.Get("/", v1.GetAllNhomTour_Component)
// 	nhomtour.Get("/trangthai/status", v1.GetAllNhomTourTrue_Component)

// 	loaitau := apiv1.Group("/loaitau")
// 	loaitau.Get("/", v1.GetAllLoaiTau_Component)
// 	loaitau.Post("/", v1.CreateLoaiTau_Component)
// 	loaitau.Put("/:id", v1.UpdateLoaiTau_Component)
// 	loaitau.Get("/:id", v1.GetIDLoaiTau_Component)
// 	loaitau.Get("/trangthai/status", v1.GetAllLoaiTauTrue_Component)

// 	ghichu := apiv1.Group("/ghichu")
// 	ghichu.Get("/", v1.GetAllGhiChu_Component)
// 	ghichu.Post("/", v1.CreateGhiChu_Component)
// 	ghichu.Put("/:id", v1.UpdateGhiChu_Component)
// 	ghichu.Get("/search", v1.GetAllGhiChuTrue_Component)

// 	diadiem := apiv1.Group("/diadiem")
// 	diadiem.Get("/", v1.GetAllDiaDiem_Component)
// 	diadiem.Post("/", v1.CreateDiaDiem_Component)
// 	diadiem.Put("/:id", v1.UpdateDiaDiem_Component)
// 	diadiem.Get("/search", v1.GetAllDiaDiemTrue_Component)

// 	loaihinhkinhdoanh := apiv1.Group("/loaihinhkinhdoanh")
// 	loaihinhkinhdoanh.Get("/", v1.GetAllLoaiHinhKinhDoanh_Component)
// 	loaihinhkinhdoanh.Post("/", v1.CreateLoaiHinhKinhDoanh_Component)
// 	loaihinhkinhdoanh.Put("/:id", v1.UpdateLoaiHinhKinhDoanh_Component)
// 	loaihinhkinhdoanh.Get("/search", v1.GetAllLoaiHinhKinhDoanhTrue_Component)

// 	danhsachdiem := apiv1.Group("/danhsachdiem")
// 	danhsachdiem.Get("/", v1.GetAllDanhSachDiem_Component)
// 	danhsachdiem.Get("/diemcang", v1.GetDiemCang_Component)
// 	danhsachdiem.Post("/", v1.CreateDanhSachDiem_Component)
// 	danhsachdiem.Put("/:id", v1.UpdateDanhSachDiem_Component)

// 	cotluong := apiv1.Group("/cotluong")
// 	cotluong.Get("/", v1.GetAllCotLuong_Component)
// 	cotluong.Post("/", v1.CreateCotLuong_Component)
// 	cotluong.Put("/:id", v1.UpdateCotLuong_Component)
// 	cotluong.Get("/trangthai/status", v1.GetAllCotLuongTrue_Component)

// 	hoatdongtour := apiv1.Group("/hoatdongtour")
// 	hoatdongtour.Get("/", v1.GetAllHoatDongTour_Component)
// 	hoatdongtour.Post("/", v1.CreateHoatDongTour_Component)
// 	hoatdongtour.Put("/:id", v1.UpdateHoatDongTour_Component)
// 	hoatdongtour.Get("/trangthai/status", v1.GetAllHoatDongTourTrue_Component)
// 	hoatdongtour.Get("/search/", v1.SearchHoatDongTour_Component)

// 	camketthanhtoan := apiv1.Group("/camketthanhtoan")
// 	camketthanhtoan.Get("/", v1.GetAllCamKetThanhToan_Component)
// 	camketthanhtoan.Post("/", v1.CreateCamKetThanhToan_Component)
// 	camketthanhtoan.Put("/:id", v1.UpdateCamKetThanhToan_Component)
// 	camketthanhtoan.Get("/trangthai/status", v1.GetAllCamKetThanhToanTrue_Component)

// 	donhang := apiv1.Group("/donhang")
// 	donhang.Get("/", v1.GetAllDonHang_Component)
// 	donhang.Get("/tinh", v1.TinhMonNuoc)
// 	donhang.Get("/tinhthoigianthuytrieu", v1.TinhThuyTrieuCangTuyen)
// 	donhang.Get("/tinhthoigianthuytrieudichchuyen", v1.TinhThuyTrieuCangTuyenDichChuyen)
// 	donhang.Post("/", v1.CreateDonHang_Component)
// 	donhang.Post("/shipping", v1.CreateShippingDonHang_Component)
// 	donhang.Put("/:id", v1.UpdateDonHang_Component)
// 	donhang.Put("/donhangcertificate/:id", v1.UpdateDonHangCertificate_Component)
// 	donhang.Put("/approval/:id", v1.UpdateApprovalDonHang_Component)
// 	donhang.Put("/updatelaixe/:id", v1.UpdateDonHangLaiXe_Component)
// 	donhang.Put("/updatecano/:id", v1.UpdateDonHangCano_Component)
// 	donhang.Put("/updatelaixeve/:id", v1.UpdateDonHangLaiXeVe_Component)
// 	donhang.Put("/updatecanove/:id", v1.UpdateDonHangCanoVe_Component)
// 	donhang.Put("/bangkesanluong/approval/:id", v1.UpdateApprovalBangKe_Component)
// 	donhang.Put("/khuyennghihoatieu/:id", v1.UpdateHoaTieuDonHang_Component)
// 	donhang.Get("/order/monitor", v1.GetAllDonHangMonitor_Component)
// 	//donhang.Get("/bangke/approval", v1.GetBangKeSanLuongApproval_Component)
// 	donhang.Get("/donhanghuy/:id", v1.GetAllDonHangTheoDoiHuyDon_Component)
// 	donhang.Get("/alldonhanghuy/", v1.GetAllDonHangHuy_Component)
// 	donhang.Get("/alldonhangchapnhan/", v1.GetAllDonHangChapNhan_Component)
// 	donhang.Get("/order/bangkesanluong", v1.GetBangKeSanLuongMonitor_Component)
// 	donhang.Get("/order/certificate/", v1.GetAllDonHangCertificate_Component)
// 	donhang.Get("/dartboard", v1.GetAllDonHangDartBoard_Component)
// 	donhang.Put("/lydohuy/:id", v1.LyDoHuyDonHang_Component)
// 	donhang.Put("/donhanghuy/:id", v1.UpdateDonHangHuy_Component)
// 	donhang.Get("/theodoi", v1.GetAllDonHangTheoDoi_Component)
// 	donhang.Get("/nametuyen", v1.GetTuyenName_Component)
// 	donhang.Get("/khuyennghihoatieu", v1.KhuyenNghiHoaTieu_Component)
// 	donhang.Get("/hoatieutuor", v1.HoaTieuTour)
// 	donhang.Get("/theodoibyiduser", v1.GetAllDonHangTheoDoiByIdUser_Component)
// 	donhang.Get("/alldonhanghuybyiduser/", v1.GetAllDonHangHuyByIdUser_Component)
// 	// donhang.Get("/sotua", v1.CountHoatieuAppearances_Component)

// 	quoctich := apiv1.Group("/quoctich")
// 	quoctich.Get("/", v1.GetAllQuocTich_Component)
// 	quoctich.Post("/", v1.CreateQuocTich_Component)
// 	quoctich.Put("/:id", v1.UpdateQuocTich_Component)
// 	quoctich.Get("/:id", v1.GetIDQuocTich_Component)
// 	quoctich.Get("/trangthai/status", v1.GetAllQuocTichTrue_Component)

// 	cano := apiv1.Group("/cano")
// 	cano.Get("/", v1.GetAllCaNo_Component)
// 	cano.Post("/", v1.CreateCaNo_Component)
// 	cano.Put("/:id", v1.UpdateCaNo_Component)
// 	cano.Get("/trangthai/status", v1.GetAllCaNoTrue_Component)
// 	cano.Get("/search/", v1.SearchCaNo_Component)
// 	cano.Get("/multiple/cano", v1.GetMultipleCanoTrue_Component)

// 	xe := apiv1.Group("/xe")
// 	xe.Get("/", v1.GetAllXe_Component)
// 	xe.Post("/", v1.CreateXe_Component)
// 	xe.Put("/:id", v1.UpdateXe_Component)
// 	xe.Get("/trangthai/status", v1.GetAllXeTrue_Component)
// 	xe.Get("/search/", v1.SearchXe_Component)

// 	nhomcano := apiv1.Group("/nhomcano")
// 	nhomcano.Get("/", v1.GetAllNhomCaNo_Component)
// 	nhomcano.Post("/", v1.CreateNhomCaNo_Component)
// 	nhomcano.Put("/:id", v1.UpdateNhomCaNo_Component)
// 	nhomcano.Get("/trangthai/status", v1.GetAllNhomCaNoTrue_Component)

// 	bangdiem := apiv1.Group("/bangdiem")
// 	bangdiem.Get("/", v1.GetAllBangDiem_Component)
// 	bangdiem.Get("/ngay", v1.GetAllBangDiemNgay_Component)
// 	bangdiem.Post("/", v1.CreateBangDiem_Component)
// 	bangdiem.Post("/hoatieu2", v1.CreateBangDiem2_Component)
// 	// bangdiem.Get("/:id", v1.TinhDiemCongTac_Component)
// 	bangdiem.Get("/tinhdiem", v1.TinhDiemTuorDay_Component)

// 	quanlycano := apiv1.Group("/quanlycano")
// 	quanlycano.Get("/", v1.GetAllQuanLyCaNo_Component)
// 	quanlycano.Get("/trangthai/status", v1.GetAllQuanLyCaNoTrue_Component)
// 	quanlycano.Post("/", v1.CreateQuanLyCaNo_Component)
// 	quanlycano.Put("/:id", v1.UpdateQuanLyCaNo_Component)

// 	thongtincongty := apiv1.Group("/genaral")
// 	thongtincongty.Get("/", v1.GetAllThongTinCongTy_Component)
// 	thongtincongty.Post("/", v1.CreateThongTinCongTy_Component)
// 	thongtincongty.Put("/:id", v1.UpdateThongTinCongTy_Component)
// 	thongtincongty.Put("/image/:id", v1.PutIamgeThongTinCongTy_Component)

// 	certificate := apiv1.Group("/certificate")
// 	certificate.Get("/theodoi", v1.GetAllCertificateTheoDoi_Component)
// 	certificate.Post("/", v1.CreateCertificate_Component)
// 	certificate.Put("/approval/:id", v1.UpdateApprovalCertificate_Component)
// 	certificate.Put("/:id", v1.UpdateCertificate_Component)
// 	certificate.Get("/approval/", v1.GetAllCertificateApproval_Component)
// 	certificate.Get("/order/bangkesanluong/", v1.GetAllCertificateLaiXeVeFalse_Component)

// 	certificate.Put("/updatelaixeve/:id", v1.UpdateCertificateLaiXeVe_Component)
// 	certificate.Put("/updatecanove/:id", v1.UpdateCertificateCanoVe_Component)

// 	report := apiv1.Group("/report")
// 	report.Get("/", v1.GetChartDashboard_Router)

// 	thuytrieu := apiv1.Group("/thuytrieu")
// 	thuytrieu.Get("/execl", v1.GetThuyTrieuExecl_Component)
// 	thuytrieu.Get("/", v1.GetAllThuyTrieu_Component)
// 	// thuytrieu.Post("/", v1.TruyTrieu_Component)
// 	thuytrieu.Post("/", v1.ThuyTrieu_Component)
// 	thuytrieu.Post("/themmoi", v1.CreateThuyTrieu_Component)
// 	thuytrieu.Put("/:id", v1.UpdateThuyTrieu_Component)

// 	phoihop := apiv1.Group("/phoihop")
// 	phoihop.Get("/", v1.GetAllPhoiHop_Component)
// 	phoihop.Post("/", v1.CreatePhoiHop_Component)
// 	phoihop.Put("/:id", v1.UpdatePhoiHop_Component)
// 	phoihop.Get("/trangthai/status", v1.GetAllPhoiHopTrue_Component)
// 	phoihop.Get("/search/", v1.SearchPhoiHop_Component)

// 	// khu vuc
// 	khuvuc := apiv1.Group("/khuvuc")
// 	khuvuc.Post("/", v1.CreateKhuvuc_Component)
// 	khuvuc.Put("/:id", v1.UpdateKhuvuc_Component)
// 	khuvuc.Get("/", v1.GetAllKhuvuc_Component)
// 	khuvuc.Get("/search/", v1.SearchKhuVuc_Component)
// }
