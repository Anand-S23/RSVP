package main

import (
	"log"
	"net/http"

	"github.com/Anand-S23/rsvp/backend/app"
)

func main()  {
    env := app.LoadEnv()

    db := app.InitDB(env.DB_URI, env.PRODUCTION)

    store := app.NewStore(db)
    controller := app.NewController(store, env.PRODUCTION)

    populateDB(store)
    
    log.Println("Server running on port: ", env.PORT);
    baseRouter := app.NewRouter(controller)
    router := app.NewCorsRouter(baseRouter, env.FE_URI)
    http.ListenAndServe(":" + env.PORT, router)
}

func populateDB(store *app.Store) {
    names := []string{"Test", "Test2", "Test3", "Scarlet", "Irene", "Sukhmeet", "Shweta", "Prableen"}

    log.Println("populating db...")
    for _, name := range(names) {
        log.Println(name)
        err := store.CreatePerson(app.NewPerson(name))
        if err != nil {
            log.Fatal(err)
        }
    }

    log.Println("completed")
}
