const express = require('express');
const axios = require('axios');
const bodyParser = require('body-parser');

const app = express();
const port = 3000;

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.get('/', (req, res) => {
  res.sendFile(__dirname + '/index.html');
});

app.get('/greet', async (req, res) => {
  try {
    const response = await axios.get('http://localhost:8080/greet');
    res.send(response.data);
  } catch (error) {
    res.status(500).send('Error fetching greeting');
  }
});

app.post('/book', async (req, res) => {
  const booking = {
    firstName: req.body.firstName,
    lastName: req.body.lastName,
    email: req.body.email,
    userTickets: parseInt(req.body.userTickets),
  };

  try {
    const response = await axios.post('http://localhost:8080/book', booking);
    res.send(response.data);
  } catch (error) {
    res.status(500).send('Error booking tickets');
  }
});

app.get('/bookings', async (req, res) => {
  try {
    const response = await axios.get('http://localhost:8080/bookings');
    res.send(response.data);
  } catch (error) {
    res.status(500).send('Error fetching bookings');
  }
});

app.listen(port, () => {
  console.log(`Frontend server running at http://localhost:${port}`);
});
