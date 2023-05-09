package service

import "pokedex/src/domain/repository"

type PokemonService struct {
	pokemonRepository repository.PokemonRepository
}

type Pokemon struct {
	Name string `json:"name"`
}

func (service PokemonService) GetPokemon(pokemonId int) Pokemon {
	return Pokemon{
		Name: "Bulbasaur",
	}
}
