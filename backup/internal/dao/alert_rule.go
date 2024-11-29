// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"login-demo/internal/dao/internal"
)

// internalAlertRuleDao is internal type for wrapping internal DAO implements.
type internalAlertRuleDao = *internal.AlertRuleDao

// alertRuleDao is the data access object for table alert_rule.
// You can define custom methods on it to extend its functionality as you wish.
type alertRuleDao struct {
	internalAlertRuleDao
}

var (
	// AlertRule is globally public accessible object for table alert_rule operations.
	AlertRule = alertRuleDao{
		internal.NewAlertRuleDao(),
	}
)

// Fill with you ideas below.