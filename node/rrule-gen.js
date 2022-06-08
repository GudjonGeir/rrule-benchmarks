import smu from "rrule";
const { RRuleSet, RRule } = smu;

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

export function generateEvents() {
  return [
    ...rule1.between(
      new Date(Date.UTC(2021, 7, 1)),
      new Date(Date.UTC(2022, 7, 1))
    ),
    ...rule2.between(
      new Date(Date.UTC(2021, 7, 1)),
      new Date(Date.UTC(2022, 7, 1))
    ),
    ...rule3.between(
      new Date(Date.UTC(2021, 7, 1)),
      new Date(Date.UTC(2022, 7, 1))
    ),
  ];
}

function generateRandomDate(interval) {
  const startDate = new Date(2002,1,1).valueOf();
  const timestamp = Math.round(Math.random() * startDate);
  return [new Date(timestamp), new Date(timestamp + interval)];
}

export function projectMaxExpectedEvents() {

  const iterations = 1000000;
  let res = [];

  for (let i = 0; i < iterations; i++) {
    const [min, max] = generateRandomDate(60 * 60);
    res = [...res, rule1.between(min, max,)];
  }

  return res;
}

