# Jira-Clone
Software Project Management Tool plays a pivotal role in the lifecycle and success of any Software Project. This project by us strives to address that issue. This project aims to provide a clean UI with minimalistic but most important features of a project management tool.  

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

</br>
</br>

## The project is accessible at -
- Project has been deployed on a *Digital Ocean* Server
- **Frontend** - https://app.se-gators.co.in
- **Backend** - https://api.se-gators.co.in

## Member List -
  
|Project Group Members|Github User Name|Developer Type
-|-|-
Ashish Sanjay Mhaske| amhaske | Frontend
Divya Bhutani | divyabhutani | Frontend
Mandar Palkar | pypalkar23 | Backend
Vaishnavi Dongre | vaishnavi3001 | Backend


## Technical Stack
- Frontend - **Angular 12, Angular Material**
- Backend - **Golang 1.17**


## How to Build and Run -
### Prerequisites: 
- NodeJS & NPM - https://nodejs.org/en/     
- Golang - https://go.dev/

### How to run for development purposes
    # Frontend -
    1. Go to jira-frontend directory
    2. Run the following commands  
        - `npm i`   
        - `ng serve`
    The frontend will run on `http://localhost:4200`

    # Backend -
    1. Go to jira-backend directory
    2. Run the following commands
        - `go run main.go`
    The backend will run on `0.0.0.0:6000`

### How to create build executables for production
    # Frontend -
        1. Go to jira-frontend directory
        2. Run the following command
            - ng build --prod 
        3. Then go to node-server & run :
            - node app.js
        The frontend will run on `0.0.0.0:8500`

    #Backend -
        1. Go to jira-backend directory
        2. Run the following command
            - go build prod
        This will create an executable called jira-backend in the same location
        3. Execute the shell script using `./jira-backend`
        The backend will run on '0.0.0.0:5000' with prod configuration.

## Sprint 4 Project Board Link:
 https://github.com/users/vaishnavi3001/projects/6/views/1

## Work Done in Sprint 4 is described here:
 https://github.com/vaishnavi3001/jira-project/blob/main/Sprint4.md

## API Docs Can be found here:
 https://github.com/vaishnavi3001/jira-project/blob/main/API_DOCS.md

## Backend Testing Execution:
https://user-images.githubusercontent.com/89871418/164365728-f13841e7-332a-4d78-bd96-99994af997b8.mp4


        


















