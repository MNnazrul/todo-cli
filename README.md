# todo-cli

A command-line TODO application written in Go, backed by MongoDB. Manage your tasks efficiently with commands like `add`, `list`, and more!

---

Project Structure:

- .env.example → Example environment variables
- db.go → MongoDB connection logic
- go.mod → Go module definitions
- go.sum → Dependency checksums
- main.go → CLI entry point
- todo.go → Todo model and CRUD operations

---

Features:

- Add a new task with optional description and status
- List all tasks in a table view
- Store tasks in a MongoDB database
- Status support: "todo", "in-progress", "done"
- Clean CLI interface using standard os.Args

---

Requirements:

- Go 1.18 or later
- MongoDB (local or Atlas)
- Internet connection (for MongoDB Atlas)

---

Setup:

1. Clone the repository:

   git clone [https://github.com/your-username/todo-cli.git](https://github.com/your-username/todo-cli.git)
   cd todo-cli

2. Setup environment variables:

   Copy `.env.example` to `.env`:

   cp .env.example .env

   Then edit `.env`:

   MONGODB_URI=mongodb+srv://your-user\:your-password\@cluster.mongodb.net
   MONGODB_DB=todo_db
   MONGODB_COLLECTION=todos

3. Install dependencies:

   go mod tidy

---

Usage:

To add a task:

go run . add "Write blog" "Write a blog post on Go modules"

To list all tasks:

go run . list

---

Status Field:

Each task has a `status` field that can be one of the following:

- todo – Not started
- in-progress – Currently working on
- done – Task completed

---

### Example Output

| #   | Task           | Status      | Description                   | Created At               |
| --- | -------------- | ----------- | ----------------------------- | ------------------------ |
| 1   | Write blog     | todo        | Write a blog post on Go       | Jan 08, 2025, 03:45:02PM |
| 2   | Build todo CLI | in-progress | Go project for managing tasks | Jan 09, 2025, 12:30:10PM |

---

### Credits

Project idea inspired by: [Task Tracker - roadmap.sh](https://roadmap.sh/projects/task-tracker)

---

Author:

    Nazrul Islam
    GitHub: [@MNnazrul](https://github.com/MNnazrul)

---
