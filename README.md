# DSA Sheet Solutions in Go 🚀

Solutions to the [Apna College DSA Sheet](https://dsa.apnacollege.in/) implemented in **Go (Golang)**.

Each topic is organized as its own **Go package**, keeping the codebase clean and modular.

---

## 📁 Project Structure

```
.
├── main.go
├── topics/
│   ├── topic.go           # Interface definitions
│   └── questions.go       # Implementations
```

---

## 📦 Packages
Lets take arrays as example
### `arrays`

Solutions to array-based problems from the DSA sheet.
Go in the arrays package
you will find all the questions interface there.


Lets take Majority Element as a example

| Problem | Approaches |
|---|---|
| Majority Element | Brute Force, HashMap, Sorting, Boyer-Moore Voting |

#### Majority Element
go to majority_element.go 
you will find all the implementation

## 🚀 Getting Started

### Prerequisites

- Go 1.18+

### Run

```bash
git clone https://github.com/<your-username>/<your-repo>.git
cd <your-repo>
go run main.go
```

### Run Tests (coming soon)

```bash
go test ./...
```

---

## 🛠️ How to Use a Different Method

In `main.go`, swap the method call:

```go
// HashMap approach
ans := array.MajorityElementHashMap(nums)

// Sorting approach
ans := array.MajorityElementSorting(nums)

// Boyer-Moore (optimal)
ans := array.MajorityElementBoyerMoore(nums)
```


## 🤝 Contributing

This is a personal learning repo, but feel free to open issues or PRs if you spot a bug or a better approach!

---

## 📄 License

MIT
