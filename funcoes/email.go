package funcoes

import (
	"fmt"
	"strings"
)

func testeEmail() {
	email := "hans@email.com"
	usuario, dominio := partesDoEmail(email)

	fmt.Println("Usuario:", usuario)
	fmt.Println("Dominio:", dominio)
}

func partesDoEmail(email string) (usuario, dominio string) {
	partes := strings.Split(email, "@")
	usuario = partes[0]
	dominio = partes[1]

	return usuario, dominio
}
