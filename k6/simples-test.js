import { sleep } from "k6";
import { SharedArray } from "k6/data";
import http from "k6/http";

export let options = {
  stages: [
    { duration: "30s", target: 2000 }, // simulate ramp-up of traffic from 1 to 200 users over 30 seconds.
    { duration: "1m", target: 2000 }, // stay at 200 users for 1 minutes
    { duration: "30s", target: 0 }, // ramp-down to 0 users
  ],
  thresholds: {
    http_req_duration: ["p(95)<100"], // 95% of requests must complete below 250ms
  },
};

const createId = new SharedArray("body", () => {
  const ids = [];
  for (let i = 0; i < 100; i++) {
    const id = Math.floor(Math.random() * 1000) + 1;
    ids.push(id);
  }
  return ids;
});

export default () => {
  const id = createId[Math.floor(Math.random() * createId.length)];
  http.get("http://localhost:3001/posts/" + id);
  sleep(1);
};
