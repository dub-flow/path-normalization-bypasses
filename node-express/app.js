const express = require('express');
const app = express();
const port = 5000;

// Home route
app.get('/', (req, res) => {
  res.send('Hello World from Express!');
});

// Admin route
app.get('/admin', (req, res) => {
  res.send('Admin area');
});

// Start the server
app.listen(port, () => {
  console.log(`Server running on http://localhost:${port}`);
});
