package rank

type RankName string

type Level int
type Radius int
type PostAndShareNumber int
type ExperienceNumber int

/* Define Rank Name for each LevelUp*/
const (
	NguoiMoi_I   = RankName("Nguoi Moi I")
	NguoiMoi_II  = RankName("Nguoi Moi II")
	NguoiMoi_III = RankName("Nguoi Moi III")

	PhongTrao_I   = RankName("Phong Trao I")
	PhongTrao_II  = RankName("Phong Trao II")
	PhongTrao_III = RankName("Phong Trao III")

	NghiepDu_I   = RankName("Nghiep Du I")
	NghiepDu_II  = RankName("Nghiep Du II")
	NghiepDu_III = RankName("Nghiep Du III")
	NghiepDu_IV  = RankName("Nghiep Du IV")

	ChuyenVien_I   = RankName("Chuyen Vien I")
	ChuyenVien_II  = RankName("Chuyen Vien Ii")
	ChuyenVien_III = RankName("Chuyen Vien III")
	ChuyenVien_IV  = RankName("Chuyen Vien IV")
	ChuyenVien_V   = RankName("Chuyen Vien V")

	ChuyenNghiep_I   = RankName("Chuyen Nghiep I")
	ChuyenNghiep_II  = RankName("Chuyen Nghiep II")
	ChuyenNghiep_III = RankName("Chuyen Nghiep III")
	ChuyenNghiep_IV  = RankName("Chuyen Nghiep IV")
	ChuyenNghiep_V   = RankName("Chuyen Nghiep V")

	ChuyenGia_I   = RankName("Chuyen Gia I")
	ChuyenGia_II  = RankName("Chuyen Gia II")
	ChuyenGia_III = RankName("Chuyen Gia III")
	ChuyenGia_IV  = RankName("Chuyen Gia IV")
	ChuyenGia_V   = RankName("Chuyen Gia V")

	CoVan_I   = RankName("Co Van I")
	CoVan_II  = RankName("Co Van II")
	CoVan_III = RankName("Co Van III")
	CoVan_IV  = RankName("Co Van IV")
	CoVan_V   = RankName("Co Van V")
)

const (
	/*Define Radius of RankName by meters */
	Radius_NguoiMoi_I   = Radius(500)
	Radius_NguoiMoi_II  = Radius(600)
	Radius_NguoiMoi_III = Radius(700)

	Radius_PhongTrao_I   = Radius(800)
	Radius_PhongTrao_II  = Radius(900)
	Radius_PhongTrao_III = Radius(1000)

	Radius_NghieDu_I   = Radius(1200)
	Radius_NghieDu_II  = Radius(1400)
	Radius_NghieDu_III = Radius(1600)
	Radius_NghieDu_IV  = Radius(1800)

	Radius_ChuyenVien_I   = Radius(2000)
	Radius_ChuyenVien_II  = Radius(2250)
	Radius_ChuyenVien_III = Radius(2500)
	Radius_ChuyenVien_IV  = Radius(2750)
	Radius_ChuyenVien_V   = Radius(3000)

	Radius_ChuyenNghiep_I   = Radius(3500)
	Radius_ChuyenNghiep_II  = Radius(4000)
	Radius_ChuyenNghiep_III = Radius(4500)
	Radius_ChuyenNghiep_IV  = Radius(5000)
	Radius_ChuyenNghiep_V   = Radius(5500)

	Radius_ChuyenGia_I   = Radius(6000)
	Radius_ChuyenGia_II  = Radius(7000)
	Radius_ChuyenGia_III = Radius(8000)
	Radius_ChuyenGia_IV  = Radius(9000)
	Radius_ChuyenGia_V   = Radius(10000)

	Radius_CoVan_I   = Radius(12000)
	Radius_CoVan_II  = Radius(14000)
	Radius_CoVan_III = Radius(16000)
	Radius_CoVan_IV  = Radius(18000)
	Radius_CoVan_V   = Radius(20000)
)

