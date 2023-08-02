### vercel周りの設定

**vercel.json**
- `builds`
    - srcでserverless functionsで使いたいファイル，useでビルドするモジュールを指定する
    - https://vercel.com/docs/concepts/projects/project-configuration#builds

- `route`
    - 基本的にはsrcでパス，destでファイルを指定する
    - パスには正規表現が使える
    - https://vercel.com/docs/concepts/projects/project-configuration#routes