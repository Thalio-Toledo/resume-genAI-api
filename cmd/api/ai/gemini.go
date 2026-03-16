package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/genai"
)

func Generate(prompt string) ([]string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-3-flash-preview",
		genai.Text(
			"Voce é um assistente especialista descrição de vagas.\n"+
				"Retorne SOMENTE um array JSON de strings de strings com as habilidades necessarias para a descrição dessa vaga.\n\n"+
				"Exemplo: [\"Go\", \"Docker\"]\n\n"+
				prompt,
		),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Text())
	var skills []string
	err = json.Unmarshal([]byte(result.Text()), &skills)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return skills, nil
}

func GenerateEmbedding(prompt string) ([]float32, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	contents := []*genai.Content{
		genai.NewContentFromText(prompt, genai.RoleUser),
	}
	result, err := client.Models.EmbedContent(ctx,
		"gemini-embedding-001",
		contents,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	// embeddings, err := json.MarshalIndent(result.Embeddings, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(embeddings))

	return result.Embeddings[0].Values, nil
}
