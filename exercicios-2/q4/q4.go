package q4

import "fmt"

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {
	mapDestinationCities := make(map[string]int)
	mapOriginCities := make(map[string]int)

	for _, route := range caminhos {
		mapDestinationCities[route[1]]++
		mapOriginCities[route[0]]++
	}
	for city := range mapDestinationCities {
		if _, ok := mapOriginCities[city]; !ok {
			return city, nil
		}
	}
	return "", fmt.Errorf("zero paths")
}
