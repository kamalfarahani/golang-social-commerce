package constants

type Error string

func (e Error) Error() string { return string(e) }

const (
	PRODUCT_NAME_ERR = Error("Product Name not found")
	PRODUCT_ID_ERR   = Error("The product ID is wrong")

	CAT_END_PRODUCT_ERR = Error("Category has no more products")
	CAT_NO_PRODUCT_ERR  = Error("Category has no product")
	CAT_NAME_ERR        = Error("Category name is wrong")
	CAT_ID_ERR          = Error("Category ID is wrong")

	COL_NO_PRODUCT_ERR = Error("Collection has no product")
	COL_NAME_ERR       = Error("Collection name is wrong")
	COL_ID_ERR         = Error("Collection ID is wrong")

	NAME_PASSWORD_WRONG_ERR = Error("User name  or password is wrong")
	SHORT_PASSWORD_ERR      = Error("Password is too short")
	INVALID_EMAIL_ERR       = Error("Email is invalid")
	SHORT_NAME_ERR          = Error("Name is too short")
	USER_ID_ERR             = Error("User ID is wrong")
	NAME_ERR                = Error("User name is wrong")

	WRONG_CAPTCHA_ERR = Error("The captcha is wrong try again")
	WRONG_DATA_ERR    = Error("Wrong data")
	USER_NOT_LOGIN    = Error("The user did not login")
	END_PAGE_ERR      = Error("The page is bigger than page conut")
)
