package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"

	"gorm.io/gorm"
)

type CouponAdapter struct {
	DB *gorm.DB
}

func NewCouponAdapter(db *gorm.DB) repositories.Coupons {
	return &CouponAdapter{DB: db}
}

func (coupon *CouponAdapter) AddCoupon(req helperstructs.CouponReq) error {

	fmt.Println(req.Description)

	query := `INSERT INTO coupons (off,code,give_on_purchase_above,apply_on_purchase_above,is_welcome,description) VALUES ($1,$2,$3,$4,$5,$6)`

	return coupon.DB.Exec(query, req.OFF, req.Code, req.GiveOnPurchaseAbove, req.ApplyOnPurchaseAbove, req.IsWelcome, req.Description).Error

}

func (coupon *CouponAdapter) GetAllCouponsByEmail(email string) ([]responce.CouponData, error) {

	var coupondta []responce.CouponData

	query := `SELECT * FROM coupons WHERE id IN (SELECT coupon FROM coupon_items WHERE email = $1)`

	return coupondta, coupon.DB.Raw(query, email).Scan(&coupondta).Error

}

func (coupon *CouponAdapter) GetAllCoupons() ([]responce.CouponData, error) {

	var coupondta []responce.CouponData

	query := `SELECT * FROM coupons`

	return coupondta, coupon.DB.Raw(query).Scan(&coupondta).Error

}

func (coupon *CouponAdapter) GetCouponByCode(code int) (responce.CouponData, error) {

	var coupondta responce.CouponData

	query := `SELECT * FROM coupons WHERE code = $1`

	return coupondta, coupon.DB.Raw(query, code).Scan(&coupondta).Error

}

func (coupon *CouponAdapter) RemoveCouponFromUser(id uint, email string) error {

	query := `DELETE FROM coupon_items WHERE email = $1 AND coupon = $2`

	return coupon.DB.Exec(query, email, id).Error

}

func (coupon *CouponAdapter) UpdateCoupon(req helperstructs.CouponReq) error {

	query := `UPDATE coupons SET off = $1,give_on_purchase_above = $2,apply_on_purchase_above = $3,description = $4 WHERE id = $5`

	return coupon.DB.Exec(query, req.OFF, req.GiveOnPurchaseAbove, req.ApplyOnPurchaseAbove, req.Description, req.ID).Error

}

func (coupon *CouponAdapter) DeleteCoupon(id uint) error {

	fmt.Println(id)

	query := `DELETE FROM coupons WHERE id = $1`

	return coupon.DB.Exec(query, id).Error

}

func (coupon *CouponAdapter) ListofCouponsAvailableForThisOrder(price int) ([]uint, error) {

	var ids []uint

	query := `SELECT id FROM coupons WHERE give_on_purchase_above < $1 AND is_welcome = false`

	return ids, coupon.DB.Raw(query, price).Scan(&ids).Error

}

func (coupon *CouponAdapter) CreditUserWithCoupon(email string, id uint) error {

	query := `INSERT INTO coupon_items (email,coupon) VALUES($1,$2)`

	return coupon.DB.Exec(query, email, id).Error

}

func (coupon *CouponAdapter) GetAllWelcomeCoupons() ([]uint, error) {

	var coupons []uint

	query := `SELECT id FROM coupons WHERE is_welcome = true`

	return coupons, coupon.DB.Raw(query).Scan(&coupons).Error

}
