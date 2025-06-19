import { useEffect, useState } from "react";
import { getLanguages } from  "../services/api";

export default function CodeExecution() {
  const [languages, setLanguages] = useState([]);
  const [selectedLanguage, setSelectedLanguage] = useState("");
  const [code, setCode] = useState("");
  const [stdin, setStdin] = useState("");
  const [output, setOutput] = useState("");

  useEffect(() => {
    getLanguages().then(setLanguages);
  }, []);

  async function handleRun() {
    const res = await fetch("http://localhost:8080/api/executions", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        languageId: parseInt(selectedLanguage),
        code,
        stdin,
      }),
    });

    const result = await res.text();
    setOutput(result);
  }

  return (
    <div>
      <h2>Run Code</h2>

      <label>
        Language:
        <select
          value={selectedLanguage}
          onChange={(e) => setSelectedLanguage(e.target.value)}
        >
          <option value="">-- Select --</option>
          {languages.map((lang) => (
            <option key={lang.id} value={lang.id}>
              {lang.name}
            </option>
          ))}
        </select>
      </label>

      <div>
        <label>Code:</label>
        <textarea
          rows="10"
          cols="80"
          value={code}
          onChange={(e) => setCode(e.target.value)}
          placeholder="Write your code here..."
        />
      </div>

      <div>
        <label>Stdin:</label>
        <textarea
          rows="4"
          cols="80"
          value={stdin}
          onChange={(e) => setStdin(e.target.value)}
          placeholder="Optional input..."
        />
      </div>

      <button onClick={handleRun}>Run</button>

      {output && (
        <div style={{ marginTop: "1rem" }}>
          <h3>Output:</h3>
          <pre>{output}</pre>
        </div>
      )}
    </div>
  );
}
