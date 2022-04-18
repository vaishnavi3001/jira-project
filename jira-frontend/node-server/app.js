#!/usr/bin/env node

const express = require('express')
let path = require('path');
const app = express()
const port = 8500


app.use('/', express.static(path.join(__dirname,'assets', 'jira-frontend')));
app.use('/*', (req, res) => res.sendFile(path.join(__dirname,'assets', 'jira-frontend','index.html')));


app.listen(port, () => {
  console.log(`Frontend app listening on port ${port}`)
})