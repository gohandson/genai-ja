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

* STEP01: Hello, net/http
* STEP02: チャットボットAPI
* STEP03: JSONで返そう
* STEP04: html/templateを使おう

## ソースコードの取得

まず、事前に`gonew`コマンドをインストールしておきます。

```bash
$ go install golang.org/x/tools/cmd/gonew@latest
```

インストールされているか`which`コマンドで確かめましょう。
Windowsの場合は、`where`コマンドを使ってください。

```bash
$ which gonew
```

`gonew`コマンドのパスが表示されればOKです。

`gonew`コマンドが見つからない場合は、以下のように`PATH`を設定してください。
`.bashrc`や`.zshrc`などに記載すると良いでしょう。

```bash
$ export PATH=$PATH:$(go env GOPATH)/bin
```

`PATH`を設定した場合は、再度`which`コマンドを実行して確認してみましょう。

セクション1のSTEP01のソースコードを取得するには以下のようなコマンドを実行します。

```
$ mkdir -p section01/step01
$ cd section01/step01
$ gonew github.com/gohandson/genai-ja/skeleton/section03/step02/genaichat example.com/genaichat
```

`gonew`コマンドの詳細な使い方は、[ドキュメント](https://pkg.go.dev/golang.org/x/tools/cmd/gonew)をご覧ください。

## ライセンス

<a href="https://creativecommons.org/licenses/by-nc/4.0/legalcode.ja">
	<img width="200" src="by-nc.eu.png">
</a>
