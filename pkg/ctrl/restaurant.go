package ctrl

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gofoody/common/generator"
	"github.com/gofoody/restaurant-service/pkg/model"
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
)

type RestaurantCtrl interface {
	BaseCtrl
	Show(rw http.ResponseWriter, r *http.Request)
	Create(rw http.ResponseWriter, r *http.Request)
}

type restaurantCtrl struct {
	restaurants map[int]*model.Restaurant
	nextID      func() int
}

type restaurantResponse struct {
	id   int    `json:"id"`
	name string `json:"name"`
}

func NewRestaurantCtrl() RestaurantCtrl {
	c := &restaurantCtrl{}
	c.init()
	return c
}

func (c *restaurantCtrl) Name() string {
	return "restaurant controller"
}

func (c *restaurantCtrl) Show(rw http.ResponseWriter, r *http.Request) {
	log.Debugf("restaurant.Show(), url=%v", r.URL)

	id := mux.Vars(r)["restaurantId"]
	restaurant := c.restaurants[cast.ToInt(id)]
	if restaurant == nil {
		log.Debugf("restaurant with id=%s not found", id)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	restaurantRes := convert(restaurant)
	payload, err := json.Marshal(restaurantRes)
	if err != nil {
		log.Errorf("restaurant to json failed, error:%v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Write([]byte(payload))
}

func (c *restaurantCtrl) Create(rw http.ResponseWriter, r *http.Request) {
	log.Debugf("restaurant.Create()")

	decoder := json.NewDecoder(r.Body)
	var restaurant model.Restaurant
	err := decoder.Decode(&restaurant)
	if err != nil {
		log.Errorf("json to restaurant failed, error:%v", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	restaurant.ID = c.nextID()
	c.restaurants[restaurant.ID] = &restaurant
	for i, menuItem := range restaurant.Menu.MenuItems {
		menuItem.ID = generator.FactorID(restaurant.ID, (i + 1))
	}

	rw.Write([]byte(cast.ToString(restaurant.ID)))
}

func (c *restaurantCtrl) init() {
	c.nextID = generator.SeqID()
	c.restaurants = make(map[int]*model.Restaurant)

	id := c.nextID()
	restaurant := &model.Restaurant{
		ID:   id,
		Name: "Meghana Foods",
		Menu: &model.RestaurantMenu{
			MenuItems: []*model.MenuItem{
				&model.MenuItem{
					ID:    generator.FactorID(id, 1),
					Name:  "Chicken Biryani",
					Price: 200.0,
				},
				&model.MenuItem{
					ID:    generator.FactorID(id, 2),
					Name:  "Egg Masala",
					Price: 100.0,
				},
			},
		},
	}
	c.restaurants[restaurant.ID] = restaurant

	id = c.nextID()
	restaurant = &model.Restaurant{
		ID:   id,
		Name: "Dominos Pizza",
		Menu: &model.RestaurantMenu{
			MenuItems: []*model.MenuItem{
				&model.MenuItem{
					ID:    generator.FactorID(id, 1),
					Name:  "Chicken Feasta",
					Price: 300.0,
				},
				&model.MenuItem{
					ID:    generator.FactorID(id, 2),
					Name:  "Garlic Bread",
					Price: 150.0,
				},
			},
		},
	}
	c.restaurants[restaurant.ID] = restaurant
}

func convert(from *model.Restaurant) *restaurantResponse {
	to := &restaurantResponse{
		id:   from.ID,
		name: from.Name,
	}
	return to
}
