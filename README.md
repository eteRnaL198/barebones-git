# Barebones Git

Barebones Git は Git の小規模クローンです。add と commit ができます。

## インストール

Release ページからバイナリをダウンロードしてください。

## 使い方の例

まず`./bbgit init`を実行してリポジトリを作成します。

```bash
$ ./bbgit init
```

すると、`.bbgit/` と `root/` が作成されます。
`root/` 内のファイルが追跡されるので、この中にファイルを追加してください。今回は`root/dir/`を作成し、`hello.txt`を追加します。

```bash
$ mkdir -p root/dir
$ echo "Hello, World" > root/dir/hello.txt
```

次に、`./bbgit add`を実行して`root/`内のファイルをステージングエリアに追加します。

```bash
$ ./bbgit add
```

このとき`.bbgit/objects/`に Blob オブジェクトが作成されます。Blob オブジェクトは、ファイルの内容を保存するオブジェクトです。この場合、ファイル名が`hello.txt`の SHA-1 ハッシュで、中身が`hello.txt`の内容の Blob オブジェクトが作成されます。このとき`dir/`のツリーオブジェクトはまだ生成されません。

```bash
$ ls .bbgit/objects/
# 4ab299c8ad6ed14f31923dd94f8b5f5cb89dfb54
$ cat .bbgit/objects/4ab299c8ad6ed14f31923dd94f8b5f5cb89dfb54
# blob
# Hello, World
```

また、`./bbgit/index`にハッシュとファイルのパスの対応が記録されます。

```bash
$ cat .bbgit/index
# 4ab299c8ad6ed14f31923dd94f8b5f5cb89dfb54 root/dir/hello.txt
```

次に、`./bbgit commit`を実行してコミットします。

```bash
$ ./bbgit commit
```

すると、`.bbgit/objects/`に コミットオブジェクトとルートのツリーオブジェクト、`dir`のツリーオブジェクトが作成されます。

```bash
$ ls .bbgit/objects/
# 46f200660ad2b3f5c30be6a8eaceb451815936df ← ルートのツリーオブジェクト
# 4ab299c8ad6ed14f31923dd94f8b5f5cb89dfb54 ← Blob オブジェクト
# bcd504671588f90e08a5e7e2cdd35ed67173d71a ← dirのツリーオブジェクト
# ca85ce5ffaab5de52ca3416d2f993be20f8ccb8f ← コミットオブジェクト
```

ルートのツリーオブジェクトには、`dir`のツリーオブジェクトのハッシュが記録されます。

```bash
$ cat .bbgit/objects/46f200660ad2b3f5c30be6a8eaceb451815936df
# tree bcd504671588f90e08a5e7e2cdd35ed67173d71a
```

dir のツリーオブジェクトには、`hello.txt`の Blob オブジェクトのハッシュが記録されます。

```bash
$ cat .bbgit/objects/bcd504671588f90e08a5e7e2cdd35ed67173d71a
# blob 4ab299c8ad6ed14f31923dd94f8b5f5cb89dfb54 hello.txt
```

コミットオブジェクトには、ルートのツリーオブジェクトのハッシュと親コミットのハッシュが記録されます。今回は初回のコミットなので親コミットはありません。

```bash
$ cat .bbgit/objects/ca85ce5ffaab5de52ca3416d2f993be20f8ccb8f
# tree 46f200660ad2b3f5c30be6a8eaceb451815936df
```

また、コミット時には`./bbgit/HEAD`に最新のコミットのハッシュが記録されます。

```bash
$ cat .bbgit/HEAD
# ca85ce5ffaab5de52ca3416d2f993be20f8ccb8f
```

最後に、`./bbgit log`を実行するとコミットの履歴を確認できます。

```bash
$ ./bbgit log
# Commit: ca85ce5ffaab5de52ca3416d2f993be20f8ccb8f
# Date: Thu Jun 6 10:11:48 2024 +0900
```

以上が一連の流れです。

以降では、さらにファイルを変更した場合の動作について簡単に説明します。

1. `echo "Goodbye, World" > root/goodbye.txt`を実行してファイルを追加します。
1. `./bbgit add`を実行すると、変更がステージングエリアに追加されます。
   - `hello.txt`の Blob オブジェクトは再利用され、`goodbye.txt`の Blob オブジェクトが作成されます。(`.bbgit/objects/08d2fd03dcaa4762b06309e147e77484bfea6a40`)
   - `.bbgit/index`には`goodbye.txt`のハッシュとパスが追加されます。
1. `./bbgit commit`を実行すると、ステージングエリアの変更がコミットされます。
   - 新たにコミットオブジェクト(`bd2fb4d3014fa124bc4a9542fa403a6378382b4e`)が作成され、`.bbgit/HEAD`の内容が更新されます。
   - コミットオブジェクトの parent には前回のコミットのハッシュが記録されます。
   - `dir/`には変更が無いため、`dir/`のツリーオブジェクトが再利用されます。
1. `./bbgit log`により、コミットの履歴が追加されたことを確認できます。
