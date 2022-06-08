import smu from "benchmark";
const { Suite } = smu;
import {
  projectDailies,
  projectMultipleRulesDaily,
  projectMultipleRulesDailyOld,
  projectMultipleRulesWeekly,
  projectMultipleRulesWeeklyOld,
  projectWeekly,
} from "./rrule-gen.js";

const suite = new Suite();

suite
  // .add(
  //   "rrule-gen",
  //   () => {
  //     generateEvents();
  //   }
  //   // { minSamples: 1000 }
  // )
  .add("projectDailies", () => {
    projectDailies();
  })
  .add("projectMultipleRulesDaily", () => {
    projectMultipleRulesDaily();
  })
  .add("projectMultipleRulesDailyOld", () => {
    projectMultipleRulesDailyOld();
  })
  .add("projectWeekly", () => {
    projectWeekly();
  })
  .add("projectMultipleRulesWeekly", () => {
    projectMultipleRulesWeekly();
  })
  .add("projectMultipleRulesWeeklyOld", () => {
    projectMultipleRulesWeeklyOld();
  })
  .on("cycle", function (event) {
    console.log(String(event.target));
  })
  .run({ async: false });

// console.log(projectDailies().length);
// console.log(projectMultipleRulesDaily().length);
// console.log(projectMultipleRulesDailyOld().length);
// console.log(projectWeekly().length);
// console.log(projectMultipleRulesWeekly().length);
// console.log(projectMultipleRulesWeeklyOld().length);
