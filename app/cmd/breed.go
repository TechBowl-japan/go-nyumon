/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

type DogBreedResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

var breed string

// breedCmd represents the breed command
var breedCmd = &cobra.Command{
	Use:   "breed",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			return fmt.Errorf("犬種が指定されていません")
		}

		breed := args[0]

		client := http.Client{
			Timeout: 30 * time.Second,
		}

		url := fmt.Sprintf("https://dog.ceo/api/breed/%s/images/random", breed)

		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
		if err != nil {
			return fmt.Errorf("リクエストの作成に失敗しました: %v", err)
		}

		req.Header.Add("X-My-Client", "Go-Learn")
		res, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("リクエストの作成に失敗しました: %v", err)
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("無効な犬種です")
		}

		// TechTrainのテストの挙動を鑑みてコメントアウト
		// fmt.Println(res.Header.Get("Content-Type"))

		var data DogBreedResponse
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			fmt.Printf("JSONのデコードに失敗しました: %v", err)
		}

		//fmt.Printf("選択した犬種: %+v\n", breed)
		fmt.Println(data.Message)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(breedCmd)

	//breedCmd.Flags().StringVarP(&breed, "select", "s", "", "犬種を指定する")
	//rootCmd.MarkFlagRequired("select")
}
