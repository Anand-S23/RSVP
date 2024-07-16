package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func ErrMsg(message string) map[string]string {
    return map[string]string {"error": message}
}

type Controller struct {
    store      *Store
    production bool
}

func NewController(s *Store, production bool) *Controller {
    return &Controller {
        store: s,
        production: production,
    }
}

func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) error {
    return WriteJSON(w, http.StatusOK, "Pong")
}

func (c *Controller) GetStatus(w http.ResponseWriter, r *http.Request) error {
    id := r.PathValue("id")
    p, err := c.store.ReadPerson(id)
    if err != nil {
        log.Println("error getting person data from db")
        return WriteJSON(w, http.StatusBadRequest, ErrMsg("Internal server error, please try again later"))
    }

    return WriteJSON(w, http.StatusOK, p.Answered)
}

func (c *Controller) RSVP(w http.ResponseWriter, r *http.Request) error {
    var personData Person
    err := json.NewDecoder(r.Body).Decode(&personData)
    if err != nil {
        log.Println("error parsing register data")
        return WriteJSON(w, http.StatusBadRequest, ErrMsg("Could not parse data"))
    }

    p, err := c.store.ReadPerson(personData.ID)
    if err != nil {
        log.Println("error getting person data from db")
        return WriteJSON(w, http.StatusBadRequest, ErrMsg("Internal server error, please try again later"))
    }

    if p.Answered {
        log.Println("already has been answered")
        return WriteJSON(w, http.StatusBadRequest, ErrMsg("Already answered, cannot answer again"))
    }

    err = c.store.UpdatePerson(personData)
    if err != nil {
        log.Println("error updating rsvp details")
        return WriteJSON(w, http.StatusBadRequest, ErrMsg("Could update person"))
    }
    log.Println("updated rsvp details")

    successMsg := map[string]string {
        "message": "Update RSVP details",
    }
    return WriteJSON(w, http.StatusOK, successMsg)
}

func (c *Controller) GetAllResponses(w http.ResponseWriter, r *http.Request) error {
    people, err := c.store.ReadPeople()
    if err != nil {
        log.Println("error getting people from db")
        return WriteJSON(w, http.StatusBadRequest, ErrMsg("Internal server error, please try again later"))
    }

    successMsg := map[string][]Person {
        "reponses": people,
    }
    return WriteJSON(w, http.StatusOK, successMsg)
}

