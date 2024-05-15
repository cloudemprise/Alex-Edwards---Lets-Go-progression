package models

/*
Alex Edwards : Let's Go

Chapter 4 Database-driven responses

Section 4.6 Executing SQL statements
*/

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {

	// SQL statement. Backticks for multi-line string.
	// Write the SQL statement with the RETURNING clause to retrieve the id
	stmt := `INSERT INTO snippets (title, content, created, expires)
             VALUES($1, $2, NOW(), NOW() + ($3 * INTERVAL '1 DAY')) RETURNING id`

	/* // MySQL version
	// ? = placeholders for query parameters
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))` */

	// Potentially vulnerable to SQL injection attacks because of the modification.
	var id int
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	/* // MySQL version:
	// Use the Exec() method on the embedded connection pool to execute the
	// statement. The first parameter is the SQL statement, followed by the
	// title, content and expiry values for the placeholder parameters. This
	// method returns a sql.Result type, which contains some basic
	// information about what happened when the statement was executed.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	} */

	/* // MySQL version:
	// Use the LastInsertId() method on the result to get the ID of our
	// newly inserted record in the snippets table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	} */

	return id, nil

	/* // MySQL version:
	// The ID returned has the type int64, so we convert it to an int type
	// before returning.
	return int(id), nil */
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}

/*
Since PostgreSQL does not support the `LastInsertId()` method, you'll need to modify the code to retrieve the ID of the newly inserted row in a different way. Here's how you can do it:

1. In your SQL `INSERT` statement, add the `RETURNING id` clause at the end. This will make PostgreSQL return the ID of the newly inserted row.

```sql
INSERT INTO snippets (title, content, created, expires)
VALUES($1, $2, NOW(), NOW() + ($3 * INTERVAL '1 DAY'))
RETURNING id
```

2. Modify the `Insert` method in `snippets.go` to use the `QueryRow` method instead of `Exec`. This will allow you to retrieve the returned ID.

```go
func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
    stmt := `INSERT INTO snippets (title, content, created, expires)
             VALUES($1, $2, NOW(), NOW() + ($3 * INTERVAL '1 DAY')) RETURNING id`

    var id int
    err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)
    if err != nil {
        return 0, err
    }

    return id, nil
}
```

In this modified code:

- The `INSERT` statement now includes the `RETURNING id` clause, which will make PostgreSQL return the ID of the newly inserted row.
- Instead of using `Exec`, we use `QueryRow` to execute the statement and retrieve the returned ID.
- The `Scan` method is used to copy the returned ID into the `id` variable.

With these changes, your code should work correctly with PostgreSQL, and you'll be able to retrieve the ID of the newly inserted row without relying on the `LastInsertId()` method.
*/
