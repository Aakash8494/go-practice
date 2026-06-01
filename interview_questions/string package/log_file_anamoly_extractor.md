Here is a complex, LeetCode-style interview question that is heavily tailored toward testing a developer's mastery of the Go `strings` package, combined with algorithmic frequency counting.

This type of question is frequently used in machine coding rounds for backend roles, as it mimics real-world text parsing and log processing tasks.

---

## **Problem: The Log File Anomaly Extractor**

**Difficulty:** Medium-Hard
**Time Limit:** 45 Minutes

### **The Scenario**

You are building an alerting system for a highly distributed Go backend. The system receives massive, unformatted chunks of text representing aggregated system logs. Your job is to parse this giant raw string, extract the valid critical anomalies, clean them up, and find the most frequent offenders.

### **The Requirements**

Write a Go function with the signature `func TopAnomalies(rawLogs string) []string` that processes the `rawLogs` string according to the following strict rules:

1. **Line Parsing:** The raw string contains multiple log entries separated by newline characters (`\n`). You must process each line individually.
2. **Filtering by Prefix:** You only care about lines that begin with exactly `"CRITICAL:"` or `"FATAL:"`. Ignore all other log levels (e.g., `"INFO:"`, `"WARN:"`).
3. **Keyword Exclusion:** If a valid critical/fatal line contains the exact, case-sensitive substring `"RETRY"`, you must completely ignore that line (it is being handled by a different self-healing system).
4. **Message Extraction & Cleaning:** For the valid lines, you must extract the core error message (everything *after* the prefix).
5. **Whitespace & Punctuation Trimming:** The extracted message must have all leading and trailing whitespace removed. Additionally, you must strip any trailing punctuation marks specifically limited to periods (`.`), exclamation points (`!`), and question marks (`?`).
6. **Frequency Aggregation:** Count the occurrences of each perfectly cleaned message.
7. **Return Value:** Return a slice of strings containing the top 3 most frequent cleaned error messages, ordered from most frequent to least frequent. If there is a tie in frequency, their relative order does not matter. If there are fewer than 3 unique anomalies, return all of them.

---

### **Example Input**

```go
rawLogs := `
INFO: System boot successful.
CRITICAL:   Database connection timeout   
WARN: Memory usage at 75%
FATAL: Segmentation fault in worker thread!
CRITICAL: Database connection timeout.
INFO: Request served in 20ms
CRITICAL: Cache miss on user profile. RETRY initiated.
FATAL:   Segmentation fault in worker thread   ?
CRITICAL: Disk space critically low!
CRITICAL: Database connection timeout!
`

```

### **Expected Output**

```go
[]string{
    "Database connection timeout", 
    "Segmentation fault in worker thread", 
    "Disk space critically low",
}

```

*Note: "Database connection timeout" appeared 3 times after cleaning. "Segmentation fault in worker thread" appeared 2 times after cleaning. The "RETRY" line was ignored.*

---

## **What the Interviewer is Secretly Grading**

This problem is a gauntlet designed to test exactly how well you know the standard library. If a candidate tries to solve this using raw `for` loops and byte-by-byte iteration, they will run out of time.

Here is what the interviewer wants to see you use:

* **`strings.Split(rawLogs, "\n")`:** To break the massive chunk of text into a slice of individual lines. *(Bonus points: A true senior might suggest using `bufio.Scanner` instead of `strings.Split` to avoid loading the entire array into memory if the string is incredibly large).*
* **`strings.HasPrefix(line, "CRITICAL:")`:** To perform a highly optimized, allocation-free check on the start of the string, rather than manually slicing `line[:9]`.
* **`strings.Contains(line, "RETRY")`:** To quickly discard the excluded messages.
* **`strings.TrimPrefix(line, "CRITICAL:")`:** To cleanly strip the log level without using string slicing index math (which can panic if the string is too short).
* **`strings.TrimSpace(message)`:** To wipe out the messy leading and trailing spaces.
* **`strings.TrimRight(message, ".!?")`:** This is the ultimate "gotcha" of the problem. Many candidates will try to write a loop with `strings.HasSuffix`. A senior Go developer knows that `TrimRight` takes a "cutset" string and strips *any* of those characters from the end of the string.
* **Data Structures:** Proper use of a `map[string]int` to aggregate the frequencies, followed by creating a slice of structs to sort the map's results using `sort.Slice`.