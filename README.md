# 🛠️ yamake – Build Ninja from YAML

**yamake** is a minimal tool that lets you configure C/C++ projects via YAML – and automatically generate `ninja` build files. No more bloated CMake for simple projects – just write `build.yaml` and go.

---

## 🚀 Features

- 🧾 Configure via `build.yaml`
- ⚙️ Automatically generates Ninja build files
- 📦 Supports `static`, `shared`, and `executable` targets
- ✨ Optional: `compile_commands.json` for IDEs / clangd

---

## 📁 Project Structure

```bash
yamake/
├── Makefile
├── src             # go source files
│   ├── compiledb.go
│   ├── config.go
│   ├── main.go
│   └── ninja.go
└── test_project/   # sample project with build.yml and C/C++ files
```

---

## 🛠 Build & Run

Use the provided `Makefile`:

```bash
make build       # Build the binary and move it to test_project/
make run         # Run the binary in test_project/ and generate ninja.build and compile_commands.json
make run_ninja   # ... with --build flag to invoke ninja and compile project
```

Alternatively via Go directly:

```bash
go build -o yamake ./src
cd test_project
./yamake --build
```

---

## 🧾 Example: `build.yaml`

```yaml
project: myapp
type: executable
sources:
  - main.cpp
  - util.cpp
compiler: g++
cflags: -O2
include_dirs:
  - include
libraries:
  - m
```

---

## 💡 Roadmap / Ideas

- [ ] `yamake install` to copy libs/headers
- [ ] Watch mode (auto rebuild on file changes)
- [ ] Support for header-only libraries
- [ ] Multi target projects

---

## ⚡️ Why?

Because CMake can be overkill – and `ninja` is just fast.  
`yamake` gives you **simple YAML config** and full control – with tooling you already like.

---

## 🧑‍💻 License

MIT

---

