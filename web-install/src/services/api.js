const BASE_URL = "http://localhost:8080/api";

export async function getLanguages() {
    const res = await fetch(`${BASE_URL}/languages`);
    return res.json();
}

export async function addLanguage(data) {
    await fetch(`${BASE_URL}/languages`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    });
}

export async function deleteLanguage(id) {
    const res = await fetch(`${BASE_URL}/languages/${id}`, { method: "DELETE" });
    return res.json();
}

export async function installLanguage(id) {
    const res = await fetch(`${BASE_URL}/languages/${id}/installations`, { method: "POST" });
    return res.json();
}

export async function uninstallLanguage(id) {
    const res = await fetch(`${BASE_URL}/languages/${id}/installations`, { method: "DELETE" });
    return res.json();
}

export async function getJobStatus(jobId) {
    const res = await fetch(`${BASE_URL}/jobs/installations/${jobId}/status`);
    return res.json();
}

export async function runCode(data) {
    const res = await fetch(`${BASE_URL}/executions`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data)
    });

    return res;
}
