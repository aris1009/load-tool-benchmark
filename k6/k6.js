import http from 'k6/http';
import { check } from 'k6';

export const options = {
    vus: 50,
    duration: '5m'
};

export default function () {
    const res = http.get('http://127.0.0.1:3005');

    const checkRes = check(res, {
        'status is 200': (r) => r.status === 200,
    });
}
