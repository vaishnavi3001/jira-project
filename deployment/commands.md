### Services Start Instructions
1. sudo systemctl start jira-frontend
2. sudo systemctl start jira-backend
<br/>

### Service File Locations
Backend Service - /lib/systemd/system/jira-backend.service
Frontend Service - /lib/systemd/system/jira-frontend.service
"Symbolic Links Are Stored In" - ```/etc/systemd/system```     
<br/>

### Certbot Commands
1. sudo apt install certbot 
2. sudo apt install python3-certbot-nginx 
3. sudo certbot --nginx -d app.se-gators.co.in -d api.se-gators.co.in

