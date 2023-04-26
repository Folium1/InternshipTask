## Start the application
1.Copy the key for your mysql db source to __MYSQL__ variable in .env file<br>
2.Start mysql server. <br>
3.Run ```go run main.go```  <br>
Requests:<br>
-(POST)localhost:9090/create/ <br>
-(GET)localhost:9090/books/<br>
-(GET)localhost:9090/boo/:id<br>
-(DELETE)localhost:9090/delete/:id<br>
-(PUT)localhost:9090/update/<br>
