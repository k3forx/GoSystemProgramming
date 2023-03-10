# GoSystemProgramming

## 低レベルアクセスへの入り口 2: `io.Reader`

### エンディアン

https://ja.wikipedia.org/wiki/%E3%82%A8%E3%83%B3%E3%83%87%E3%82%A3%E3%82%A2%E3%83%B3

> 十六進法で表現すると 1234ABCD という 1 ワードが 4 バイトのデータを、バイト毎に上位側（通常左側）から「12 34 AB CD」のように並べる順序はビッグエンディアン (英: big-endian)、下位側（通常右側）から「CD AB 34 12」のように並べる順序はリトルエンディアン (英: little-endian)

- x86, x86_64 はリトルエンディアン

### PNG の画像形式

https://qiita.com/spc_ehara/items/c748ec636283df805926#:~:text=png%E3%81%AB%E3%81%AF%E6%A7%98%E3%80%85%E3%81%AA,IDAT%E3%83%81%E3%83%A3%E3%83%B3%E3%82%AF%E3%80%81IEND%E3%83%81%E3%83%A3%E3%83%B3%E3%82%AF%E3%81%A7%E3%81%99%E3%80%82

### `io.Reader` / `io.Writer` でストリームを自由に操る

データの流れを自由に制御するために使える構造体

- `io.MultiReader`
- `io.TeeReader`
- `io.Pipe` (`io.PipeReader` と `io.PipeWriter`)

## 低レベルアクセスへの入り口 3: チャネル

### チャネル

- チェネルは、データを順序よく受け渡すためのデータ構造
- チャネルは、並列処理されても正しくデータを受け渡す同期機構
- チャネルは、読み込み・書き込みで準備ができるまでブロックする機能

#### チャネルの使用方法

- バッファなしのチャネルでは、受け取り側が受信をしないと送信側もブロックされる

## システムコール

##

## ファイルシステムの最新部を扱う Go 言語の関数

- POSIX 系の OS の場合、 `syscall.Flock()` というシステムコールが利用できる

### 同期・非同期

- OS のシステムコールにおいて整備するためのモデルとなるのが、**同期処理** と **非同期処理**、**ブロッキング処理** と **ノンブロッキング処理**
  - 同期処理: OS に I/O タスクを投げて、入出力の準備ができたらアプリケーションに処理が返ってくる
  - 非同期処理: OS に I/O タスクを投げて、入出力の準備ができたら通知をもらう
  - ブロッキング処理: お願いした I/O タスクの結果の準備が出来るまで待つ (自分は停止)
  - ノンブロッキング処理: お願いした I/O タスクの結果の準備ができるのを待たない (自分は停止しない)

## プロセスの役割と Go 言語による操作

### プロセスに含まれるもの (Go 言語視点)

- 実行ファイルパス
- プロセス ID
- プロセスグループ ID、セッショングループ ID
- ユーザー ID、グループ ID
- 実行ユーザー ID、実行グループ ID
- カレントフォルダ
- ファイルディスクリプタ

## 第 13 章 シグナルによるプロセス間の通信

- **プロセス間通信**: カーテルが仲介して、あるプロセスから、別のプロセスに対してシグナルを送ることができる。自分自身に対してシグナルを送ることも可能。
- **ソフトウェア割り込み**: システムで発生したイベントは、シグナルとしてプロセスに送られる。シグナルを受け取ったプロセスは、現在行っているタスクを中断して、あらかじめ登録しておいた登録ルーチンを実行する。

### 13.1 シグナルのライフサイクル

1.

### 13.2 シグナルの種類

```bash
# macOS/BSD
man signal
```

#### 13.2.1 ハンドルできないシグナル

強制力を持ち、アプリケーションではハンドルできないシグナル

- SIGKILL: プロセスを強制終了
- SIGSTOP: プロセスを一時停止して、バックグランドジョブにする

#### 13.2.2 サーバーアプリケーションでハンドルするシグナル

- SIGTERM: `kill()` システムコールや `kill` コマンドがデフォルトで送信するシグナルで、プロセスを終了するもの
- SIGHUP: 設定ファイルの再読み込みを外部から指示する用途で使われることがデファクトスタンダード
