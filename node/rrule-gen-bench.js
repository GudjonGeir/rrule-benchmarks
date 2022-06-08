import smu from "benchmark";
const { Suite } = smu;
import {generateEvents,  projectMaxExpectedEvents } from "./rrule-gen.js";

const suite = new Suite();

suite
  .add(
    "rrule-gen",
    () => {
      generateEvents();
    }
    // { minSamples: 1000 }
  )
.add("projections", () => {projectMaxExpectedEvents()})
  .on("cycle", function (event) {
    console.log(String(event.target));
  })
  .run({ async: false });
