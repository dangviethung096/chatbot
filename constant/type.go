package constant

type LogoutStatus string

const (
	Success LogoutStatus = "Success"
	Fail    LogoutStatus = "Fail"
)

type LeadSource string

const (
	LeadSource_Facebook LeadSource = "facebook"
	LeadSource_Zalo     LeadSource = "zalo"
)

type LeadStatus string

const (
	LeadStatus_New   LeadStatus = "new"
	LeadStatus_Done  LeadStatus = "done"
	LeadStatus_Reuse LeadStatus = "reuse"
	LeadStatus_Lost  LeadStatus = "lost"
	LeadStatus_Win   LeadStatus = "win"
)

type ActivityType string

const (
	ActivityType_Message ActivityType = "message"
)

type ActivitySource string

const (
	ActivitySource_FacebookMessage ActivitySource = "facebook-message"
	ActivitySource_ZaloMessage     ActivitySource = "zalo-message"
)

var RoleMap = map[int64]Role{
	1: Role_Admin,
	2: Role_Director,
	3: Role_RegionalManager,
	4: Role_Manager,
	5: Role_Staff,
}

type Role struct {
	Id   int64
	Name string
}

var (
	Role_Admin           Role = Role{Id: 1, Name: "role_admin"}
	Role_Director        Role = Role{Id: 2, Name: "role_director"}
	Role_RegionalManager Role = Role{Id: 3, Name: "role_regional_manager"}
	Role_Manager         Role = Role{Id: 4, Name: "role_manager"}
	Role_Staff           Role = Role{Id: 5, Name: "role_staff"}
)

type OrderStatus string

const (
	OrderStatus_New     OrderStatus = "new"
	OrderStatus_Done    OrderStatus = "done"
	OrderStatus_Lost    OrderStatus = "lost"
	OrderStatus_Receive OrderStatus = "receive"
)

var UnitMap = map[string]string{
	"unit_bo":       UnitType_Bo,
	"unit_cai":      UnitType_Cai,
	"unit_cap":      UnitType_Cap,
	"unit_cay":      UnitType_Cay,
	"unit_chiec":    UnitType_Chiec,
	"unit_hop":      UnitType_Hop,
	"unit_md":       UnitType_Md,
	"unit_thanh":    UnitType_Thanh,
	"unit_vong":     UnitType_Vong,
	"unit_dientich": UnitType_Dientich,
	"unit_doi":      UnitType_Doi,
	"unit_cum":      UnitType_Cum,
}

const (
	UnitType_Bo       string = "unit_bo"
	UnitType_Cai      string = "unit_cai"
	UnitType_Cap      string = "unit_cap"
	UnitType_Cay      string = "unit_cay"
	UnitType_Chiec    string = "unit_chiec"
	UnitType_Hop      string = "unit_hop"
	UnitType_Md       string = "unit_md"
	UnitType_Thanh    string = "unit_thanh"
	UnitType_Vong     string = "unit_vong"
	UnitType_Dientich string = "unit_dientich"
	UnitType_Doi      string = "unit_doi"
	UnitType_Cum      string = "unit_cum"
)

type SubscriberType string

var SubscriberTypeMap = map[string]SubscriberType{
	"TT": SubscriberType_DEBIT,
	"TS": SubscriberType_CREDIT,
}

const SubscriberType_DEBIT SubscriberType = "TRA_TRUOC"
const SubscriberType_CREDIT SubscriberType = "TRA_SAU"

type CustomerType string

var CustomerTypeMap = map[string]CustomerType{
	"KHCN": CustomerType_CA_NHAN,
	"KHDN": CustomerType_DOANH_NGHIEP,
}

const CustomerType_CA_NHAN CustomerType = "KH_CA_NHAN"
const CustomerType_DOANH_NGHIEP CustomerType = "KH_DOANH_NGHIEP"

const (
	TaskStatus_TODO        string = "TODO"
	TaskStatus_DONE        string = "DONE"
	TaskStatus_IN_PROGRESS string = "IN_PROGRESS"
	TaskStatus_CLOSED      string = "CLOSED"
	TaskStatus_ALL         string = "ALL"
)

type TaskType string

const (
	TaskType_EXPIRING_CUSTOMER TaskType = "EXPIRING_CUSTOMER"
)

type ExpiringCustomerStatus string

const (
	ExpiringCustomerStatus_NeedSupport ExpiringCustomerStatus = "need_support"
	ExpiringCustomerStatus_Supported   ExpiringCustomerStatus = "supported"
	ExpiringCustomerStatus_Renewed     ExpiringCustomerStatus = "renewed"
	ExpiringCustomerStatus_None        ExpiringCustomerStatus = ""
)

const (
	SubscriberStatus_ACTIVE           string = "00"
	SubscriberStatus_BLOCK1WAY1DEBIT  string = "11"
	SubscriberStatus_BLOCK1WAY2DEBIT  string = "12"
	SubscriberStatus_BLOCK1WAY        string = "01"
	SubscriberStatus_BLOCK2WAY        string = "02"
	SubscriberStatus_BLOCK2WAY2DEBIT  string = "22"
	SubscriberStatus_BLOCK1WAYREQUIRE string = "10"
	SubscriberStatus_BLOCK2WAYREQUIRE string = "20"
	SubscriberStatus_BLOCK2WAY1CREDIT string = "21"
)

var SubscriberStatusList = []string{
	SubscriberStatus_ACTIVE,
	SubscriberStatus_BLOCK1WAY1DEBIT,
	SubscriberStatus_BLOCK1WAY2DEBIT,
	SubscriberStatus_BLOCK1WAY,
	SubscriberStatus_BLOCK2WAY,
	SubscriberStatus_BLOCK2WAY2DEBIT,
	SubscriberStatus_BLOCK1WAYREQUIRE,
	SubscriberStatus_BLOCK2WAYREQUIRE,
	SubscriberStatus_BLOCK2WAY1CREDIT,
}

var SubscriberStatusMap = map[string]string{
	SubscriberStatus_ACTIVE:           "Bình thường",
	SubscriberStatus_BLOCK1WAY1DEBIT:  "Chặn 1 chiều KHYC + 1 chiều nợ cước",
	SubscriberStatus_BLOCK1WAY2DEBIT:  "Chặn 1 chiều KHYC + 2 chiều nợ cước",
	SubscriberStatus_BLOCK1WAY:        "Chặn 1 chiều nợ cước",
	SubscriberStatus_BLOCK2WAY:        "Chặn 2 chiều nợ cước",
	SubscriberStatus_BLOCK2WAY2DEBIT:  "Chặn 2 chiều KHYC + 2 chiều nợ cước",
	SubscriberStatus_BLOCK1WAYREQUIRE: "Chặn 1 chiều KHYC",
	SubscriberStatus_BLOCK2WAYREQUIRE: "Chặn 2 chiều KHYC",
	SubscriberStatus_BLOCK2WAY1CREDIT: "Chặn 2 chiều KHYC + 1 chiều nợ cước",
	"ALL":                             "Tất cả",
}

const (
	PackageType_SinglePeriod string = "don_ky"
	PackageType_LongTerm     string = "dai_ky"
)
