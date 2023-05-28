package main

import (
	"go-project-12_05_23/internal/ports/conexao"

	"github.com/edgarrps/political-struct/internal/adapter/conexaoMockada"
)

func main() {

	var conexao conexao.Conexao

	conexao = conexaoMockada.ConexaoMockada{}
}
