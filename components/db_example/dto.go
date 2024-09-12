package dbexample

type UserCreate struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

type UserGet struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}
