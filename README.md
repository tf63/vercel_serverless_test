### Vercel Serverless Functionsを使ってみる

### 対応言語

(参考)
- https://vercel.com/docs/concepts/functions/serverless-functions/supported-languages

**Nextjs**
- API Routeがデフォルトでserverless functionsとしてデプロイされる

**Nodejs**
- `api/*.ts`で`export default`された関数がserverless functionsとしてデプロイされる
- vercelの型情報は次でインストールできる
```
    npm i -D @vercel/node
```

```
    import type { VercelRequest, VercelResponse } from '@vercel/node';
    
    export default (request: VercelRequest, response: VercelResponse) => {
        const { name } = request.query;
        response.status(200).send(`Hello ${name}!`);
    };
```

**Go**
- `api/*.go`で`net/http`にマッチする関数がserverless functionsとしてデプロイされる
- デフォルトでは`Go 1.16`
```
    package handler
    
    import (
        "fmt"
        "net/http"
        "time"
    )
    
    func Handler(w http.ResponseWriter, r *http.Request) {
        currentTime := time.Now().Format(time.RFC850)
        fmt.Fprintf(w, currentTime)
    }
```

**Python**
- `api/*.py`で`BaseHTTPRequestHandler`を継承したクラス，または`wsgi` or `asgi`アプリケーション
- デフォルトでは`python3.9`らしい
```
from http.server import BaseHTTPRequestHandler
from datetime import datetime
 
class handler(BaseHTTPRequestHandler):
 
  def do_GET(self):
    self.send_response(200)
    self.send_header('Content-type', 'text/plain')
    self.end_headers()
    self.wfile.write(str(datetime.now().strftime('%Y-%m-%d %H:%M:%S')).encode())
    return
```

**Ruby**
- 知らない
