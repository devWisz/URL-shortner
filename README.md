# URL-shortner

# 🔗 URL Shortener in Go

A simple and efficient URL shortener built with Go. This project allows you to shorten long URLs and get quick redirects via short codes. Perfect for learning Go web development, URL routing, and working with hashes and maps.

---

## 🚀 Features

- 🔒 Generates short, unique URLs using hashing
- ↩️ Redirects to the original URL when a short code is visited
- 📦 Lightweight and fast — runs on `localhost:3000`
- 💻 Minimal and clean Go code

---

## 🛠️ Built With

- [Go](https://golang.org/) – Backend and HTTP Server
- Standard Library only – No third-party dependencies

---


---

## ⚙️ How It Works

1. A long URL is submitted via `/shorten?url=https://example.com`
2. The server returns a short path like `/redirect/a1b2c3d`
3. Visiting `/redirect/a1b2c3d` will redirect you to the original long URL

---

## 📦 Installation

### Prerequisites

- Go installed (v1.16 or later)

### Clone the repository

```bash
git clone https://github.com/your-username/url-shortener.git
cd url-shortener



