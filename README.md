# timesheet-generator
Generate a timesheet for billing a client

# Building
**NO dependencies**. 
Edit the variable ``RootPath`` in ``main.go``, set it to the directory you want your timesheet written to.
Run go build timetracker in the root directory.

# Usage
```sh
./timetracker add <sheetname>
./timetracker <sheetname> # to toggle timing, from on->off or off->on,
./timetracker export <sheetname> # export from the plain text format to a latex table.
```
  
