package main

import (
	"errors"
	"net/http"

	"github.com/philjestin/ranked-talishar/internal/data"
)

func (app *application) showAllHeroesHandler(w http.ResponseWriter, r *http.Request) {
	heroes, err := app.models.Heroes.GetAllHeroes()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"heroes": heroes}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showHeroHandler(w http.ResponseWriter, r *http.Request) {
	id := app.readUuidParam(r)

	hero, err := app.models.Heroes.GetHeroById(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"hero": hero}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
