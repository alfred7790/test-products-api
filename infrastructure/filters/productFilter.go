package filters

type ProductFilter struct {
	ID       string  `form:"id"`
	Code     string  `form:"code"`
	Name     string  `form:"name"`
	MinPrice float64 `form:"minPrice"`
	MaxPrice float64 `form:"maxPrice"`
	Page     int     `form:"page"`
	Limit    int     `form:"limit"`
}
