## Q1

## Q2

## Q3

`strings.Reader` を使って zip ファイルを作成する。ただし、以下のような `newfile.txt` という実際のファイルを作成するコードは禁止。

```golang
zipWriter := zip.NewWriter(file)
defer zipWriter.Close()

writer, err := zipWriter.Create("newfile.txt")
```
