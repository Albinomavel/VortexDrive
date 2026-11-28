package main

type Problem struct {
	Options string
}

var Catalog = map[string][]Problem{
	"🌐 Internet": {
		{Options: "Computador sem internet"},
		{Options: "Internet lenta"},
		{Options: "Conexão caindo toda hora"},
	},
	"🐢 Lentidão": {
		{Options: "Sistema lento ou travando"},
		{Options: "Computador demora pra ligar"},
		{Options: "Computador esquentando ou fazendo muito barulho"},
	},
	"🖨️ Impressora": {
		{Options: "Impressora não imprime"},
		{Options: "Fila travada"},
		{Options: "Impressora não encontrada"},
	},
}
