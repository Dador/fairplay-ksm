# FairPlay Key Service Management

Apple FairPlay Streaming ([FPS](https://developer.apple.com/streaming/fps/)) securely delivers keys to Apple mobile devices, Apple TV, and Safari on OS X, which will enable playback of encrypted video content.

Forked from [easonlin404/ksm](https://github.com/easonlin404/ksm) and [payt0nc/fairplay-ksm](https://github.com/payt0nc/fairplay-ksm).

## Usage

### Start using it

Create `.env` file in project folder:
```
FAIRPLAY_ACCESS_TOKEN= # Access token, any string, used in `Auth-Code` header
FAIRPLAY_CERTIFICATION= # Certificate, encoded in base64
FAIRPLAY_PRIVATE_KEY= # Private key in PKCS #1 format, encoded in base64
FAIRPLAY_APPLICATION_SERVICE_KEY= # Application secret key (ASK) in hex
```

Start container:

```bash
docker compose up --build
```

Use `https://localhost:8080/fps/license` as license server URL. In most cases you shouldn't use it directly, but check auth in your app and proxy request from client to that URL.

Headers:
```
Content-Type: application/json or application/x-www-form-urlencoded
key: (content key)
iv: (iv for content key)
Auth-Code: (auth code from .env)
```
Body (`application/json`):
```
{"spc": "(request from client)"}
```
Body (`application/x-www-form-urlencoded`):
```
spc=(request from client)
```

Response example:
```

{
    "ckc": "(response for client)"
}
```

You can use Kubernetes config in `k8s` as an example how to deploy it to Kubernetes.

### How to verifying Key Security Module (KSM) Implementation?

[https://developer.apple.com/library/archive/technotes/tn2454/_index.html#//apple_ref/doc/uid/DTS40017630-CH1-VERIFYING_KEY_SECURITY_MODULE__KSM__IMPLEMENTATION](https://developer.apple.com/library/archive/technotes/tn2454/_index.html#//apple_ref/doc/uid/DTS40017630-CH1-VERIFYING_KEY_SECURITY_MODULE__KSM__IMPLEMENTATION)

### How to Debugging FairPlay Streaming?

[https://developer.apple.com/library/archive/technotes/tn2454/_index.html](https://developer.apple.com/library/archive/technotes/tn2454/_index.html)
