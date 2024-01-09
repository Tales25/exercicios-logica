package q1

//Na loja de animais à venda, existem alguns tipos de ração disponíveis para compra, sendo eles:
//
//* Ração para cachorro
//* Ração para gato
//* Ração universal
//
//O estoque dessas rações está representado em um mapa, onde a chave é o nome da ração e o valor correspondente é a
//quantidade disponível em estoque.
//
//Polycarp possui 𝑥 cães e 𝑦 gatos. Gostaríamos de determinar se é possível para ele comprar comida suficiente para todos
//os seus animais na loja. Cada um dos seus cães e gatos deve receber um pacote de ração adequado para sua espécie.

func CanBuyFood(stock map[string]int, dogs, cats int) bool {
	remainingDogs := dogs - stock["dog"]
	remainingCats := cats - stock["cat"]

	if remainingDogs < 0 {
		remainingDogs = 0
	}
	if remainingCats < 0 {
		remainingCats = 0
	}

	remainingPets := remainingDogs + remainingCats
	remainingStock := stock["universal"] - remainingPets

	return remainingStock >= 0
}
