## 美股爬蟲 CLI

**Overview:**
The main function of the program is to fetch stock information for the user by inputting stock symbols. Users can set parameters to determine whether to trigger Line Notify. The program also allows querying and deletion actions.

# Usage:
### Main Command: 

golmy

### Subcommands:

**catch "symbols..."**: Fetch specified stocks and store the stock information in the database.

**-o**: Activate ongoing mode, checking every 15 seconds, without storing stock information in the database, "Ctrl+C" to exit the program.

**-l**: Activate Line Notify, input parameters required; trigger Line Notify when the stock price changes by n%.

**Example**: go run main.go catch NVDA -o -l=5

Fetch NVIDIA stocks, activate continuous monitoring mode, and trigger Line Notify when there is a 5% increase.

**see**: If no parameter is provided, view all stock information in the database.

**see "symbols..."**: If a parameter is provided, view information for specified stocks in the database.

**-p**: View all stocks in the database with prices below n, requiring input parameters but no symbols.

**Example**: go run main.go see MU or go run main.go see -p=100

View information for Micron stocks in the database or view stocks in the database priced below $100.

**down "symbols..."**: Delete information for specified stocks from the database.

**Example**: go run main.go down TSLA

Delete information for Tesla stocks from the database.