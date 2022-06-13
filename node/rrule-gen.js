import smu from "rrule";
const { RRuleSet, RRule } = smu;

export function generateEvents() {
  // 3 days a week
  const rule1 = new RRule(
    {
      freq: RRule.WEEKLY,
      dtstart: new Date(Date.UTC(2021, 1, 1, 10, 30)),
      interval: 1,
      byweekday: [RRule.MO, RRule.WE, RRule.FR],
    },
    true
  );
  // Daily
  const rule2 = new RRule(
    {
      freq: RRule.DAILY,
      dtstart: new Date(Date.UTC(2021, 1, 1, 10, 30)),
    },
    true
  );

  // With exclusions and byweekday
  const rule3 = new RRuleSet(true);
  rule3.rrule(
    new RRule({
      freq: RRule.WEEKLY,
      dtstart: new Date(Date.UTC(2021, 1, 1, 17, 30)),
      interval: 1,
    })
  );

  rule3.exdate(new Date(Date.UTC(2021, 7, 16, 17, 30)));
  rule3.exdate(new Date(Date.UTC(2021, 8, 27, 17, 30)));
  rule3.rdate(new Date(Date.UTC(2021, 7, 5, 10, 30)));
  rule3.rdate(new Date(Date.UTC(2021, 7, 16, 10, 30)));
  rule3.rdate(new Date(Date.UTC(2021, 8, 27, 10, 30)));

  return [
    ...rule1.between(
      new Date(Date.UTC(2021, 7, 1)),
      new Date(Date.UTC(2021, 7, 7, 23, 59, 59)),
      true
    ),
    ...rule2.between(
      new Date(Date.UTC(2021, 7, 1)),
      new Date(Date.UTC(2021, 7, 7, 23, 59, 59)),
      true
    ),
    ...rule3.between(
      new Date(Date.UTC(2021, 7, 1)),
      new Date(Date.UTC(2021, 7, 7, 23, 59, 59)),
      true
    ),
  ];
}

// function generateRandomDate(interval) {
//   const startDate = new Date(2002, 1, 1).valueOf();
//   const timestamp = Math.round(Math.random() * startDate);
//   return [new Date(timestamp), new Date(timestamp + interval)];
// }

// export function projectMaxExpectedEvents() {
//   const iterations = 1000000;
//   let res = [];

//   for (let i = 0; i < iterations; i++) {
//     const [min, max] = generateRandomDate(60 * 60);
//     res = [...res, rule1.between(min, max)];
//   }

//   return res;
// }

const numberOfEvents = 10000;

export function projectDailies() {
  const dtstart = new Date(Date.UTC(2021, 1, 1, 10, 30));
  const rule = new RRule(
    {
      freq: RRule.DAILY,
      dtstart: dtstart,
    },
    true
  );

  const dtend = new Date(dtstart);
  dtend.setDate(dtend.getDate() + numberOfEvents);

  return rule.between(dtstart, dtend, true);
}

export function projectMultipleRulesDaily() {
  const dtstart = new Date(Date.UTC(2021, 1, 1, 10, 30));
  const rule = new RRule(
    {
      freq: RRule.DAILY,
      dtstart: dtstart,
    },
    true
  );
  const iterations = numberOfEvents;
  let res = [];

  for (let i = 0; i < iterations; i++) {
    res = [...res, rule.between(dtstart, dtstart, true)];
  }
  return res;
}

export function projectMultipleRulesDailyOld() {
  const dtstart = new Date(Date.UTC(2021, 1, 1, 10, 30));
  const rule = new RRule(
    {
      freq: RRule.DAILY,
      dtstart: new Date(Date.UTC(2019, 1, 1, 10, 30)),
    },
    true
  );

  const iterations = numberOfEvents;
  let res = [];

  for (let i = 0; i < iterations; i++) {
    res = [...res, rule.between(dtstart, dtstart, true)];
  }
  return res;
}

export function projectWeekly() {
  const dtstart = new Date(Date.UTC(2021, 1, 1, 10, 30));
  const rule = new RRule(
    {
      freq: RRule.WEEKLY,
      dtstart: dtstart,
      interval: 1,
      byweekday: [RRule.MO, RRule.WE, RRule.FR],
    },
    true
  );

  const dtend = new Date(dtstart);
  dtend.setDate(dtend.getDate() + (numberOfEvents / 3) * 7);

  return rule.between(dtstart, dtend, true);
}

export function projectMultipleRulesWeekly() {
  const dtstart = new Date(Date.UTC(2021, 1, 1, 10, 30));
  const rule = new RRule(
    {
      freq: RRule.WEEKLY,
      dtstart: dtstart,
      interval: 1,
      byweekday: [RRule.MO, RRule.WE, RRule.FR],
    },
    true
  );
  const iterations = numberOfEvents / 3;

  const dtend = new Date(dtstart);
  dtend.setDate(dtend.getDate() + 6);

  let res = [];

  for (let i = 0; i < iterations; i++) {
    res = [...res, ...rule.between(dtstart, dtend, true)];
  }
  return res;
}

export function projectMultipleRulesWeeklyAsync() {
  console.log("A");
  return new Promise((resolve) => {
    const dtstart = new Date(Date.UTC(2021, 1, 1, 10, 30));
    const rule = new RRule(
      {
        freq: RRule.WEEKLY,
        dtstart: dtstart,
        interval: 1,
        byweekday: [RRule.MO, RRule.WE, RRule.FR],
      },
      true
    );
    const iterations = numberOfEvents / 3;

    const dtend = new Date(dtstart);
    dtend.setDate(dtend.getDate() + 6);

    let res = [];

    for (let i = 0; i < iterations; i++) {
      res = [...res, ...rule.between(dtstart, dtend, true)];
    }
    console.log("B");
    return resolve(res);
  });
}

export function projectMultipleRulesWeeklyOld() {
  const dtstart = new Date(Date.UTC(2021, 1, 1, 10, 30));
  const rule = new RRule(
    {
      freq: RRule.WEEKLY,
      dtstart: new Date(Date.UTC(2019, 1, 1, 10, 30)),
      interval: 1,
      byweekday: [RRule.MO, RRule.WE, RRule.FR],
    },
    true
  );
  const iterations = numberOfEvents / 3;

  const dtend = new Date(dtstart);
  dtend.setDate(dtend.getDate() + 6);

  let res = [];

  for (let i = 0; i < iterations; i++) {
    res = [...res, ...rule.between(dtstart, dtend, true)];
  }
  return res;
}

// console.log("projectDailies: ", projectDailies().length);
// console.log("projectMultipleRulesDaily: ", projectMultipleRulesDaily().length);
// console.log(
//   "projectMultipleRulesDailyOld: ",
//   projectMultipleRulesDailyOld().length
// );
// console.log("projectWeekly: ", projectWeekly().length);
// console.log(
//   "projectMultipleRulesWeekly: ",
//   projectMultipleRulesWeekly().length
// );
// console.log(
//   "projectMultipleRulesWeeklyOld: ",
//   projectMultipleRulesWeeklyOld().length
// );
