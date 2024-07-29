package domain

// User 用户实体（赞数不考虑mysql）
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}