const (
	/*Define Number's Post and Share of RankName */
	PostAndShareNumber_NguoiMoi_I       = PostAndShareNumber(1)
	PostAndShareNumber_NguoiMoi_II      = PostAndShareNumber(1)
	PostAndShareNumber_NguoiMoi_III     = PostAndShareNumber(1)
	PostAndShareNumber_PhongTrao_I      = PostAndShareNumber(2)
	PostAndShareNumber_PhongTrao_II     = PostAndShareNumber(2)
	PostAndShareNumber_PhongTrao_III    = PostAndShareNumber(2)
	PostAndShareNumber_NghieDu_I        = PostAndShareNumber(2)
	PostAndShareNumber_NghieDu_II       = PostAndShareNumber(2)
	PostAndShareNumber_NghieDu_III      = PostAndShareNumber(2)
	PostAndShareNumber_NghieDu_IV       = PostAndShareNumber(2)
	PostAndShareNumber_ChuyenVien_I     = PostAndShareNumber(3)
	PostAndShareNumber_ChuyenVien_II    = PostAndShareNumber(3)
	PostAndShareNumber_ChuyenVien_III   = PostAndShareNumber(3)
	PostAndShareNumber_ChuyenVien_IV    = PostAndShareNumber(3)
	PostAndShareNumber_ChuyenVien_V     = PostAndShareNumber(3)
	PostAndShareNumber_ChuyenNghiep_I   = PostAndShareNumber(3)
	PostAndShareNumber_ChuyenNghiep_II  = PostAndShareNumber(3)
	PostAndShareNumber_ChuyenNghiep_III = PostAndShareNumber(3)
	PostAndShareNumber_ChuyenNghiep_IV  = PostAndShareNumber(3)
	PostAndShareNumber_ChuyenNghiep_V   = PostAndShareNumber(3)
	PostAndShareNumber_ChuyenGia_I      = PostAndShareNumber(4)
	PostAndShareNumber_ChuyenGia_II     = PostAndShareNumber(4)
	PostAndShareNumber_ChuyenGia_III    = PostAndShareNumber(4)
	PostAndShareNumber_ChuyenGia_IV     = PostAndShareNumber(4)
	PostAndShareNumber_ChuyenGia_V      = PostAndShareNumber(4)
	PostAndShareNumber_CoVan_I          = PostAndShareNumber(5)
	PostAndShareNumber_CoVan_II         = PostAndShareNumber(5)
	PostAndShareNumber_CoVan_III        = PostAndShareNumber(5)
	PostAndShareNumber_CoVan_IV         = PostAndShareNumber(5)
	PostAndShareNumber_CoVan_V          = PostAndShareNumber(5)
)

const (
	/*Define Number's Experience of RankName */
	ExperienceNumber_NguoiMoi_I   = ExperienceNumber(0)
	ExperienceNumber_NguoiMoi_II  = ExperienceNumber(10)
	ExperienceNumber_NguoiMoi_III = ExperienceNumber(25)

	ExperienceNumber_PhongTrao_I   = ExperienceNumber(45)
	ExperienceNumber_PhongTrao_II  = ExperienceNumber(75)
	ExperienceNumber_PhongTrao_III = ExperienceNumber(120)

	ExperienceNumber_NghieDu_I   = ExperienceNumber(180)
	ExperienceNumber_NghieDu_II  = ExperienceNumber(300)
	ExperienceNumber_NghieDu_III = ExperienceNumber(500)
	ExperienceNumber_NghieDu_IV  = ExperienceNumber(850)

	ExperienceNumber_ChuyenVien_I   = ExperienceNumber(1500)
	ExperienceNumber_ChuyenVien_II  = ExperienceNumber(2500)
	ExperienceNumber_ChuyenVien_III = ExperienceNumber(4000)
	ExperienceNumber_ChuyenVien_IV  = ExperienceNumber(7000)
	ExperienceNumber_ChuyenVien_V   = ExperienceNumber(12000)

	ExperienceNumber_ChuyenNghiep_I   = ExperienceNumber(20000)
	ExperienceNumber_ChuyenNghiep_II  = ExperienceNumber(35000)
	ExperienceNumber_ChuyenNghiep_III = ExperienceNumber(60000)
	ExperienceNumber_ChuyenNghiep_IV  = ExperienceNumber(100000)
	ExperienceNumber_ChuyenNghiep_V   = ExperienceNumber(150000)

	ExperienceNumber_ChuyenGia_I   = ExperienceNumber(220000)
	ExperienceNumber_ChuyenGia_II  = ExperienceNumber(300000)
	ExperienceNumber_ChuyenGia_III = ExperienceNumber(410000)
	ExperienceNumber_ChuyenGia_IV  = ExperienceNumber(540000)
	ExperienceNumber_ChuyenGia_V   = ExperienceNumber(690000)

	ExperienceNumber_CoVan_I   = ExperienceNumber(890000)
	ExperienceNumber_CoVan_II  = ExperienceNumber(1110000)
	ExperienceNumber_CoVan_III = ExperienceNumber(1350000)
	ExperienceNumber_CoVan_IV  = ExperienceNumber(1610000)
	ExperienceNumber_CoVan_V   = ExperienceNumber(2000000)
)
