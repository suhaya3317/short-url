# short-url

## 1. Create Firebase Project
https://console.firebase.google.com

## 2. Get Firebase API KEY
Firebase Project > Setting > Web API KEY

## 3. Set Dynamic links
https://firebase.google.com/docs/dynamic-links/

## 4. Create config.env
```
FIREBASE_KEY=***********************
DOMAIN=https://example.page.link
```

## 5. Build
```terminal
$ go build main.go
```

## 6. Create Short URL
```terminal
$ ./main -url URL
```

結構雑っす。
