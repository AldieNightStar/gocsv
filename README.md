# CSV Loader

# Import
```go
import "github.com/AldieNightStar/gocsv"
```

# Usage
```go
// Load the file
file, _ := os.ReadFile("someFile.csv")

// Load CSV from this file
// Accepts io.Reader
csv := gocsv.Load(file)

// Print all values at first row
// type: string
firstRow := csv.Rows[0]
for i := range firstRow.Cols {
    t.Log(firstRow.Cols[i])
}

// Get First row 0, 1 value
print(firstRow.Cols[0])
print(firstRow.Cols[1])

// Get value from 1,2
val := csv.Rows[1].Cols[2]
print(val)
```