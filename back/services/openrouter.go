package services

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
    "os"
	"ia-api/models"
)

func CallOpenRouter(messages []models.ChatMessage, model string) (string, error) {
    if model == "" {
        model = os.Getenv("DEFAULT_MODEL")
        if model == "" {
            model = "openrouter/horizon-beta"
        }
    }

    apiKey := os.Getenv("OPENROUTER_API_KEY")
    if apiKey == "" {
        return "", errors.New("clé API manquante")
    }

    payload := map[string]interface{}{
        "model":    model,
        "messages": messages,
    }

    bodyJSON, _ := json.Marshal(payload)

    req, _ := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(bodyJSON))
    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("HTTP-Referer", "http://localhost:8080")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    respBody, _ := io.ReadAll(resp.Body)

    var result struct {
        Choices []struct {
            Message struct {
                Content string `json:"content"`
            } `json:"message"`
        } `json:"choices"`
    }

    if err := json.Unmarshal(respBody, &result); err != nil {
        return "", err
    }

    if len(result.Choices) == 0 {
        return "", fmt.Errorf("réponse vide : %s", string(respBody))
    }

    return result.Choices[0].Message.Content, nil
}

