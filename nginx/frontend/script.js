const go_url = 'http://localhost:8000/go/sha/';
const node_url = 'http://localhost:8000/node/sha/';


function isRawStringValid(rawString) {
    if (rawString.length < 8)
        return false;
    return true;
}

async function getapi(url) {
    // Storing response
    const response = await fetch(url);
    // Storing data in form of JSON
    var data = await response.json();
    console.log(data);
    return data;
}

async function submitRawString(url, body) {
    const json_body = JSON.stringify(body)
    console.log(`Going to send post request, url=${url} body=${json_body}`);
    const rawResponse = await fetch(url, {
        method: "POST",
        headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
        },
        body: json_body,
    });

    const content = await rawResponse.json();
    return content;
}

const node_btn = document.querySelector('#rawStringSubmitButton');

node_btn.addEventListener('click', (event) => {
    const rawString = document.querySelector("#rawStringInput").value;
    console.log(rawString);
    if (isRawStringValid(rawString)) {
        const response = submitRawString(node_url, { raw_string: rawString });
        console.log(response);
    } else {
        console.log('Input should at least be 8 characters.');
    }
});