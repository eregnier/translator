package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func translateText(targetLanguage string, translations []string) ([]translate.Translation, error) {
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return nil, fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer func(client *translate.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalln("Error closing client")
		}
	}(client)

	response, err := client.Translate(ctx, translations, lang, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("no enough arguments : expecting <input_file> <output_file>")
		os.Exit(0)
	}
	translations, err := translateText("fr", readFile())
	if err != nil {
		log.Fatalln("Error while translating", err)
	}
	writeFile(translations)
}

func readFile() []string {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

func writeFile(translations []translate.Translation) {

	file, err := os.OpenFile(os.Args[2], os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	writer := bufio.NewWriter(file)

	for _, translation := range translations {
		_, _ = writer.WriteString(translation.Text + "\n")
	}
	fmt.Println(len(translations))
	err = writer.Flush()
	if err != nil {
		log.Fatalln(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatalln(err)

	}

}
