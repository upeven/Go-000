package service

import (
	"database/sql"
	"errors"
	"log"
)

// 获取用户支付金额
func GetUserPaidAmount(uid int64) (paidAmount int64, err error) {

	paid, err := dal.GetUserPaid(uid)
	if err != nil {
		unWrappedErr := errors.Unwrap(err)
		// 没有记录很常见，所以要排除开，只打印日志
		if errors.Is(unWrappedErr, sql.ErrNoRows) {
			log.Fatal(err)
			return 0, nil
		}
		// 其他错误直接返回
		return 0, unWrappedErr

	}
	return paid, nil

}
