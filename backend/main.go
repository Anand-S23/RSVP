package main

import (
	"log"
	"net/http"
)

func main()  {
    env := LoadEnv()

    db := InitDB(env.DB_URI, env.PRODUCTION)
    // TODO: Create Admin
    // TODO: Load in People from file

    store := NewStore(db)
    controller := NewController(store, env.PRODUCTION)
    
    log.Println("Server running on port: ", env.PORT);
    baseRouter := NewRouter(controller)
    router := NewCorsRouter(baseRouter, env.FE_URI)
    http.ListenAndServe(":" + env.PORT, router)
}
