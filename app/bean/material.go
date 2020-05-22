package bean

import "huage.tech/mini/app/config"

type Material struct {
	ID             int64  `json:"id" gorm:"primary_key"`
	Code           string `form:"code" json:"code" binding:"required"`
	Name           string `form:"name" json:"name" binding:"required"`
	Specs          string `form:"specs" json:"specs" binding:"required"`
	PackSpecs      string `form:"pack_specs" json:"pack_specs" binding:"required"`
	Batch          string `form:"batch" json:"batch" binding:"required"`
	BatchUnit      string `form:"batch_unit" json:"batch_unit" binding:"required"`
	Storage        string `form:"storage" json:"storage" binding:"required"`
	Validity       string `form:"validity" json:"validity" binding:"required"`
	Standard       string `form:"standard" json:"standard" binding:"required"`
	ApprovalNumber string `form:"approval_number" json:"approval_number" binding:"required"`
	QualityName    string `form:"quality_name" json:"quality_name" binding:"required"`
	ProcessName    string `form:"process_name" json:"process_name" binding:"required"`
	ProcessFile    int64  `form:"process_file" json:"process_file"`
	TypeId         int64  `form:"type_id" json:"type_id"`
	Status         int    `form:"status" json:"status" default:1 binding:"oneof=1 2"`
	TenantId       int64  `json:"-"`
}

func (Material) TableName() string {
	return config.Prefix + "material"
}

type MaterialWithFile struct {
	Material
	Url  string `json:"url"`
	File string `json:"file"`
}

func (MaterialWithFile) TableName() string {
	return config.Prefix + "material"
}
