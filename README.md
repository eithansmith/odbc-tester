# ODBC Tester

A simple Go utility to test ODBC connections, specifically designed for DB2/iSeries but adaptable for other ODBC sources.

## Description

`odbc-tester` is a command-line tool that verifies your ODBC configuration and connectivity by:
1. Opening a connection using a DSN, UID, and PWD.
2. Pinging the database with a 5-minute timeout.
3. Executing a sample query (`SELECT COUNT(*) FROM PCIDLIB.CIMASTRN`) to ensure data can be retrieved.

## Prerequisites

- **Go**: Version 1.25 or later is recommended.
- **ODBC Driver**: An appropriate ODBC driver (e.g., IBM i Access Client Solutions) must be installed and configured on your system.
- **DSN**: A Data Source Name (DSN) must be set up in the ODBC Data Source Administrator.

## Installation

```bash
git clone github.com/eithansmith/odbc-tester
cd odbc-tester
go build -o odbc-tester.exe
```

## Configuration

The application requires three environment variables to be set:

| Variable | Description |
|----------|-------------|
| `DB2_DSN` | The name of the ODBC Data Source Name (DSN) |
| `DB2_UID` | The database username |
| `DB2_PWD` | The database password |

### Setting environment variables in PowerShell:
```powershell
$env:DB2_DSN = "YourDSNName"
$env:DB2_UID = "YourUsername"
$env:DB2_PWD = "YourPassword"
```

### Setting environment variables in Command Prompt:
```cmd
set DB2_DSN=YourDSNName
set DB2_UID=YourUsername
set DB2_PWD=YourPassword
```

## Usage

Once the environment variables are set, run the executable:

```bash
./odbc-tester.exe
```

### Expected Output
If successful, you will see the row count from the test table:
```
Count: 1234
```

If there is a connection error, the tool will provide detailed error information:
```
PING ERR: ...
DB2 ping failed: ...
```

## Implementation Details

- Uses `github.com/alexbrainman/odbc` driver.
- Limits the connection pool to 1 open and 1 idle connection (`SetMaxOpenConns(1)`, `SetMaxIdleConns(1)`).
- Implements a 5-minute context timeout for database operations.
