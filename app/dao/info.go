// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"NoticeServices/app/dao/internal"
)

// infoDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type infoDao struct {
	*internal.InfoDao
}

var (
	// Info is globally public accessible object for table info operations.
	Info = &infoDao{
		internal.Info,
	}
)
