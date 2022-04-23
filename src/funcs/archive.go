package funcs

import (
	"github.com/cihub/seelog"
	"github.com/YongboStudio/satagent/src/common"
	"strconv"
)

//clear timeout alert table
func ClearArchive() {
	seelog.Info("[func:ClearArchive] ", "starting run ClearArchive ")
	common.DLock.Lock()
	common.Db.Exec("delete from alertlog where logtime < date('now','start of day','-" + strconv.Itoa(common.Cfg.Base["Archive"]) + " day')")
	common.Db.Exec("delete from mappinglog where logtime < date('now','start of day','-" + strconv.Itoa(common.Cfg.Base["Archive"]) + " day')")
	common.Db.Exec("delete from pinglog where logtime < date('now','start of day','-" + strconv.Itoa(common.Cfg.Base["Archive"]) + " day')")
	common.DLock.Unlock()
	seelog.Info("[func:ClearArchive] ", "ClearArchive Finish ")
}
