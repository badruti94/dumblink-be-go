package link

type Link struct {
	Id          int    `json:"id" form:"id"`
	Uniqid      string `json:"uniqid" form:"uniqid"`
	Type        string `json:"type" form:"type"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Vlog        string `json:"vlog" form:"vlog"`
	Galery      string `json:"galery" form:"galery"`
	Contact     string `json:"contact" form:"contact"`
	About       string `json:"about" form:"about"`
	Facebook    string `json:"facebook" form:"facebook"`
	Instagram   string `json:"instagram" form:"instagram"`
	Twitter     string `json:"twitter" form:"twitter"`
	Youtube     string `json:"youtube" form:"youtube"`
	Whatsapp    string `json:"whatsapp" form:"whatsapp"`
	Phone       string `json:"phone" form:"phone"`
	Web         string `json:"web" form:"web"`
	Lazada      string `json:"lazada" form:"lazada"`
	Photo       string `json:"photo" form:"photo"`
	View        int    `json:"view" form:"view"`
	UserId      int    `json:"userId" gorm:"column:userId"`
	CreatedAt   string `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt   string `json:"updatedAt" gorm:"column:updatedAt"`
}
