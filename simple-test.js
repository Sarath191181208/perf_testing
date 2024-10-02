import http from "k6/http";
import { check } from "k6";

export const options = {
  stages: [
    { duration: "30s", target: 20 },
    { duration: "1m", target: 10 },
    { duration: "30s", target: 0 },
    { duration: "30s", target: 100 },
    { duration: "1m", target: 50 },
    { duration: "30s", target: 0 },
  ],
};

export default function() {
  const url = "http://127.0.0.1:3000/register";
  const payload = JSON.stringify({
    name: "sarath",
    email: "vssarathc04@gmail.com",
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  const response = http.post(url, payload, params);

  check(response, {
    "is status 201": (r) => r.status === 201,
    "response body is not empty": (r) => r.body.length > 0,
  });
}
