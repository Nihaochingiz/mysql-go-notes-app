

1 запуск
docker run --name mysql-db -e MYSQL_ROOT_PASSWORD=0000 -e MYSQL_DATABASE=runningdb -p 3308:3306 -d mysql:latest
2, далее
docker exec -it mysql-db mysql -uroot -p runningdb -e "CREATE TABLE running_statistics (date DATE, distance VARCHAR(10), time VARCHAR(10));"


1. docker-compose up
2. go run create_note.go
3. docker exec -it docker-mysql-db-1 mysql -uexample_user -pexample_password example_db
4. select * from example_db;
5. show tables;
6. select * from notes;