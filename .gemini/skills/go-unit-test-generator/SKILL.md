---
name: go-unit-test-generator
description: Generate Go unit tests following the project's table-driven testing pattern and BDD naming conventions. Use when adding new features or fixing bugs in the Go codebase.
---

# Go Unit Test Generation

Follow this pattern to ensure consistent, testable, and maintainable Go code.

## Testing Pattern

The project follows a consistent table-driven testing pattern using `testify` mocks and **Behavior-Driven Development (BDD)** style naming:

- **BDD Naming:** Test case `name` fields should follow a "should [behavior] when [condition]" format (e.g., `"should successfully load .env when file exists"`).
- **Mock Grouping:** Define a `mocks` struct within the test function to hold all dependency mocks.
- **Table-Driven Structure:** Use a slice of anonymous structs for test cases, including a `setup` func to define mock expectations.
- **Mock Initialization:** Initialize mocks using `NewDependency(t)` (or equivalent) for each test iteration.
- **Expectations:** Use `.EXPECT()` on mocks within the `setup` function.
- **Assertions:**
    - Use `errors.Is(gotErr, tt.wantErr)` for error validation.
    - Use `reflect.DeepEqual` or `testify` assertions for state validation.
- **Context:** Pass `context.Background()` to methods requiring a `context.Context`.

### Example Structure

```go
func TestService_Method(t *testing.T) {
    type mocks struct {
        dep *dep_mocks.Mock
    }
    tests := []struct {
        name    string
        setup   func(m mocks)
        wantErr error
    }{
        {
            name: "success",
            setup: func(m mocks) {
                m.dep.EXPECT().DoSomething().Return(nil)
            },
            wantErr: nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            m := mocks{dep: dep_mocks.NewMock(t)}
            if tt.setup != nil { tt.setup(m) }
            // ... execute and assert
        })
    }
}
```
