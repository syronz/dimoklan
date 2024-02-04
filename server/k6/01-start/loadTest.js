import http from 'k6/http';
import { check } from 'k6';

export let options = {
  stages: [
    { duration: '10s', target: 100 }, // Stage 1: Ramp up to 50 virtual users 
    { duration: '30s', target: 100 }, // Stage 2: Maintain 100 virtual users 
    { duration: '10s', target: 0 },   // Stage 3: Ramp down to 0 virtual users 
  ],
};

export default function () {


    // Generate a random string for the 'name' field
    const name = 'Name_' + Math.floor(Math.random() * 10000000000);

    // Generate a random string for the 'username' field
    const username = 'User_' + Math.floor(Math.random() * 10000000000);

    // Generate a random password
    const password = 'Pass_' + Math.floor(Math.random() * 1000);


    // Construct the payload
    const payload = {
        name: name,
        username: username,
        password: password,
        language: 'ku',
        email: name,
    };

    // Define the headers
    const headers = {
        'Content-Type': 'application/json',
    };

    // Make a POST request to the endpoint
    const res = http.post('http://127.0.0.1:3000/users', JSON.stringify(payload), { headers: headers });
    // const res = http.get('http://127.0.0.1:3000/');

    // Check if the request was successful
    check(res, {
        'is status 200': (r) => r.status === 200,
    });
}
