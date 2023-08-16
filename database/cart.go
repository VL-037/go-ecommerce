package database

import "errors"

var (
	ErrProductNotFound = errors.New("product not found")
	ErrDecodeProductsFailed = errors.New("product not found")
	ErrUserIdNotValid = errors.New("user not valid")
	ErrUpdateUserFailed = errors.New("update user failed")
	ErrRemoveItemCartFailed = errors.New("remove item cart failed")
	ErrGetCartItemFailed = errors.New("get cart item failed")
	ErrCheckoutFailed = errors.New("checkout failed")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func CartItemCheckOut() {

}

func InstantBuyer() {

}