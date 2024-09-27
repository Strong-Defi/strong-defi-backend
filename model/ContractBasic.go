package model

import (
	"gorm.io/gorm"
	"time"
)

type ContractBasic struct {
	ID                    uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ContractCode          string    `gorm:"type:varchar(200);not null" json:"contract_code"`
	ContractTitle         string    `gorm:"type:varchar(333);not null" json:"contract_title"`
	ContractName          string    `gorm:"type:varchar(333);not null" json:"contract_name"`
	ContractType          int16     `gorm:"not null" json:"contract_type"`
	ContractCategory      string    `gorm:"type:varchar(50);not null" json:"contract_category"`
	StageSnap             int16     `gorm:"not null" json:"stage_snap"`
	StartDate             time.Time `gorm:"type:datetime;not null" json:"start_date"`
	EndDate               time.Time `gorm:"type:datetime;not null" json:"end_date"`
	ProjectID             string    `gorm:"type:varchar(50);not null" json:"project_id"`
	ProjectName           string    `gorm:"type:varchar(300);not null" json:"project_name"`
	SecretFlag            bool      `gorm:"not null" json:"secret_flag"`
	EmergencyFlag         int8      `gorm:"not null" json:"emergency_flag"`
	TemplateFlag          int16     `gorm:"not null" json:"template_flag"`
	CompanyID             string    `gorm:"type:varchar(50);not null" json:"company_id"`
	CompanyName           string    `gorm:"type:varchar(333);not null" json:"company_name"`
	LevelType             int16     `gorm:"not null" json:"level_type"`
	PayType               int16     `gorm:"not null" json:"pay_type"`
	PayMethod             int16     `gorm:"not null" json:"pay_method"`
	CurrencyCode          string    `gorm:"type:varchar(20);not null" json:"currency_code"`
	Amount                float64   `gorm:"type:decimal(19,2);not null" json:"amount"`
	FinanceRemark         string    `gorm:"type:varchar(2500);not null" json:"finance_remark"`
	OrderID               string    `gorm:"type:varchar(50);not null" json:"order_id"`
	TaskID                string    `gorm:"type:varchar(50);not null" json:"task_id"`
	ApproveAds            string    `gorm:"type:varchar(255);not null" json:"approve_ads"`
	NodeName              string    `gorm:"type:varchar(100);not null" json:"node_name"`
	SubmitDate            time.Time `gorm:"type:datetime;not null" json:"submit_date"`
	ParentContractID      int64     `gorm:"not null" json:"parent_contract_id"`
	IsDeleted             bool      `gorm:"not null" json:"is_deleted"`
	Ctime                 time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"ctime"`
	Mtime                 time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"mtime"`
	ContractOwner         string    `gorm:"type:varchar(50);not null" json:"contract_owner"`
	CreatedBy             string    `gorm:"type:varchar(50);not null" json:"created_by"`
	BusinessCC            string    `gorm:"type:varchar(50);not null" json:"business_cc"`
	Business              string    `gorm:"type:varchar(500);not null" json:"business"`
	RelativeURL           string    `gorm:"type:varchar(255);not null" json:"relative_url"`
	IsTemplate            int8      `gorm:"not null" json:"is_template"`
	MemberCooperateAmount float64   `gorm:"type:decimal(13,2);not null" json:"member_cooperate_amount"`
	ContractOrderCode     string    `gorm:"type:varchar(50);not null" json:"contract_order_code"`
	NodeCtime             time.Time `gorm:"type:datetime;not null" json:"node_ctime"`
	IsFrameworkContract   int8      `gorm:"not null" json:"is_framework_contract"`
	Version               int       `gorm:"not null" json:"version"`
	EquivalentRMB         float64   `gorm:"type:decimal(19,2);not null" json:"equivalent_rmb"`
	RequestLevel          int8      `gorm:"not null" json:"request_level"`
	CorrelationCode       string    `gorm:"type:varchar(25);not null" json:"correlation_code"`
	OwnerMtime            time.Time `gorm:"type:datetime;not null" json:"owner_mtime"`
	RequestChannel        int8      `gorm:"not null" json:"request_channel"`
	IsOverseas            bool      `gorm:"not null" json:"is_overseas"`
	AgentSponsorAccount   string    `gorm:"type:varchar(50);not null" json:"agent_sponsor_account"`
	AgentFlag             int8      `gorm:"not null" json:"agent_flag"`
	CreatedByNickName     string    `gorm:"type:varchar(100);not null" json:"created_by_nick_name"`
	CreatedByDept         string    `gorm:"type:varchar(100);not null" json:"created_by_dept"`
	SourceID              int16     `gorm:"not null" json:"source_id"`
	DeptNameSeq           string    `gorm:"type:varchar(500);not null" json:"dept_name_seq"`
	TotalAmount           float64   `gorm:"type:decimal(20,2);not null" json:"total_amount"`
	TotalEquivalentRMB    float64   `gorm:"type:decimal(20,2);not null" json:"total_equivalent_rmb"`
	SequenceNumber        int16     `gorm:"not null" json:"sequence_number"`
	SonContractOrderCode  string    `gorm:"type:varchar(100);not null" json:"son_contract_order_code"`
	ContractStatus        int16     `gorm:"not null" json:"contract_status"`
	MajorContractID       int64     `gorm:"not null" json:"major_contract_id"`
	NftOrderNo            string    `gorm:"type:varchar(500);not null" json:"nft_order_no"`
	HighRisk              int8      `gorm:"not null" json:"high_risk"`
	HighRiskTip           string    `gorm:"type:varchar(1000);not null" json:"high_risk_tip"`
	OtherID               string    `gorm:"type:varchar(50);not null" json:"other_id"`
	CompanyCodeJD         string    `gorm:"type:varchar(50);not null" json:"company_code_jd"`
	CompanyCurrency       string    `gorm:"type:varchar(50);not null" json:"company_currency"`
	CompanyCurrencyCode   string    `gorm:"type:varchar(50);not null" json:"company_currency_code"`
	FirstDepartment       string    `gorm:"type:varchar(50);not null" json:"first_department"`
	FinanceType           string    `gorm:"type:varchar(100);not null" json:"finance_type"`
	TaxPayType            int16     `gorm:"not null" json:"tax_pay_type"`
	TaxFeeNature          int16     `gorm:"not null" json:"tax_fee_nature"`
	TaxAmountHongKong     float64   `gorm:"type:decimal(18,2);not null" json:"tax_amount_hongkong"`
	TaxAmountNotHongKong  float64   `gorm:"type:decimal(18,2);not null" json:"tax_amount_not_hongkong"`
	CompanyPlace          string    `gorm:"type:varchar(100);not null" json:"company_place"`
	ProjectCode           string    `gorm:"type:varchar(500);not null" json:"project_code"`
	TrackKeys             string    `gorm:"type:varchar(256);not null" json:"track_keys"`
}

func (ContractBasic) TableName() string {
	return "contract_basic"
}

func SelectUser(db *gorm.DB, id int64) ([]ContractBasic, error) {
	var contractBasic []ContractBasic
	result := db.Where("id = ?", id).First(&contractBasic)

	if result.Error != nil {
		panic(result.Error)
	}
	return contractBasic, nil
}

func CustomSelect(db *gorm.DB, query interface{}, args ...interface{}) ([]ContractBasic, error) {
	var contractBasic []ContractBasic
	result := db.Where(query, args).Find(&contractBasic)

	if len(contractBasic) <= 0 {
		return contractBasic, result.Error
	}
	return contractBasic, nil
}

func Create(db *gorm.DB, contractBasic *ContractBasic) {
	result := db.Create(contractBasic)

	if result.Error != nil {
		panic(result.Error)
	}
}

func Update(db *gorm.DB, contractName string, query interface{}, queryArgs ...interface{}) (int, error) {
	var contractBasic ContractBasic
	result := db.Model(contractBasic).Where(query, queryArgs).Update("contract_name", contractName)

	if result.Error != nil {
		return 0, result.Error
	} else {
		return 1, nil
	}
}
