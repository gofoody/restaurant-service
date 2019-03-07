package model

type Restaurant struct {
	ID   int
	Name string
	Menu *RestaurantMenu
}

type RestaurantMenu struct {
	MenuItems []*MenuItem
}

type MenuItem struct {
	ID    int
	Name  string
	Price float32
}
