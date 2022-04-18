# Sprint 3
**Major Features Implemented**:
- Sign In
- Register
- JWT Authentication Across APIs.
- User Access Management In Backend.
- Backend Testcases were refactored to make use of test environment rather than the original development environment.

**User Stories Implemented**:
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

**Note**: All the APIs refactored as a result of JWT authentication that was implemented in the project. As a result user-id is not explicitly passed through the request body but is encrypted in the token which is then decrypted to retrieve the user-id of the user who made the request.

The token can be passed in the cookie
```
access_token=********; Path=/; Domain=jira-clone.com; HttpOnly; Expires=Mon, 04 Apr 2022 15:55:06 GMT;
```

OR as a **Authentication Header**
```
Authorization: Bearer ********.
```
# Updated Backend APIs Endpoints
- Any request that goes without token gets 401 status code in reply and response as follows:
    ```
    {
    "message": "ACTION_NOT_AUTHORIZED",
    "status": false
    }
    ```
- Any request with expired token gets 400 status code in reply and response as follows:
    ```
    {
    "message": "EXPIRED_TOKEN",
    "status": false
    }
    ```

-  **Create Projects**  
	*Request Format* -
``` 
	POST /api/project/create
	{
        "name" :"Project 6"
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
    
``` 
    {
        "message": "",
        "resp": {
            "project_name": "Project 6",
            "project_id": 7,
            "created_at": "2022-04-02T01:41:11.449013-04:00"
        },
    "status": true
    }
```  
<br/>

-  **List Projects**  
	*Request Format* -
``` 
	POST /api/project/list
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
    
``` 
         {
        "message": "",
        "resp": {
            "projects": [
                {
                    "name": "Project 1",
                    "id": 1,
                    "created_at": "2022-03-23T15:05:23.536313-04:00",
                    "user_role": 1
                },
                {
                    "name": "Project 5",
                    "id": 5,
                    "created_at": "2022-03-31T19:12:09.045909-04:00",
                    "user_role": 1
                },
                {
                    "name": "Project Sample PS1",
                    "id": 6,
                    "created_at": "2022-04-01T18:33:53.976626-04:00",
                    "user_role": 1
                },
                {
                    "name": "Project 6",
                    "id": 7,
                    "created_at": "2022-04-02T01:41:11.449013-04:00",
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
    GET /api/project/info
    {
     "project_id": 6,
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
        "message": "",
        "resp": {
            "project_id": 6,
            "project_name": "Project Sample PS1",
            "owner_uname": "pypalkar23",
            "owner_id": 1,
            "owner_fname": "Mandar",
            "owner_lname": "Palkar",
            "created_at": "2022-04-01T18:33:53.976626-04:00"
        },
        "status": true
    }
```  

<br/>

- **Delete a Project**	  
  *Request Format* -
``` 
    POST /api/project/delete
    {
        "project_id": 7,
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
        "message": "PROJECT_DELETE_SUCCESS",
        "resp": "",
        "status": true
    }
```  

- **List Members**	  
  *Request Format* -
``` 
    POST /api/project/members
    {
        "project_id": 1,
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
    "message": "",
    "resp": {
        "members": [
            {
                "user_id": 1,
                "first_name": "Mandar",
                "last_name": "Palkar",
                "user_role": 1
            },
            {
                "user_id": 2,
                "first_name": "Ashish",
                "last_name": "Mhaske",
                "user_role": 2
            }
        ]
    },
    "status": true
}
```  

-  **Create Sprints**  
	*Request Format* -
``` 
	POST /api/sprint/create
	{   
        "sprint_name" :"Sprint 5",
        "project_id" : 1,
        "start_date": "2021-03-07",
        "end_date" : "2021-03-21"
    }

``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
    
``` 
	{
    "message": "",
    "resp": {
        "sprint_name": "Sprint 5",
        "sprint_id": 6
    },
    "status": true
    }
```  
<br/>

-  **List Sprints**  
	*Request Format* -
``` 
	GET /api/sprint/list
	{
    	  "project_id":1
	}
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
    
``` 
{
    "message": "",
    "resp": {
        "sprints": [
            {
                "name": "Sprint 1",
                "id": 1,
                "start_date": "2021-03-06T00:00:00Z",
                "end_date": "2021-03-21T00:00:00Z",
                "created_at": "2022-03-23T15:23:01.116584-04:00",
                "project_id": 1
            },
            {
                "name": "Sprint 3",
                "id": 3,
                "start_date": "2021-03-07T00:00:00Z",
                "end_date": "2021-03-21T00:00:00Z",
                "created_at": "2022-03-29T19:55:12.21721-04:00",
                "project_id": 1
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
    POST /api/sprint/info
    {   
        "sprint_id" :3
    }

``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
        "message": "",
        "resp": {
            "sprint_id": 3,
            "sprint_name": "Sprint 1",
            "created_at": "2022-03-29T19:55:12.21721-04:00",
            "start_date": "2021-03-07T00:00:00Z",
            "end_date": "2021-03-21T00:00:00Z"
        },
        "status": true
    }
```  
<br/>

- **Delete a Sprint**  	  
	*Request Format* -
``` 
    POST /api/sprint/delete
    {
     "sprint_id": 6,
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
        "message": "SPRINT_DELETE_SUCCESS",
        "resp": "",
        "status": true
    }
``` 
<br/> 

-  **To Create an Issue**. 
	*Request Format* -
