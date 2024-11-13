import http from 'k6/http';

export const options = {
    vus: 500,
    duration: '5s',
    rps: 500,
}

export default function () {
    const bannerId = "634dfb86-f492-4f12-b524-2b3d35f2c5a3"
    const url = `http://localhost:8080/api/v1/counter/${bannerId}`;

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.get(url, params);
}