# Sprint 2

In Sprint 2 the following user stories were successfully implemented:
1. As a user, I should be able register in the system.
	* In this user story, the user has the ability to register on the platform .
2. As a user, I should be able to login to the system.
	* In this user story, the user has the ability to login to the plaform. 
3. As a user, I should be able to reassign an issue from one user to another.
	* In this user story, the user has the ability transfer the assignment of an issue from one user to another.
4. As a user, I should be able add comments to an issue.
	* In this user story, the user has the ability to add comments in order to discuss and comment on issues.
5. As a user, I should be able to view comments for an issue.
	* In this user story, the user has the ability to view all previously added comments for an issue.
6. As a user of type admin, I should be able to delete issues.
	* In this user story, the user has the ability to  delete issues in case it is of type admin
7. As a user, I should be able to view the issue details.
	* In this user story, the user has the ability to view issue details. 


## Video Demo Link

https://drive.google.com/drive/folders/1b5kKf9ewZIVxmkyzyQO8p1Qys7SZbx8j?usp=sharing

# SQLite Models of Tables
- projects
	```
	project_id INTEGER PRIMARY KEY
	project_name       TEXT
	is_deleted NUMERIC
	owner_id   NUMERIC  (FOREIGN KEY Users)
	created_at DATETIME 
	updated_at  DATETIME
    name        TEXT
	```

- ROLES
	```
	role_id INTEGER PRIMARY KEY
    role_name TEXT
	permission_ref INTEGER PRIMARY KEY
	permission_id INTERER 
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
	firstname  TEXT
    lastname    TEXT
	is_deleted TEXT
	email_id   TEXT
	created_at DATETIME
    name    TEXT
	```

- Permissions
	```
	permission_id    INTEGER PRIMARY KEY
	permission_name   TEXT
	
	```

- SPRINTS
	```
	sprint_id    INTEGER PRIMARY KEY
	sprint_name   TEXT
	project_ref  INTEGER
    created_at    DATETIME
	updated_at DATETIME
    start_date  DATETIME
    end_date    DATETIME
	status   INTEGER
	is_deleted NUMERIC
    project_id    INTEGER
	```


- ISSUE
	```
	issue_id    INTEGER PRIMARY KEY
	status   INTEGER
	type  INTEGER
    title    TEXT
	created_by INTEGER
    sprint_ref  INTEGER
    assignee_id    INTEGER
	project_ref   INTEGER
	is_deleted NUMERIC
    updated_at    DATETIME
    description TEXT
    created_at	DATETIME
	sprint_id   INTEGER
	```

-COMMENT
    	```
    comment_id INTEGER PRIMARY KEY
	issue_id    INTEGER 
    comment_text TEXT
	user_id   INTEGER
    created_at	DATETIME
    updated_at	DATETIME
    deleted_at  DATETIME
	
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

-  **Create Sprints**  
	*Request Format* -
``` 
	POST /sprint/create
	{
	   //json body
	   "sprint_name" :"Sprint 1",
       "start_date": "02/22/2022",
        "end_date": "22/22/2022",
        "project_id":2
	}
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
    
``` 
	{
	   "message": "Sprint Created Successfully",
	    "resp": {
		"sprint_name" :"Sprint 1",
        "project_id":2,
        "sprint_id": 2,
        "start_date": "02/22/2022",
        "end_date": "22/22/2022"
	    },
	    "status": true
	}
```  
<br/>

-  **List Sprints**  
	*Request Format* -
``` 
	GET /sprint/list
	{
	  //json body
    	  "project_id":2
	}
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
    
``` 
     {
     "message": "",
     "resp": {
        "sprints": [
            {
                "sprint_name" :"Sprint 1",
                "project_id":2,
                "sprint_id": 2,
                "start_date": "02/22/2022",
                "end_date": "02/29/2022"
            },
            {
                "sprint_name" :"Sprint 2",
                "project_id":2,
                "sprint_id": 2,
                "start_date": "02/22/2022",
                "end_date": "02/29/2022"
            },
            {
                "sprint_name" :"Sprint 3",
                "project_id":2,
                "sprint_id": 2,
                "start_date": "02/23/2022",
                "end_date": "22/28/2022"
            }
        ]
    },
    "status": true
   }
