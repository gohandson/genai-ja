# 生成AIでチャットを作ろう

## 解説資料

* TBA

## ハンズオンのやりかた

`skeleton`ディレクトリ以下に問題があり、3つのセクションに分けられています。さらにセクションはステップごとに分けられており、それぞれのステップを進めていくことで、Goで生成AIを作ったチャットアプリが作れるようになっています。

各セクションの各ステップにはREADMEが用意されていますので、まずは`README`を読みます。
`README`には、そのステップを理解するための解説が書かれています。

`README`を読んだら、ソースコードを開き`TODO`コメントが書かれている箇所をコメントに従って修正して行きます。
`TODO`コメントをすべて修正し終わったら、`README`に書かれた実行例に従ってプログラムをコンパイルして実行します。

途中でわからなくなった場合は、`solution`ディレクトリ以下に解答例を用意していますので、そちらをご覧ください。

`macOS`の動作結果をもとに解説しています。
`Windows`の方は、パスの区切り文字やコマンド等を適宜読み替えてください。

## 目次

### [Section 01: コマンドラインツール](./skeleton/section01)

* STEP01: Hello, 世界
* STEP02: オウム返し
* STEP03: ChatGPT APIを使おう
* STEP04: プロンプトを工夫しよう

### [Section 02: net/httpパッケージ](./skeleton/section02)

学ぶこと：コンポジット型（配列、スライス、マップ）、構造体、ユーザ定義型

* STEP01: 11連ガチャの結果を記録しよう

### [Section 03: gRPC (Connect)](./skeleton/section03)

学ぶこと：関数、ポインタ、メソッド

* STEP01: ガチャを行う関数を定義しよう

## ソースコードの取得

まず、事前に`gonew`コマンドをインストールしておきます。


```bash
$ go install golang.org/x/tools/cmd/gonew@latest
```

セクション3のSTEP2のソースコードを取得するには以下のようなコマンドを実行します。

```
$ mkdir -p solution3/step02
$ cd solution3/step02
$ go new github.com/gohandson/genai-ja/skeleton/solution03/step02/genaichat github.com/yourname/genai
```

`gonew`コマンドの詳細な使い方は、[ドキュメント](https://pkg.go.dev/golang.org/x/tools/cmd/gonew)をご覧ください。

## ライセンス

<a href="https://creativecommons.org/licenses/by-nc/4.0/legalcode.ja">
	<img width="200" src="by-nc.eu.png">
</a>
