// package main
// import (
// 	"context"
// 	"fmt"
// 	"io"

// 	translate "cloud.google.com/go/translate/apiv3"
// 	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
// )

// // translateText translates input text and returns translated text.
// func translateText(w io.Writer, projectID string, sourceLang string, targetLang string, text string) error {
// 	// projectID := "my-project-id"
// 	// sourceLang := "en-US"
// 	// targetLang := "fr"
// 	// text := "Text you wish to translate"

// 	ctx := context.Background()
// 	client, err := translate.NewTranslationClient(ctx)
// 	if err != nil {
// 			return fmt.Errorf("NewTranslationClient: %v", err)
// 	}
// 	defer client.Close()

// 	req := &translatepb.TranslateTextRequest{
// 			Parent:             fmt.Sprintf("projects/%s/locations/global", "turing-poet-305022"),
// 			SourceLanguageCode: "en",
// 			TargetLanguageCode: "en",
// 			MimeType:           "text/plain", // Mime types: "text/plain", "text/html"
// 			Contents:           []string{text},
// 	}

// 	resp, err := client.TranslateText(ctx, req)
// 	if err != nil {
// 			return fmt.Errorf("TranslateText: %v", err)
// 	}

// 	// Display the translation for each input text provided
// 	for _, translation := range resp.GetTranslations() {
// 			fmt.Fprintf(w, "Translated text: %v\n", translation.GetTranslatedText())
// 	}

// 	return nil
// }