package fundamentos

import (
	"fmt"
	"time"
)

func declaracao() {
	var nome string = "Hans Christian"
	email := "hans@email.com"
	dataCriacao := time.Now()
	const Ativo = true

	fmt.Println(nome, email, dataCriacao, Ativo)

	const (
		Juros       = 0.0167
		ParcelasMax = 36
	)

	fmt.Println(Juros, ParcelasMax)
}
