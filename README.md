### How To Run This Project
> Make Sure you have run the database.sql in your mysql


#### Run the Applications
- Set the database and port in the `config.json` file.
  ```json
  {  
    "server": {  
      "address": "8101"  
    },  
     "database": {  
      "host": "127.0.0.1",  
      "port": "3306",  
      "user": "root",  
      "password": "root",  
      "dbname": "recipe_management"  
    }  
  }
  ```
- Run the following command
  ```bash
  go run app/main.go
  ```
- Your program will start running, and the server will listen on the port you specified in the `config.json` file.