# apache-go-reverse-proxy

## 構成

- Apache
  - Port80でアクセスしてきたのを、Goが乗っているコンテナの8080番ポートにリダイレクトさせる
- Go(Gin)
  - おみくじAPIを実装
  - 実装ネタは[こちら](https://docs.google.com/presentation/d/10c3HqNCouzBDqGow8Nth_883PLHMpHU_iZ_zMuhkA8o/edit#slide=id.g8ba34b2da5_0_1185)
