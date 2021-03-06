// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"NoticeServices/app/dao/internal"
)

// appDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type appDao struct {
	*internal.AppDao
}

var (
	// App is globally public accessible object for table app operations.
	App = &appDao{
		internal.App,
	}
)

// Fill with you ideas below.