```  



<br/>

-  **To get a sprint's info**  
	*Request Format* -
``` 
    GET /sprint/info
    {
     "project_id": 6,
     "sprint_id":2
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
   {
    "message": "",
    "resp": {
        
                "sprint_name" :"Sprint 3",
                "project_id":2,
                "sprint_id": 2,
                "start_date": "02/23/2022",
                "end_date": "22/28/2022"
    },
    "status": true
   }
```  

<br/>

- **Delete a Sprint**	  
*Request Format* -
``` 
    POST /sprint/delete
    {
     "sprint_id": 6,
     "user_id":2
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
   {
    "message": "Sprint deleted successfully",
    "resp": "",
    "status": true
   }
``` 
<br/> 

-  **To Create an Issue**
    *Request Format* -
```
    GET /issue/create
    {
      “user_id”:1,
      “sprint_id”:1,
      “issue_name”:“Sample Issue 5",
      “issue_text”:“Thing ABC is not working”,
      “creator”:1,
      “assignee”:2,
      “project_id”:1
    }
```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
     {
    “message”: “Issue Created Successfully”,
    “resp”: {
        “user_id”: 0,
        “issue_id”: 5,
        “issue_name”: “Sample Issue 5"
    },
    “status”: true
}
```
<br/>
-  **To get a issue’s info**
    *Request Format* -
```
    GET /issues/update
    {
    “user_id”:1,
    “issue_id”:5,
    “sprint_id”:1,
    “issue_name”:“Sample Issue 6",
    “issue_text”:“Thing CBSX is not working”,
    “creator”:1,
    “assignee”:2,
    “project_id”:1
    }
```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
{
    “message”: “Issue Updated Successfully”,
    “resp”: “”,
    “status”: true
}
```
<br/>
- **To list issues for sprint**
*Request Format* -
```
    GET /issues/list
    {
    “project_id”:1,
    “sprint_id”:1
    }
```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
   {
    “message”: “”,
    “resp”: {
        “issues”: [
            {
                “issue_id”: 1,
                “title”: “Sample Issue 1”,
                “status”: 1
            },
            {
                “issue_id”: 2,
                “title”: “Sample Issue 1”,
                “status”: 1
            },
            {
                “issue_id”: 3,
                “title”: “Sample Issue 3”,
                “status”: 1
            },
            {
                “issue_id”: 5,
                “title”: “Sample Issue 6”,
                “status”: 1
            }
        ]
    },
    “status”: true
}
```
<br/>
<br/>
- **To Register**
*Request Format* -
```
    POST /register
    {
    "username": "dipesh97",
    "password": "1c1d15237b2433f18f588d8bf6984751",
    "firstname": "Dipesh",
    "lastname": "Palkar",
    "email_id": "dipeshpalkar@gmail.com"
    }
```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
   {
    "message": "REGISTERATION_SUCCESSFUL",
    "resp": "",
    "status": true
}
```
<br/>


- **To Login**
*Request Format* -
```
    POST /login
    {
   "username":"pypalkar23",
   "password":"dd29b8cb089a56606fca480e137c27c4"
}
```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
   {
    "message": "LOGIN_SUCCESSFUL",
    "resp": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NDg4NjE2NTMsImlhdCI6MTY0ODg1ODA1M30.W_k2WsrrP01Hzi81os5efb-9UCnGga08ueOA4l4et3c"
    },
    "status": true
    }
```
<br/>


- **To Add Comments**
*Request Format* -
```
    POST /comment/add
    {
    "issue_id": 3,
    "comment": "This is a comment for issue_id: 3"
    }

```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
   {
    “message”: “”,
    “resp”: {
        “comment”: “This is a comment for issue_id: 3”
    },
    “status”: true
}
```
<br/>

- **View Comments**
*Request Format* -
```
    GET /comment/view
    {
    "issue_id": 3
    }

```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
   {
    "message": "ACTION_NOT_AUTHORIZED",
    "status": false
}
```
<br/>

- **List Project Members**
*Request Format* -
```
api/project/members
req:
{
    “project_id”: 1
}
```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```

{
    “message”: “”,
    “resp”: {
        “members”: [
            {
                “user_id”: 1,
                “first_name”: “Mandar”,
                “last_name”: “Palkar”,
                “user_role”: 1
            },
            {
                “user_id”: 2,
                “first_name”: “Ashish”,
                “last_name”: “Mhaske”,
                “user_role”: 2
            }
        ]
    },
    “status”: true
} 

```
<br/>
- **Issue Status change**
*Request Format* -
```
    POST api/issue/move
     {
    “issue_id”:3,
    “project_id”:1,
    “status”:2
}

```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
{
    “message”: “ISSUE_UPDATE_SUCCESS”,
    “resp”: {},
    “status”: true
}
 
```
<br/>


 
