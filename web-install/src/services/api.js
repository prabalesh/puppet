const BASE_URL = "http://localhost:8080/api/languages";

export async function getLanguages() {
    const res = await fetch(BASE_URL);
    return res.json();
}

export async function addLanguage(data) {
    await fetch(BASE_URL, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    });
}

export async function deleteLanguage(id) {
    await fetch(`${BASE_URL}/${id}`, { method: "DELETE" });
}

export async function installLanguage(id) {
    await fetch(`${BASE_URL}/${id}/installations`, { method: "POST" });
}

export async function uninstallLanguage(id) {
    await fetch(`${BASE_URL}/${id}/installations`, { method: "DELETE" });
}
