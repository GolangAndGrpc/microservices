@echo off

set DATA_SOURCE_URL=root:root@tcp(localhost:3306)/orders
set APPLICATION_PORT=8089
set ENV=dev


echo APPLICATION_PORT=%APPLICATION_PORT%
echo ENV=%ENV%
echo DATA_SOURCE_URL=%DATA_SOURCE_URL%

pause