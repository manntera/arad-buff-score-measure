# どんなツール？
# やりたいこと
バフスコア測定ツールの作成。
身内用カニツールみたいな感じのを作りたい。
カニツールみたいにPCを常時放置するのは難しいので、必要な時にギルドメンバーが取り出して使える感じのツールを想定しています。
# 身内外利用について
本システムはオープンソースにし改変自由とする為、このツールを汎用的に使える様にしたい場合は自由にforkして貰って良い事にします。ただし、バックエンドに対してリクエストを投げつけられすぎると、多額のクラウド利用料金等がかかってしまう為、運用中のエンドポイントは非公開にします。APIサーバーのホスティング等の環境構築は各自で行ってください。
# どんなツールにするの？
1. 修練場にBとDの二人で入場し、D側がツールを起動する
2. D側は画面をフルスクリーンにする(解像度はFHDを想定)
3. BがバフをDにかけたら、ツールの測定開始ボタンを押下する
4. ツールからバフスコアが出力される。
# ツールの仕組み
1. 測定開始ボタンを押下すると、バフアイコンのある所を順番にマウスカーソルが自動的に動く様にする。
2. マウスカーソルがバフアイコンにオーバーした際に出てくる、バフの能力上昇値をスクリーンショットで撮影する
3. マウスカーソルはバフアイコンがあるであろう場所を全て選択し、全てのスクリーンショットを撮影する。
4. 撮影したスクリーンショットを、まんてらが用意したAPIサーバーにPOSTリクエストで投げる。
5. サーバー内で画像を解析して、バフスコアの計算を行う。
6. APIサーバーからバフスコアがレスポンスされるのでそれを表示若しくは出力する
# システム構成
## インフラ
- GooglePlatformを使用
- ホスティング
  - 基本はCloudRun
- 画像解析
  - GoogleVision
## サーバーサイド
 - golang
## フロントエンド
検討中
# ディレクトリ構成
検討中
