## 美股爬蟲 CLI

**Overview:**
This program fetches US stock information from Yahoo Finance and stores it in a database. If the program detects a stock price drop exceeding five percent, it triggers a Line notification. Additionally, it offers subcommands to view and delete stock data from the database.

**Usage:**
Main Command: golmy

Subcommands:

catch "stock symbols...": Fetch specified stocks.

see: View all stocks in the database.
see "stock symbols...": View specified stocks in the database.

down "stock symbols...": Delete specified stocks from the database.