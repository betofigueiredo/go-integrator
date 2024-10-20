import http from "k6/http";
import { sleep } from "k6";
import { SharedArray } from "k6/data";
import { namesList } from "./namesList.js";

export let options = {
  stages: [
    { duration: "30s", target: 400 }, // simulate ramp-up of traffic from 1 to 200 users over 30 seconds.
    { duration: "1m", target: 400 }, // stay at 200 users for 1 minutes
    { duration: "30s", target: 0 }, // ramp-down to 0 users
  ],
};

const generateUserInfo = (idx) => {
  const sexes = ["MALE", "FEMALE"];
  const roles = ["ADMIN", "MANAGER", "CONTRIBUTOR"];
  const actives = [true, false];
  const firstName = namesList[Math.floor(Math.random() * namesList.length)];
  const lastName = namesList[Math.floor(Math.random() * namesList.length)];
  const name = `${firstName} ${lastName}`;
  return {
    name: name,
    email: `${name.toLowerCase().replace(" ", "")}@mail${idx}.com`,
    phone: String(Math.floor(Math.random() * 1000000000)),
    sex: sexes[Math.floor(Math.random() * sexes.length)],
    birth_date: "1990-01-01",
    role: roles[Math.floor(Math.random() * roles.length)],
    is_active: actives[Math.floor(Math.random() * actives.length)],
  };
};

const createBody = new SharedArray("body", () => {
  const bodies = [];
  for (let i = 0; i < 100; i++) {
    const newBody = generateUserInfo(i);
    bodies.push(newBody);
  }
  return bodies;
});

export default () => {
  const randomBody = createBody[Math.floor(Math.random() * createBody.length)];
  http.post("http://localhost:8000/users", JSON.stringify(randomBody), {
    headers: { "Content-Type": "application/json" },
  });
  sleep(1);
};
