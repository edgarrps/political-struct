package main

import (
	"fmt"

	"github.com/edgarrps/political-struct/internal/adapter/conexao"

	"github.com/edgarrps/political-struct/internal/adapter/conexaoMockada"
)

func main() {

	var conexao conexao.Conexao

	conexao = &conexaoMockada.ConexaoMockada{}

	fmt.Println(conexao.BuscaDeputado())
}
