# Terraform for Arukas

[Terraform](https://www.terraform.io)で[Arukas](https://arukas.io)を操作するためのTerraform providerプラグインです。

## 注意

当プラグインは[Terraform](https://terraform.io) v0.8.3にてTerraform本体に組み込まれました。  
今後のバグフィックスなどは[Terraform本体のリポジトリ](https://github.com/hashicorp/terraform)にて行われます。  
      
## クイックスタート

#### 前提条件

- Arukas APIキーを取得しておく

Dockerがない場合は、[Wiki:インストール](https://github.com/yamamoto-febc/terraform-provider-arukas/wiki/Install)を参考にバイナリファイルをインストールすることでTerraform for Arukasの利用が可能です。

Arukas APIキーの取得方法は[こちら](https://github.com/yamamoto-febc/terraform-provider-arukas/wiki/Install#arukas-apiキーの取得)を参照してください。

以下はArukas上にNginxコンテナを立ち上げる例です。

```bash
#################################################
# Terraform定義ファイル作成
#################################################
$ mkdir ~/work; cd ~/work #作業用ディレクトリ
$ tee arukas.tf <<-'EOF'

resource "arukas_container" "demo"{
    name = "arukas-quick-start"
    image = "nginx:latest"
    ports = {
        protocol = "tcp"
        number = "80"
    }
}

EOF

#################################################
# Terraformでインフラ作成
#################################################
$ docker run -it --rm \
         -e ARUKAS_JSON_API_TOKEN=[Arukas APIトークン] \
         -e ARUKAS_JSON_API_SECRET=[Arukas APIシークレット] \
         -v $PWD:/work \
         aquarium/terraform-arukas apply

#################################################
# 確認
#################################################
$ docker run -it --rm \
         -e ARUKAS_JSON_API_TOKEN=[Arukas APIトークン] \
         -e ARUKAS_JSON_API_SECRET=[Arukas APIシークレット] \
         -v $PWD:/work \
         aquarium/terraform-arukas show

#################################################
# 削除
#################################################
$ docker run -it --rm \
         -e ARUKAS_JSON_API_TOKEN=[Arukas APIトークン] \
         -e ARUKAS_JSON_API_SECRET=[Arukas APIシークレット] \
         -v $PWD:/work \
         aquarium/terraform-arukas destroy
```

## インストール

[リリースページ](https://github.com/yamamoto-febc/terraform-provider-arukas/releases/latest)から最新のバイナリを取得し、
Terraformバイナリと同じディレクトリに展開してください。

詳細は[Wiki:インストール](https://github.com/yamamoto-febc/terraform-provider-arukas/wiki/Install)を参照してください。

## 使い方/各リソースの設定方法

Terraform定義ファイル(tfファイル)を作成してご利用ください。
設定ファイルの記載方法は[Wikiページ](https://github.com/yamamoto-febc/terraform-provider-arukas/wiki)を参照ください。

## License

  This project is published under [Apache 2.0 License](LICENSE).

## Author

  * Kazumichi Yamamoto ([@yamamoto-febc](https://github.com/yamamoto-febc))
