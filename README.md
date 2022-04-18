# Jira-Clone
Software Project Management Tool plays a pivotal role in the lifecycle and success of any Software Project. This project by us strives to address that issue. This project aims to provide a clean UI with minimalistic but most important features of a project management tool.  
  
## The project is accessible at -
- **Frontend** - https://app.se-gators.co.in
- **Backend** - https://api.se-gators.co.in


|Project Group Members|Github User Name|Developer Type
-|-|-
Ashish Sanjay Mhaske| amhaske | Frontend
Divya Bhutani | divyabhutani | Frontend
Mandar Palkar | pypalkar23 | Backend
Vaishnavi Dongre | vaishnavi3001 | Backend


## Technical Stack
- Frontend - **Angular 12, Angular Material**
- Backend - **Golang 1.17**


## What is included in the project -

- **Sprint Tracking View** -  
A view that shows the issues on a "Kanban" type board, highlighting the status of the issue for the sprint like "In Progress", "Completed" etc.  

- **Issue View** -  
A view that gives details about the issue like when was the issue created, who has been assigned to work on the issue etc. 

- **Project Settings** -  
A view that allows the user to modify the settings for the project like editing details, select issue types to include in a project etc.  

- **User Hierarchy** -    
User Hierarchy will allow the users to take roles like project owner, developer, etc.

- **User Profile** -  
User Profile will allow user to set his personal details as well as to view all issues assigned to him at a glance.

- **User Invite** -  
User Invite feature will allow the project owner to invite other users to join the project.

- **Sign up / Sign in** -
This feature will help authenticate the user.


## To Build and Run Go Package Of Angular Build
- Navigate to `jira-frontend` directory
- make sure `outputPath` property in `angular.json` to `build/assets/jira-frontend`
- navigate to `build` directory
- run `go build jira-build`
- run `go run jira-build` or `./jira-build`
- Server will run on `0.0.0.0:8080`








