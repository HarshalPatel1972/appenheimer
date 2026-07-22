import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  scenarios: {
    search_traffic: { executor: 'constant-vus', vus: 70, duration: '2m', exec: 'search' },
    details_traffic: { executor: 'constant-vus', vus: 15, duration: '2m', exec: 'details' },
    admin_traffic: { executor: 'constant-vus', vus: 10, duration: '2m', exec: 'admin' },
    publish_traffic: { executor: 'constant-vus', vus: 5, duration: '2m', exec: 'publish' },
  },
};

export function search() {
  http.get('http://localhost/api/v1/search?q=demo');
  sleep(1);
}

export function details() {
  http.get('http://localhost/api/v1/apps/123');
  sleep(1);
}

export function admin() {
  http.get('http://localhost/api/v1/admin/queue');
  sleep(2);
}

export function publish() {
  // Simulate POST request to publish
  sleep(5);
}
