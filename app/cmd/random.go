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

type Data struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// type DogResponse struct {
// 	Message string `json:"message"`
// 	Status  string `json:"status"`
// }

var images int

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("random called")

		// 本にあったコードを参考
		client := http.Client{
			Timeout: 30 * time.Second,
		}

		for i := range make([]struct{}, images) {
			url := "https://dog.ceo/api/breeds/image/random"
			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
			if err != nil {
				fmt.Printf("リクエスト%d回目", i)
				fmt.Printf("リクエストの作成に失敗しました: %v", err)
			}

			res, err := client.Do(req)
			if err != nil {
				fmt.Printf("リクエストの作成に失敗しました: %v", err)
			}

			if res.StatusCode != http.StatusOK {
				fmt.Printf("APIからエラーステータスが返されました: %s", res.Status)
			}
			// TechTrainのテストの挙動を鑑みてコメントアウト
			// fmt.Println(res.Header.Get("Content-Type"))

			var data Data
			err = json.NewDecoder(res.Body).Decode(&data)
			if err != nil {
				fmt.Printf("JSONのデコードに失敗しました: %v", err)
			}

			res.Body.Close()
			fmt.Printf("%+v\n", data.Message)

		}

		// 自分で書いたコード
		// var images int
		// //a := make([]struct{}, images)
		// for i := range make([]struct{}, images) {
		// 	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
		// 	if err != nil {
		// 		// エラーハンドリング
		// 		fmt.Printf("リクエスト%d回目", i)
		// 		fmt.Printf("リクエストエラー:%v", err)
		// 	}

		// 	// defer resp.Body.Close()
		// 	/*
		// 		body, err := io.ReadAll(resp.Body)
		// 		if err != nil {
		// 			// エラーハンドリング
		// 		}
		// 	*/

		// 	var dogResponse DogResponse
		// 	dec := json.NewDecoder(resp.Body)
		// 	if err := dec.Decode(&dogResponse); err != nil {
		// 		// エラーハンドリング
		// 	}

		// 	resp.Body.Close()

		// 	fmt.Println(dogResponse.Message)
		// }

	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	randomCmd.Flags().IntVarP(&images, "images", "i", 1, "取得する画像の数を指定する")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
