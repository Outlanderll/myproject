package underly

import "awesomeProject1/modules"

func GetAllOutboundlist() (error, []modules.OunboundList) {
	var outboundlistData []modules.OunboundList
	err := db.Table("outbound_list").Order("parts_name ASC").Find(&outboundlistData).Error
	return err, outboundlistData
}
