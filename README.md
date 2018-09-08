# GoTask
by: Coleman Word

## API Design
Code available in main.go

| ROUTE          	| METHOD 	| DESCRIPTION             	|
|----------------	|--------	|-------------------------	|
| /add/          	| POST   	| Add a new task          	|
| /              	| GET    	| Get Tasks               	|
| /complete/     	| GET    	| Show Pending Tasks      	|
| /deleted/      	| GET    	| Show Deleted Tasks      	|
| /edit/<id>     	| POST   	| Edit Post               	|
| /edit/<id>     	| GET    	| Show Edit View          	|
| /trash/<id>    	| POST   	| Send to Trash           	|
| /delete/<id>   	| POST   	| Permanently Delete Task 	|
| /complete/<id> 	| POST   	| Complete Tasks          	|
| /login/        	| POST   	| Login                   	|
| /login/        	| GET    	| Show Login Page         	|
| /logout/       	| POST   	| Logout                  	|
| /restore/<id>  	| POST   	| Restore a Task          	|
| /update/<id>   	| POST   	| Update a Task           	|
| /change/       	| GET    	| Change Password         	|
| /register/     	| GET    	| Show Register Page      	|
| /register/     	| POST   	| Register in Database    	|
  
## Screenshots
![](https://github.com/ops2go/screenshots/blob/master/gotask/gotask-login.png?raw=true)
![](https://github.com/ops2go/screenshots/blob/master/gotask/gotask-home.png?raw=true)
![](https://github.com/ops2go/screenshots/blob/master/gotask/gotask-categories.png?raw=true)
![](https://github.com/ops2go/screenshots/blob/master/gotask/gotask-add.png?raw=true)


## TO DO
* Move HTML Rendering Client Side
* Update Color Scheme to Black, White, and Blue
* Add Angular/Ajax/React Front End


## DOCKER
```
docker build -t gotask .
```
```
docker run -d -p 8080:8081 gotask
```
