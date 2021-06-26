package services

import (
	"awesomeProject1/modules"
	"awesomeProject1/underly"
)

func GetAllOutboundlist() (error, []modules.OunboundList) {
	err, outboundlistData := underly.GetAllOutboundlist()
	return err, outboundlistData
}
