# Sprint 2

In Sprint 1 the following user stories were successfully implemented:
1. As a user, I should be able to create a new sprint.
	* In this user story, the user has the ability to create new sprint by providing the sprint's name, .
2. As a user, I should be able to view the list of sprints.
	* In this user story, the user has the ability to view the list of sprints the user has created under that specific project. 
3. As a user, I should be able to edit the sprint details.
	* In this user story, the user has the ability to edit a sprint. The user can edit the sprint's details.
4. As a user, I should be able to delete a sprint.
	* In this user story, the user has the ability to delete the sprint.
5. As a user, I should be able to create a new issue.
	* In this user story, the user has the ability to create new issue by providing the project's name.
6. As a user, I should be able to view the list of issues.
	* In this user story, the user has the ability to view the list of issues the user has created under that specific project. 
7. As a user, I should be able to edit the issue details.
	* In this user story, the user has the ability to edit issue details. 
8. As a user, I should be able to delete an issue.
	* In this user story, the user has the ability to delete the issue.
9. As a user, I should be able to move an issue from one sprint to another.
	* In this user story, the user has the ability to move the issue from one sprint other using drag and drop.

## Video Demo Link

https://drive.google.com/drive/folders/18pB8snG21xTkadQ4i2WlHt_MyGADyhIP?usp=sharing

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





 
