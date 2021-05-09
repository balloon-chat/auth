# Auth

## 概要

OAuth認証、セッション管理、プロフィール取得を行うAPI

## 環境変数

| 環境変数名                     | 説明                                                                       | サンプル                                         |
| ------------------------------ | -------------------------------------------------------------------------- | ------------------------------------------------ |
| ENV                            | アプリケーションの実行モード(この値に対応する`.env`ファイルが読み込まれる) | `development`, `production`                      |
| GOOGLE_CLIENT_ID               | Google OAuth認証のクライアントID                                           |                                                  |
| GOOGLE_CLIENT_SECRET           | Google OAuth認証のクライアントシークレット                                 |                                                  |
| SESSION_SECRET_KEY             | セッションCookieを暗号化する鍵                                             | secret                                           |
| BASE_URL                       | APIを動作させるサーバーを指すURL                                           | `http://localhost:8080`                          |
| COOKIE_DOMAIN                  | CookieのDomain属性                                                         | localhost                                        |
| CLIENT_ENTRY_POINT             | クライアントのサーバーを指すURL                                            | `http://localhost:3000`                          |
| CLIENT_SIGN_IN_URL             | クライアントのサインインページのURL                                        | `http://localhost:3000/signin`                   |
| CLIENT_LOGIN_URL               | クライアントのログインページのURL                                          | `http://localhost:3000/login`                    |
| GOOGLE_APPLICATION_CREDENTIALS | 秘密鍵ファイルのパス(ローカルでデバッグ時に必須)                           | `/home/user/Downloads/service-account-file.json` |

### `.env`ファイルのテンプレート
```.env
BASE_URL=
COOKIE_DOMAIN=

CLIENT_ENTRY_POINT=
CLIENT_SIGN_IN_URL=
CLIENT_LOGIN_URL=
```
```.env
ENV=debug

GOOGLE_CLIENT_ID=
GOOGLE_CLIENT_SECRET=

GOOGLE_APPLICATION_CREDENTIALS=

SESSION_SECRET_KEY=
```
