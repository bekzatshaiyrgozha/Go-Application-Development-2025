1. шаги из w7/cmd/README.md если ещё не сделаны
2. go run main.go
3. New in browser (создание записи)
4. Edit (обновление записи)
5. Del (удаление записи)
6. go run gorm.go
7. change volume path in docker-compose:mysql to './_1/injection_db.sql:/docker-entrypoint-initdb.d/injection_db.sql'
8. cd ../ 
9. docker-compose down 
10. docker-compose up -d
11. cd _1/
12. go run sql_injection.go
13. login with `user` in form
14. login with `404' or login='admin`