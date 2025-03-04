package main

import (
	"fmt"
	"os"

	"github.com/Renaud42/bref/parsing"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// Vérifie si un fichier d'entrée est passé en argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file>")
		return
	}

	// Charge le fichier à analyser
	inputFile := os.Args[1]
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Erreur de lecture du fichier %s: %v\n", inputFile, err)
		return
	}

	// Crée un flux d'entrée ANTLR à partir du fichier
	is := antlr.NewInputStream(string(data))

	// Crée un lexer et un stream de tokens
	lexer := parsing.NewbrefLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Crée un parseur
	parser := parsing.NewbrefParser(stream)

	// Active l'analyseur pour commencer l'analyse de l'entrée
	tree := parser.Prog() // Remplace 'StartRule' par la règle de départ de ton grammaire

	// Affiche l'arbre de syntaxe abstraite (AST) pour débogage
	fmt.Println("Arbre syntaxique:")
	fmt.Println(tree.ToStringTree(nil, parser))

	// Optionnel : Utilise un écouteur personnalisé pour interpréter les résultats
	// Exemple simple d'un interpréteur basé sur un écouteur
	listener := &parsing.BasebrefListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	// Affiche une sortie personnalisée (selon la logique de ton interpréteur)
	fmt.Println("Interprétation terminée.")
}
