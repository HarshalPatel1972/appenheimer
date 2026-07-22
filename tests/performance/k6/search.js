import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 50 },
    { duration: '1m', target: 500 },
    { duration: '1m', target: 1000 },
    { duration: '30s', target: 0 },
  ],
  thresholds: {
    http_req_duration: ['p(95)<150', 'p(99)<300'],
    http_req_failed: ['rate<0.001'],
  },
};

export default function () {
  const res = http.get('http://localhost/api/v1/search?q=test');
  check(res, {
    'status is 200': (r) => r.status === 200,
  });
  sleep(1);
}
