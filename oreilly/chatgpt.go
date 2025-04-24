package main

import (
	"context"
	"fmt"

	"github.com/PullRequestInc/go-gpt3"
)

func main() {

	apiKey := "sk-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" // Replace with your OpenAI API key

	ctx := context.Background()

	client := gpt3.NewClient(apiKey)

	request := gpt3.CompletionRequest{

		Prompt: []string{"How many coffees should I drink per day?"},
	}

	resp, err := client.Completion(ctx, request)

	if err != nil {

		fmt.Printf("%s\n", err)

	} else {

		fmt.Printf("Answer:\n %s\n", resp.Choices[0].Text)

	}

}
