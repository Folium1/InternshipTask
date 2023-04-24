## Start the application
1.Copy the key for your mysql db source to __MYSQL__ variable <br>
2.Start mysql server. <br>
3.Run ```go run main.go```  <br>
Requests:
-(POST)localhost:9090/create/ 
-(GET)localhost:9090/books/
-(GET)localhost:9090/boo/:id
-(DELETE)localhost:9090/delete/:id
-(PUT)localhost:9090/update/