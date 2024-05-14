import { check, sleep } from 'k6';
import http from 'k6/http';

// This is the URL we're going to test, in this case the application server.
const url = "http://assessment:8080";

// The default function is the one that will be run by k6 when it starts.
export default function () {
    const randomName = "https://ifconfig.me";

    // First check a POST to the application server, we'll use the result of the POST to ensure all is well.
    const resPost = http.post(`${url}`, JSON.stringify({ url: randomName }),
        { headers: { 'Content-Type': 'application/json' } });

    // We want to ensure that 201s are returned on a POST, and that the latency is sub-300ms.
    check(resPost, {
        'POST status was 201': (r) => r.status == 201,
        'POST transaction time below 300ms': (r) => r.timings.duration < 300,
    });
    sleep(1);
}
