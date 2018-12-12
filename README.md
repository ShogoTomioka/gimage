# go-STR

間違い探しをやらせるGo製のコードです。
使用した画像は[ちびむすドリル](https://happylilac.net/machigai-h.html)から引用しています
現在、二種類の間違い画像（A,B)が置いてあります。makeコマンドの"TYPE"で”A”か”B”を指定してね。デフォルトはA

## usage

```make TYPE=A run```

間違い探し用の画像
<div>
<img src="https://raw.githubusercontent.com/ShogoTomioka/go-image-diff/master/testdata/picture_A.png" width="300">
<img src="https://raw.githubusercontent.com/ShogoTomioka/go-image-diff/master/testdata/picture_B.png" width="300">
</div>
左：二値画像（処理後）、右：結果画像
<div>
<img src="https://raw.githubusercontent.com/ShogoTomioka/go-image-diff/master/testdata/binary.png" width="300">
<img src="https://raw.githubusercontent.com/ShogoTomioka/go-image-diff/master/testdata/filtered.png" width="300">
</div>
