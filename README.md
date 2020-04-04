# Kagawa Checker

香川県からのアクセスかどうかをIPで判定するWebAPI

1. [ipinfodb のウェブサイト](https://lite.ip2location.com/database-ip-country-region-city) から `IPV4 BIN` をダウンロードし、解凍したディレクトリにある `IP2LOCATION-LITE-DB3.BIN` を `/tmp` に移動する。
2. `go get github.com/ip2location/ip2location-go`
3. `go run main.go`
