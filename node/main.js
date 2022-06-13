import express from "express";
import {
  generateEvents
} from "./rrule-gen.js";

const app = express()
const port = 3000

app.get('/', (req, res) => {
  const events = generateEvents();
  res.send(events);
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})