```
    POST /api/issue/create
   {    
        "issue_title":"sample issue title",
        "issue_description": "sample issue description",
        "issue_type": 1,
        "creator": 2,
        "assignee": 1,
        "sprint_id": 1,
        "project_id": 1
    }
```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
    {
        "message": "",
        "resp": {
            "issue_id": 0,
            "issue_title": "sample issue title",
            "issue_description": "sample issue description",
            "issue_type": 1,
            "creator_name": "amhaske32",
            "assignee_name": "pypalkar23",
            "created_at": "2022-04-02T02:01:04.593536-04:00",
            "issue_status": 1
        },
        "status": true
    }

```
<br/>

- **To get a issue’s info**  
	*Request Format* -
```
    POST api/issues/info
    {
        "issue_id": 3
    }
```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
    {
        "message": "",
        "resp": {
            "issue_id": 3,
            "title": "Sample Issue Edited 1",
            "type": 1,
            "sprint_id": 3,
            "sprint_name": "Sprint 1",
            "project_id": 1,
            "description": "Sample Text Edited 5",
            "assignee_id": 1,
            "assignee_name": "pypalkar23",
            "creator_id": 2,
            "creator_name": "amhaske32",
            "created_at": "2022-03-23T18:05:00.58531-04:00",
            "project_name": "Project 1",
            "issue_status": 2
        },
        "status": true
    }
```
<br/>

- **To list issues for sprint**  
 	*Request Format* -
```
    GET api/issue/list
    {
        "sprint_id": 1
    }
```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
    {
        "message": "",
        "resp": {
            "issues": [
                {
                    "issue_id": 1,
                    "title": "issue title",
                    "status": 1,
                    "created_at": "2022-03-23T16:39:17.532173-04:00"
                },
                {
                    "issue_id": 4,
                    "title": "issue title 4",
                    "status": 1,
                    "created_at": "2022-03-29T22:53:55.005758-04:00"
                },
                {
                    "issue_id": 5,
                    "title": "sample issue title",
                    "status": 1,
                    "created_at": "2022-04-02T02:01:04.593536-04:00"
                }
            ]
        },
        "status": true
    }
```
<br/>

- **To delete an issue**  	  
 	*Request Format* -
``` 
    POST /api/issue/delete
    {
        "issue_id": 4
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
        "message": "ISSUE_DELETE_SUCCESS",
        "resp": "",
        "status": true
    }
``` 

- **To update an issue**  	  
 	*Request Format* -
``` 
    POST /api/issue/delete
    {
        "issue_id": 3,
        "issue_title": "Sample Issue Edited 1",
        "issue_description": "Sample Text Edited 5",
        "issue_type": 1,
        "creator": 1,
        "assignee": 2,
        "sprint_id": 3,
        "project_id": 1
    }   
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
        "message": "ISSUE_UPDATE_SUCCESS",
        "resp": {},
        "status": true
    }
``` 

- **To update an issue**  	  
	*Request Format* -
``` 
    POST /api/issue/move
    {
        "issue_id":3,
	    "project_id":1,
	    "status":2
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
        "message": "ISSUE_UPDATE_SUCCESS",
        "resp": {},
        "status": true
    }
``` 

- **To Register a User**  
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
        "password":"**********"
    }
```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
   {
        "message": "LOGIN_SUCCESSFUL",
        "resp": {
            "access_token": "******"
        },
        "status": true
    }
```
<br/>


- **To Add Comments To An Issue**  
    *Request Format* -
```
    POST /api/comment/add
    {
        "issue_id": 3,
        "comment": "This is a comment for issue_id: 3"
    }

```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
    {
        "message": "",
        "resp": {
            "comment" : "This is a comment for issue_id: 3"
        },
        "status": true
    }
```
<br/>

- **View Comments**  
    *Request Format* -
```
    POST /api/comment/view
    {
        "issue_id": 3
    }

```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
   {
        "message": "",
        "resp": {
            "comments": [
                {
                    "comment_id": 1,
                    "comment": "This is a comment for issue_id: 3"
                }
            ]
        },
        "status": true
    }
```
<br/>


## Test Case Refactor:  
<br/>
The code associated with test cases is in **tests** directory and is divided into three files.  

- **setup_test.go** : Creates a test environment(Mock Database, User Entries, Tokens) prior to running test cases.
- **tests_data.go** : Contains list of test cases
- **run_test.go** : Driver code that picks a test case one by one from *tests_data.go* and runs those testcases.

**How to run the test cases**: The test cases can be executed by running the command ```go test -v ./tests```

Rather than writing multiple test functions for multiple testcases which require setup of environment prior to running actual test code, the testcases were combined into an array of custom designed struct. The structure of a single entry for a test case that tests the API is as follows:
```
	{
	    init function,  
	    REST method - e.g. GET,POST,PUT etc) 
	    Actual API handler -e.g. "/api/sprint/delete")
	    Post Body, - e.g. `{"sprint_id":2}`)
	    Expected Status Code - e.g. 200,401,400 etc.)
	    Expected Response - e.g. `{"message":"SPRINT_DELETE_SUCCESS","resp":"","status":true}`)
	    post function(optional), - function that tests for changes happened in test environment after calling an API
	}

An example:
	{
	    func(req *http.Request) {
	    	// init function sets cookie into the request
		setTokenRequestInCookie(req)
	    },
	    "/api/project/create",
	    "POST",
	    `{"name":"Project 1"}`,
	    http.StatusOK,
	    `{"message":"","resp":{"project_name":"Project 1","project_id":([0-9]+),"created_at":"([A-Z0-9:\-\.]+)"},"status":true}`,
	    func() bool {
	        //Post function that checks whether the user entry is created or not in the user table.
		var count int64
		dbutils.DB.Where("role_id = 1 AND user_id = 1 AND project_id = 1").Find(&models.UserRole{}).Count(&count)
		return count == 1
	      },
	 },
```

 
