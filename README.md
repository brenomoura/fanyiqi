# fanyiqi

Still in Work In Progress, but already usable. It is missing the CI/CD pipeline, and also to set the binaries to a keyboard shortcut, but it is functional and ready to use.

**fanyiqi** is a lightweight and personal translation tool. It combines a graphical cross-platform desktop client (built in Go) with a backend API (built in Python with FastAPI) which integrates with several providers.

The goal is to provide a fast, customizable, and friendly translation experience for personal use, with the flexibility to run in home labs, cloud environments, or locally.

---

## Features

- Translate text between languages using offline-compatible models (CTranslate2 + Hugging Face)
- Lightweight desktop GUI using [Fyne](https://fyne.io/)
- Multiplatform (Linux, macOS, Windows)
- Backend powered by FastAPI
- Docker and Docker Compose support

---

## Running the Server

> Requires Python 3.12+ and [uv](https://github.com/astral-sh/uv) — or simply use Docker.

### Option 1: Run locally (Dev)

```bash
cd server/
make install       # installs dependencies into a virtual env
make run           # runs on http://localhost:8000
```

### Option 2: Run with Docker Compose

```bash
cd server/containerfile
docker-compose up --build
```

- First-time run will download lightweight models.
- SQLite database is stored at `server/data/app.db` (or mapped volume).
- Accessible at `http://localhost:8000`

---

## Running the Client (Desktop GUI)

### Option 1: Run from source

Requirements:
- Go <= 1.23.5
- `make`

```bash
cd client/
make run
```

### Option 2: Build binaries for target OS

```bash

make build-linux
make build-windows
make build-macos
```

Note: For now, the only tested binary is the Linux version, but you can build for Windows and macOS as well.

---

## Keyboard Shortcut
The main idea is to keep it simple and intuitive, so, with the generated binaries, you can set a global keyboard shortcut to open the app and translate text from anywhere.

---

## Supported Providers

- **Facebook M2M100 (418M)** (offline models) — [Hugging Face Model](https://huggingface.co/facebook/m2m100_418M)
- **CTranslate2 Fast M2M100 (418M)** (offline models) — [Hugging Face Model](https://huggingface.co/michaelfeil/ct2fast-m2m100_418M)
- **Google Translate**

Note: The models used are the lightest available, optimized for speed and size, so it is possible to provide some not desirable results. 

---

## 📄 License

MIT License — Free for personal or commercial use.

---

Just a personal project designed to learn, automate, and make life easier.

Feel free to contribute, open issues, or fork the project!