# restaurant-service

Run - start restaurant service at `localhost:8074`
```
go run main.go
```

Check status - returns restaurant service status
```
curl localhost:8074/api/status
```

Get restaurant - returns restaurant basic details (ID and Name) for given restaurantId
```
curl http://localhost:8074/api/restaurants/1
```

Create restaurant - creates new restaurant with given restaurant json and returns restaurantId
```
curl -X POST localhost:8074/api/restaurants -d '{"Name":"Taj Hotel","Menu":{"MenuItems":[{"Name":"Shahi Paneer","Price":500}]}}'
```

Formatted restaurant payload
```
{
  "Name": "Taj Hotel",
  "Menu": {
    "MenuItems": [
      {
        "Name": "Shahi Paneer",
        "Price": 500
      }
    ]
  }
}
```
