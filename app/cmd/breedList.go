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

type DogBreedListResponse struct {
	Message []string `json:"message"`
	Status  string   `json:"status"`
}

var breedList string

// breedListCmd represents the breedList command
var breedListCmd = &cobra.Command{
	Use:   "breed-list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// if len(args) == 0 {
		// 	return fmt.Errorf("犬種が指定されていません")
		// }

		// breedList := args[0]

		client := http.Client{
			Timeout: 30 * time.Second,
		}

		url := "https://dog.ceo/api/breeds/list"
		//url := fmt.Sprintf("https://dog.ceo/api/breeds/list", breedList)

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

		var data DogBreedListResponse
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			fmt.Printf("JSONのデコードに失敗しました: %v", err)
		}

		for _, breed := range data.Message {
			// if i < 10 {
			fmt.Printf("%+v\n", breed)
			//}
		}

		//datafs := fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n", data.Message[0], data.Message[1], data.Message[2], data.Message[3], data.Message[4])

		//fmt.Printf("選択した犬種: %+v\n", breedList)
		//fmt.Printf("%+v", datafs)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(breedListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// breedListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// breedListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
