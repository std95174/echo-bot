# 應聲蟲機器人

## 啟動

1. 設定環境變數(`.env`)
2. 產生 TLS 憑證
   ```sh
   openssl genrsa -out key.pem 2048 
    openssl req -new -x509 -key key.pem -out cert.pem -days 365
   ```
3. `docker-compose up -d`