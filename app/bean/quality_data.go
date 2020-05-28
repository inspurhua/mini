package bean

import (
	"huage.tech/mini/app/config"
	"time"
)

type QualityData struct {
	ID     int64  `json:"id" gorm:"primary_key"`
	KeyId  int64  `json:"key_id" gorm:"key_id"`
	Batch  string `json:"batch" form:"batch" gorm:"batch"`
	KeyT00 string `json:"key_t00" form:"key_t00" gorm:"key_t00"`
	KeyT01 string `json:"key_t01" form:"key_t01" gorm:"key_t01"`
	KeyT02 string `json:"key_t02" form:"key_t02" gorm:"key_t02"`
	KeyT03 string `json:"key_t03" form:"key_t03" gorm:"key_t03"`
	KeyT04 string `json:"key_t04" form:"key_t04" gorm:"key_t04"`

	DataT00 string `json:"data_t00" form:"data_t00" gorm:"data_t00"`
	DataT01 string `json:"data_t01" form:"data_t01" gorm:"data_t01"`
	DataT02 string `json:"data_t02" form:"data_t02" gorm:"data_t02"`
	DataT03 string `json:"data_t03" form:"data_t03" gorm:"data_t03"`
	DataT04 string `json:"data_t04" form:"data_t04" gorm:"data_t04"`
	DataT05 string `json:"data_t05" form:"data_t05" gorm:"data_t05"`
	DataT06 string `json:"data_t06" form:"data_t06" gorm:"data_t06"`
	DataT07 string `json:"data_t07" form:"data_t07" gorm:"data_t07"`
	DataT08 string `json:"data_t08" form:"data_t08" gorm:"data_t08"`
	DataT09 string `json:"data_t09" form:"data_t09" gorm:"data_t09"`
	DataT10 string `json:"data_t10" form:"data_t10" gorm:"data_t10"`
	DataT11 string `json:"data_t11" form:"data_t11" gorm:"data_t11"`
	DataT12 string `json:"data_t12" form:"data_t12" gorm:"data_t12"`
	DataT13 string `json:"data_t13" form:"data_t13" gorm:"data_t13"`
	DataT14 string `json:"data_t14" form:"data_t14" gorm:"data_t14"`
	DataT15 string `json:"data_t15" form:"data_t15" gorm:"data_t15"`
	DataT16 string `json:"data_t16" form:"data_t16" gorm:"data_t16"`
	DataT17 string `json:"data_t17" form:"data_t17" gorm:"data_t17"`
	DataT18 string `json:"data_t18" form:"data_t18" gorm:"data_t18"`
	DataT19 string `json:"data_t19" form:"data_t19" gorm:"data_t19"`

	DataI00 int `json:"data_i00" form:"data_i00" gorm:"data_i00"`
	DataI01 int `json:"data_i01" form:"data_i01" gorm:"data_i01"`
	DataI02 int `json:"data_i02" form:"data_i02" gorm:"data_i02"`
	DataI03 int `json:"data_i03" form:"data_i03" gorm:"data_i03"`
	DataI04 int `json:"data_i04" form:"data_i04" gorm:"data_i04"`
	DataI05 int `json:"data_i05" form:"data_i05" gorm:"data_i05"`
	DataI06 int `json:"data_i06" form:"data_i06" gorm:"data_i06"`
	DataI07 int `json:"data_i07" form:"data_i07" gorm:"data_i07"`
	DataI08 int `json:"data_i08" form:"data_i08" gorm:"data_i08"`
	DataI09 int `json:"data_i09" form:"data_i09" gorm:"data_i09"`
	DataI10 int `json:"data_i10" form:"data_i10" gorm:"data_i10"`
	DataI11 int `json:"data_i11" form:"data_i11" gorm:"data_i11"`
	DataI12 int `json:"data_i12" form:"data_i12" gorm:"data_i12"`
	DataI13 int `json:"data_i13" form:"data_i13" gorm:"data_i13"`
	DataI14 int `json:"data_i14" form:"data_i14" gorm:"data_i14"`
	DataI15 int `json:"data_i15" form:"data_i15" gorm:"data_i15"`
	DataI16 int `json:"data_i16" form:"data_i16" gorm:"data_i16"`
	DataI17 int `json:"data_i17" form:"data_i17" gorm:"data_i17"`
	DataI18 int `json:"data_i18" form:"data_i18" gorm:"data_i18"`
	DataI19 int `json:"data_i19" form:"data_i19" gorm:"data_i19"`

	DataN00 float64 `json:"data_n00" form:"data_n00" gorm:"data_n00"`
	DataN01 float64 `json:"data_n01" form:"data_n01" gorm:"data_n01"`
	DataN02 float64 `json:"data_n02" form:"data_n02" gorm:"data_n02"`
	DataN03 float64 `json:"data_n03" form:"data_n03" gorm:"data_n03"`
	DataN04 float64 `json:"data_n04" form:"data_n04" gorm:"data_n04"`
	DataN05 float64 `json:"data_n05" form:"data_n05" gorm:"data_n05"`
	DataN06 float64 `json:"data_n06" form:"data_n06" gorm:"data_n06"`
	DataN07 float64 `json:"data_n07" form:"data_n07" gorm:"data_n07"`
	DataN08 float64 `json:"data_n08" form:"data_n08" gorm:"data_n08"`
	DataN09 float64 `json:"data_n09" form:"data_n09" gorm:"data_n09"`
	DataN10 float64 `json:"data_n10" form:"data_n10" gorm:"data_n10"`
	DataN11 float64 `json:"data_n11" form:"data_n11" gorm:"data_n11"`
	DataN12 float64 `json:"data_n12" form:"data_n12" gorm:"data_n12"`
	DataN13 float64 `json:"data_n13" form:"data_n13" gorm:"data_n13"`
	DataN14 float64 `json:"data_n14" form:"data_n14" gorm:"data_n14"`
	DataN15 float64 `json:"data_n15" form:"data_n15" gorm:"data_n15"`
	DataN16 float64 `json:"data_n16" form:"data_n16" gorm:"data_n16"`
	DataN17 float64 `json:"data_n17" form:"data_n17" gorm:"data_n17"`
	DataN18 float64 `json:"data_n18" form:"data_n18" gorm:"data_n18"`
	DataN19 float64 `json:"data_n19" form:"data_n19" gorm:"data_n19"`

	QcDate   string    `json:"qc_date" form:"qc_date" gorm:"qc_date"`
	CreateAt time.Time `json:"create_at" form:"-" gorm:"create_at"`
	CreateBy int64     `json:"create_by" form:"-" gorm:"create_by"`
	TenantId int64     `json:"tenant_id" form:"-" gorm:"tenant_id"`
}

func (QualityData) TableName() string {
	return config.Prefix + "quality"
}
