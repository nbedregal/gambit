package models

type SecretRDSJson struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	Engine               string `json:"engine"`
	Host                 string `json:"host"`
	Port                 int    `json:"port"`
	Dbname               string `json:"dbname"`
	DbInstanceIdentifier string `json:"dbInstanceIdentifier"`
}

type SignUp struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserUUID"`
}

type Category struct {
	CategID   int    `json:"categID"`
	CategName string `json:"categName"`
	CategPath string `json:"categPath"`
}

type Product struct {
	ProdId          int     `json:"prodID"`
	ProdTitle       string  `json:"prodTitle"`
	ProdDescription string  `json:"prodDescription"`
	ProdCreatedAt   string  `json:"prodCreatedAt"`
	ProdUpdated     string  `json:"prodUpdated"`
	ProdPrice       float64 `json:"prodPrice,omitempty"`
	ProdPath        string  `json:"prodPath"`
	ProdCategoryId  int     `json:"prodCategoryId"`
	ProdStock       int     `json:"prodStock"`
	ProdSearch      string  `json:"search,omitempty"`
	ProdCategPath   string  `json:"categPath,omitempty"`
}

type ProductResp struct {
	TotalItems int       `json:"totalItems"`
	Data       []Product `json:"data"`
}

type User struct {
	UserUUID      string `json:"userUUID"`
	UserEmail     string `json:"userEmail"`
	UserFirstName string `json:"userFirstName"`
	UserLastName  string `json:"userLastName"`
	UserStatus    int    `json:"userStatus"`
	UserDateAdd   string `json:"userDateAdd"`
	UserDateUpg   string `json:"userDateUpg"`
}

type ListUser struct {
	TotalItems int    `json:"totalItems"`
	Data       []User `json:"data"`
}

type Address struct {
	AddId         int    `json:"addId"`
	AddTitle      string `json:"addTitle"`
	AddName       string `json:"addName"`
	AddAddress    string `json:"addAddress"`
	AddCity       string `json:"addCity"`
	AddState      string `json:"addState"`
	AddPostalCode string `json:"addPostalCode"`
	AddPhone      string `json:"addPhone"`
}

type Orders struct {
	OrderId       int            `json:"orderId"`
	OrderUserUUID string         `json:"orderUserUUID"`
	OrderAddId    int            `json:"orderAddId"`
	OrderDate     string         `json:"orderDate"`
	OrderTotal    float64        `json:"orderTotal"`
	OrderDetails  []OrderDetails `json:"orderDetails"`
}

type OrderDetails struct {
	ODId       int     `json:"odId"`
	ODOrderId  int     `json:"odOrderId"`
	ODProdId   int     `json:"odProdId"`
	ODQuantity int     `json:"odQuantity"`
	ODPrice    float64 `json:"odPrice"`
}
