package main

import (
	"fmt"
	"math"
)

func main() {
	var maxUint uint8 = math.MaxUint8
	var maxInt int8 = int8(maxUint >> 1)
	calInt := math.Cos(float64(maxInt))
	calUint := math.Cos(float64(maxUint))

	var pino int
	var gino int

	mappa := map[string]string{
		"alfio":      "giannetti",
		"paolino":    "marachella",
		"gianfranco": "deonli",
	}
	pino, gino = 8, 9

	fmt.Println("massimo valore int8: ", maxInt)
	fmt.Println("massimo valore uint8: ", maxUint)
	fmt.Printf("Il coseno di %d é:\t%.2f\n", maxInt, calInt)
	fmt.Printf("Il coseno di %d é:\t%.2f\n", maxUint, calUint)
	fmt.Printf("Il coseno di %d é:\t%.2f (calcolato nella Printf)\n", maxInt, math.Cos(float64(maxInt)))
	fmt.Printf("Ecco i valori di gino e pino:\t%v\t%v\n", gino, pino)
	// Questa si chiama COMMA OK NOTATION ed assegna il valore
	// a nome solo se la chiave nella mappa esiste
	// Se non esiste, l'espressione viene valutata falsa
	if nome, ok := mappa["paolino"]; ok {
		fmt.Printf("Abbiamo un vincitore: %v!\n", nome)
	} else {
		fmt.Println("Booo!!! Perdente!!!")
	}
}
