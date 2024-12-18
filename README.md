Nginx Config Example for Reverse Proxy
```
    location /saman {
        proxy_pass https://sep.shaparak.ir/;
        proxy_set_header Host sep.shaparak.ir;
        proxy_set_header X-Forwarded-For "";
        proxy_set_header X-Forwarded-Proto "";
    }
```