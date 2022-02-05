# Sprint 1

In Sprint 1 the following user stories were successfully implemented:
1. As a user, I should be able to create a new project.
	* In this user story, the user has the ability to create new project by providing the project's name, key.
2. As a user, I should be able to view the list of projects.
	* In this user story, the user has the ability to view the list of projects the user has created. Each item in the list contains the project's name, key, username
3. As a user, I should be able to edit the project details.
	* In this user story, the user has the ability to edit a project. The user can edit the project's name, key.
4. As a user, I should be able to delete a project.
	* In this user story, the user has the ability to delete the project.

## Video Demo Link

https://drive.google.com/drive/folders/1NO4WvOJx3Gk6evxFsizYsg27xtTjT9BG

# SQLite Models of Tables
- projects
	```
	project_id INTEGER PRIMARY KEY
	name       TEXT
	is_deleted NUMERIC
	owner_id   NUMERIC  (FOREIGN KEY Users)
	created_at DATETIME 
	UpdatedAt  DATETIME
	```

- user_roles
	```
	user_id INTEGER PRIMARY KEY
	role_id INTEGER PRIMARY KEY
	project_id INTEGER PRIMARY KEY
	membership 
	```

- Users
	```
	user_id    INTEGER PRIMARY KEY
	username   TEXT
	name 	   TEXT
	is_deleted TEXT
	email_id   TEXT
	created_at DATETIME
	```
# Backend APIs Endpoint
-  **Create Projects**  
	*Request Format* -
``` 
	POST /project/create
	{
	   //json body
	   "name" :"Boring Company",
    	   "user_id":2
	}
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
    
``` 
	{
	   "message": "Project Created Successfully",
	    "resp": {
		"project_name": "Boring Company",
		"project_id": 6
	    },
	    "status": true
	}
```  
<br/>

-  **List Projects**  
	*Request Format* -
``` 
	GET /project/list
	{
	  //json body
    	  "user_id":2
	}
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
    
``` 
     {
     "message": "",
     "resp": {
        "projects": [
            {
                "name": "First Project",
                "id": 2,
                "created_at": "2022-02-04T22:18:08.692549-05:00",
                "user_role": 1
            },
            {
                "name": "Prometheus",
                "id": 4,
                "created_at": "2022-02-04T22:18:39.179794-05:00",
                "user_role": 1
            },
            {
                "name": "Boring Company",
                "id": 6,
                "created_at": "2022-02-05T13:26:59.794848-05:00",
                "user_role": 1
            }
        ]
    },
    "status": true
   }
```  


``` 
    {
     "message": "Project Created Successfully",
     "resp": {
        "project_name": "Boring Company", 
	"project_id": 6
      },
     "status": true
    }
```  

<br/>

-  **To get a project's info**  
	*Request Format* -
``` 
    GET /project/info
    {
     "project_id": 6,
     "user_id":2
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
   {
    "message": "",
    "resp": {
        "project_id": 6,
        "owner_uname": "amhaske32",
        "owner_id": 2,
        "owner_name": "Ashish Mhaske",
        "project_name": "Boring Company",
        "created_at": "2022-02-05T13:26:59.794848-05:00"
    },
    "status": true
   }
```  

<br/>

- **Delete a Project**	  
*Request Format* -
``` 
    POST /project/delete
    {
     "project_id": 6,
     "user_id":2
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
   {
    "message": "Project deleted successfully",
    "resp": "",
    "status": true
   }
```  



 
