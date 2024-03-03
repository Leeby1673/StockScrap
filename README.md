美股爬蟲 CLI

概略說明
從 Yahoo Finance 抓取美國股票資訊存進資料庫，過程若發現股價下跌幅度超過百分之五，會觸發 Line 通知。
可利用子命令查看、刪除資料庫內的股票資料。
-----------------------------------------
使用說明

1. 主命令: golmy
2. 子命令: 
    catch <股票代號...> 
    抓取指定股票
    
    see
    查看資料庫全部股票
    see <股票代號...>
    查看資料庫指定的股票

    down <股票代號...>
    刪除資料庫指定的股票