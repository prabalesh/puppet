# 🧩 Puppet – Code Execution and Language Management API

<div align="center">
  
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.22+-00ADD8.svg)](https://golang.org)
[![React](https://img.shields.io/badge/react-18+-61DAFB.svg)](https://reactjs.org)
[![Docker](https://img.shields.io/badge/docker-20+-2496ED.svg)](https://docker.com)

*A powerful full-stack application for managing programming languages and executing code in isolated Docker environments*

</div>

---

## ✨ Overview

Puppet is a comprehensive code execution platform that provides:

🔧 **Language Management** – Add, remove, and manage programming languages with Docker environments  
💻 **Code Execution** – Run code snippets safely in isolated Docker containers  
📥 **Stdin Support** – Execute code with custom input streams  
🌐 **RESTful API** – Clean Go-based backend with comprehensive endpoints  
⚛️ **Modern Frontend** – Intuitive React interface built with Vite  

---

## 🚀 Tech Stack

<table align="center">
  <tr>
    <td><strong>Layer</strong></td>
    <td><strong>Technology</strong></td>
    <td><strong>Purpose</strong></td>
  </tr>
  <tr>
    <td>🔙 Backend</td>
    <td>Go (net/http)</td>
    <td>API server & business logic</td>
  </tr>
  <tr>
    <td>🎨 Frontend</td>
    <td>React + Vite</td>
    <td>User interface & interactions</td>
  </tr>
  <tr>
    <td>🐳 Container</td>
    <td>Docker</td>
    <td>Code execution isolation</td>
  </tr>
  <tr>
    <td>🗄️ Database</td>
    <td>PostgreSQL</td>
    <td>Data persistence</td>
  </tr>
  <tr>
    <td>🚀 Deployment</td>
    <td>Docker Compose</td>
    <td>Infrastructure orchestration</td>
  </tr>
</table>

---

## 📁 Project Structure

```
puppet/
├── 🚀 cmd/puppet-api/             # Main entry point for the Go API
├── 🔧 internal/                   # Backend application logic
│   ├── 📡 handler/                # HTTP request handlers
│   ├── 🔨 service/                # Business logic layer
│   ├── 📊 repository/             # Database abstraction layer
│   ├── 📋 model/                  # Data models & entities
│   ├── 📦 dto/                    # Data Transfer Objects
│   ├── 🧩 module/                 # Dependency injection
│   └── ⚙️  config/, db/, logging/ # Configuration & infrastructure
├── 🌐 web-install/                # React frontend application
└── 🐳 docker-compose.yml          # PostgreSQL container setup
```

---

## 🌟 Features

### 🔄 Language Management
- ➕ **Add Languages** – Register new programming languages
- 📋 **List Languages** – View all available languages
- 🗑️ **Remove Languages** – Clean up unused languages
- 📦 **Docker Management** – Install/uninstall language environments

### 💻 Code Execution
- 🏃‍♂️ **Safe Execution** – Run code in isolated Docker containers
- 📥 **Stdin Support** – Provide custom input to programs
- ⚡ **Fast Response** – Optimized execution pipeline
- 🛡️ **Security** – Sandboxed environment for code execution

### 🖥️ User Interface
- 🎨 **Modern Design** – Clean, intuitive React interface
- 📱 **Responsive** – Works seamlessly across devices
- 🔄 **Real-time** – Live code execution and results
- 🎯 **User-friendly** – Simple workflow for all skill levels

---

## ⚙️ Quick Start

### Prerequisites

Make sure you have the following installed:
- 🐳 [Docker](https://docker.com) & Docker Compose
- 🐹 [Go 1.22+](https://golang.org)
- 📦 [Node.js 16+](https://nodejs.org) & npm

### 🚀 Backend Setup

1. **Start PostgreSQL Database**
   ```bash
   docker-compose up -d
   ```

2. **Configure Database Connection**
   
   Ensure your `DB_URL` points to the running PostgreSQL container:
   ```
   postgres://user:password@localhost:5432/puppet
   ```

3. **Launch API Server**
   ```bash
   go run ./cmd/puppet-api
   ```
   
   The API will be available at `http://localhost:8080` 🎉

### 🎨 Frontend Setup

1. **Navigate to Frontend Directory**
   ```bash
   cd web-install
   ```

2. **Install Dependencies**
   ```bash
   npm install
   ```

3. **Start Development Server**
   ```bash
   npm run dev
   ```
   
   The frontend will be available at `http://localhost:5173` ✨

---

## 📡 API Reference

### 🔧 Language Management

<table>
  <tr>
    <th>Method</th>
    <th>Endpoint</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><code>GET</code></td>
    <td><code>/api/languages</code></td>
    <td>📋 List all available languages</td>
  </tr>
  <tr>
    <td><code>POST</code></td>
    <td><code>/api/languages</code></td>
    <td>➕ Add a new programming language</td>
  </tr>
  <tr>
    <td><code>DELETE</code></td>
    <td><code>/api/languages/{id}</code></td>
    <td>🗑️ Remove a language</td>
  </tr>
  <tr>
    <td><code>POST</code></td>
    <td><code>/api/languages/{id}/installations</code></td>
    <td>📦 Install Docker image for language</td>
  </tr>
  <tr>
    <td><code>DELETE</code></td>
    <td><code>/api/languages/{id}/installations</code></td>
    <td>🗑️ Uninstall Docker image</td>
  </tr>
</table>

### 💻 Code Execution

<table>
  <tr>
    <th>Method</th>
    <th>Endpoint</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><code>POST</code></td>
    <td><code>/api/executions</code></td>
    <td>🏃‍♂️ Execute code in Docker container</td>
  </tr>
</table>

#### 📝 Example Execution Request

```json
{
  "languageId": 1,
  "code": "print('Hello from Puppet! 🎭')",
  "stdin": ""
}
```

#### 📊 Example Response

```json
{
  "success": true,
  "output": "Hello from Puppet! 🎭\n",
  "error": "",
  "executionTime": "0.043s"
}
```

---

## 🐳 Docker Integration

Puppet leverages Docker for secure code execution:

- **🛡️ Isolation** – Each code execution runs in a separate container
- **🔒 Security** – Sandboxed environment prevents system access
- **📦 Flexibility** – Support for multiple programming languages
- **⚡ Performance** – Efficient container lifecycle management

---

## 🤝 Contributing

We welcome contributions! Here's how you can help:

1. 🍴 **Fork** the repository
2. 🌿 **Create** a feature branch (`git checkout -b feature/amazing-feature`)
3. 💾 **Commit** your changes (`git commit -m 'Add amazing feature'`)
4. 📤 **Push** to the branch (`git push origin feature/amazing-feature`)
5. 🔄 **Open** a Pull Request

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
<div align="center">
<b>Made with ❤️</b>

[⭐ Star this repo](../../stargazers) • [🐛 Report Bug](../../issues) • [💡 Request Feature](../../issues)

</div>
