package main

import (
	"log"
	"net/http"
)

func main()  {
    env := LoadEnv()

    db := InitDB(env.DB_URI, env.PRODUCTION)

    store := NewStore(db)
    controller := NewController(store, env.PRODUCTION)

    populateDB(store)
    
    log.Println("Server running on port: ", env.PORT);
    baseRouter := NewRouter(controller)
    router := NewCorsRouter(baseRouter, env.FE_URI)
    http.ListenAndServe(":" + env.PORT, router)
}

func populateDB(store *Store) {
    names := []string{"Test", "Test2", "Test3", "Scarlet", "Irene", "Sukhmeet", "Shweta", "Prableen"}
    for _, name := range(names) {
        err := store.CreatePerson(NewPerson(name))
        log.Fatal(err)
    }
}
