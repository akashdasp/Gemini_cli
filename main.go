package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

func readEnvFile(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist.")
			return make(map[string]string), nil
		}
		return nil, err
	}
	defer file.Close()

	env := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			env[parts[0]] = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return env, nil
}
func writeEnvFile(filename string, env map[string]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for key, value := range env {
		_, err := fmt.Fprintf(writer, "%s=%s\n", key, value)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
func main() {
	var style_1 = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingLeft(4).
		Width(55)
	var style_2 = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#1a4480")).
		Background(lipgloss.Color("#face00")).
		PaddingLeft(4).
		Width(55)
	env, err := readEnvFile(".env")
	if err != nil {
		fmt.Println("Error reading .env file:", err)
		return
	}
	apiKey := env["Google_Api"]
	apiKeyPtr := flag.String("api_key", apiKey, "Google API key")
	imagePathPtr := flag.String("image_path", "", "Path to the image file")
	prompt := flag.String("prompt", "Descibe the image with Detiled Manner", "Write the Prompt over here")

	flag.Parse()
	if *apiKeyPtr != "" && *apiKeyPtr != apiKey {
		fmt.Printf("New API key provided. Do you want to save it as default? (yes/no): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
		input := strings.ToLower(scanner.Text())
		if input == "yes" || input == "y" {
			env["Google_Api"] = fmt.Sprintf(`%s`, *apiKeyPtr)
			if err := writeEnvFile(".env", env); err != nil {
				fmt.Println("Error writing .env file:", err)
				return
			}
			fmt.Println("Default API key updated.")
		}

	}
	if *prompt != "" && *prompt != "Descibe the image with Detiled Manner" && *imagePathPtr == "" {
		ctx := context.Background()
		llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
		if err != nil {
			log.Fatal(err)
		}

		answer, err := llms.GenerateFromSinglePrompt(ctx, llm, *prompt)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(style_2.Render(answer))
		return
	}
	if *apiKeyPtr == "" || *imagePathPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *apiKeyPtr != apiKey {
		fmt.Printf("New API key provided. Do you want to save it as default? (yes/no): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
		input := strings.ToLower(scanner.Text())
		if input == "yes" || input == "y" {
			env["Google_Api"] = fmt.Sprintf(`%s`, *apiKeyPtr)
			if err := writeEnvFile(".env", env); err != nil {
				fmt.Println("Error writing .env file:", err)
				return
			}
			fmt.Println("Default API key updated.")
		}
	}

	ctx := context.Background()
	llm, err := googleai.New(ctx, googleai.WithAPIKey(*apiKeyPtr))
	if err != nil {
		log.Fatal(err)
	}

	imgData, err := os.ReadFile(*imagePathPtr)
	if err != nil {
		log.Fatal(err)
	}

	parts := []llms.ContentPart{
		llms.BinaryPart("image/png", imgData),
		llms.TextPart(*prompt),
	}

	content := []llms.MessageContent{
		{
			Role:  llms.ChatMessageTypeHuman,
			Parts: parts,
		},
	}
	fmt.Println("")
	resp, err := llm.GenerateContent(ctx, content, llms.WithModel("gemini-pro-vision"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(style_1.Render(string(resp.Choices[0].Content)))
}
