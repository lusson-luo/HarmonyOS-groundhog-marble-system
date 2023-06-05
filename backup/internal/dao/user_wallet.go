// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"login-demo/internal/dao/internal"
)

// internalUserWalletDao is internal type for wrapping internal DAO implements.
type internalUserWalletDao = *internal.UserWalletDao

// userWalletDao is the data access object for table user_wallet.
// You can define custom methods on it to extend its functionality as you wish.
type userWalletDao struct {
	internalUserWalletDao
}

var (
	// UserWallet is globally public accessible object for table user_wallet operations.
	UserWallet = userWalletDao{
		internal.NewUserWalletDao(),
	}
)

// Fill with you ideas below.
