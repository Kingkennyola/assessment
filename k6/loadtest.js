import { check, sleep } from 'k6';
import http from 'k6/http';

// This is the URL we're going to test, in this case the application server.
const url = `http://${__ENV.HOST}`;

const scrapeURLs = [
    'https://httpstat.us/200',
    'https://httpstat.us/400',
    'https://httpstat.us/401',
    'https://httpstat.us/500',
    'https://httpstat.us/503',
];

// The default function is the one that will be run by k6 when it starts.
export default function () {
    const scrapeURL = scrapeURLs[Math.floor(Math.random() * scrapeURLs.length)];

    // First check a POST to the application server, we'll use the result of the POST to ensure all is well.
    const resPost = http.post(`${url}`, JSON.stringify({ url: scrapeURL }),
        { headers: { 'Content-Type': 'application/json' } });

    // We want to ensure that 201s are returned on a POST, and that the latency is sub-5ms.
    check(resPost, {
        'POST status was 201': (r) => r.status == 201,
        'POST transaction time below 5ms': (r) => r.timings.duration < 5,
    });
    sleep(1);
}
