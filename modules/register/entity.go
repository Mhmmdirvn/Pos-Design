package register

type Admin struct {
	ID       int    `gorm:"primarykey" json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Message string
	Data Admin
}
