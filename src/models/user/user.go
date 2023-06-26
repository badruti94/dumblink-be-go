package user

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt string `json:"updatedAt" gorm:"column:updatedAt"`
}
