package repository

type PokemonRepository interface {
	Get(pokemonId int)
}
