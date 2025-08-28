package main

import "fmt"

func main() {
	books := []Book{
			{
				Title: "BookTitle",
				Author: "Tada",
				Price: 1000,
				Categories: []string{"Learn"},
			},
			{
				Title: "BookTitle2",
				Author: "Tada2",
				Price: 2000,
				Categories: []string{"Learn2"},
			},
		}

	totalPrice := Total(books)
	fmt.Println("合計", totalPrice, "円")

	authorBooksMap := ToMap(books)
	//fmt.Println(authorBooksMap)
	authorName := ""
	fmt.Println("著者: ", authorName, ", タイトル: ", authorBooksMap[authorName].Title)
}

// 構造体 Book を定義 ---

type Book struct {
	Title string
	Author string
	Price int
	Categories []string
}

// フィールド
// Title ... 文字列、本のタイトル
// Author ... 文字列、著者
// Price ... 整数、価格
// Categories ... 文字列のスライス、カテゴリ

// --------------------

// Bookの配列を引数とし、本の合計価格を戻り値とする関数 `Total` を実装する
func Total(books []Book) int {
	priceTotal := 0
	for _, v := range books {
		fmt.Println("add: ", v.Price, "円")

		priceTotal += v.Price
		fmt.Println("sum: ", priceTotal, "円")
	}
	return priceTotal
}

// 3. キーを「著者名 (Author)」、値を構造体 Book とするマップを戻り値とする関数 `ToMap` を実装する
func ToMap(books []Book) map[string]Book {
	m := make(map[string]Book)
	for _, v := range books {
		m[v.Author] = v

		//m[v.Author] = append(m[v.Author], v)
	}
	return m
}
