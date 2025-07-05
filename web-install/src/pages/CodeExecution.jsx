import { useEffect, useState } from "react";
import { getLanguages, runCode } from "../services/api";

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
        const res = await runCode({
            languageId: parseInt(selectedLanguage),
            code,
            stdin,
        });

        const result = await res.json();
        setOutput(result.stdout);
    }

    return (
        <div>
            <h2 className="text-2xl font-bold">Run Code</h2>

            <div className="flex flex-col gap-4">
                <div className="my-4 flex items-center gap-2">
                    <label>Language</label>
                    <select
                        value={selectedLanguage}
                        onChange={(e) => setSelectedLanguage(e.target.value)}
                        className="px-4 py-1 border rounded"
                    >
                        <option value="">Select</option>
                        {languages.map((lang) => (
                            <option key={lang.id} value={lang.id}>
                                {lang.name}
                            </option>
                        ))}
                    </select>
                </div>


                <div className="flex flex-col">
                    <label>Code:</label>
                    <textarea
                        rows="10"
                        cols="80"
                        value={code}
                        onChange={(e) => setCode(e.target.value)}
                        placeholder="Write your code here..."
                        className="p-2 border rounded"
                    />
                </div>

                <div className="flex flex-col">
                    <label>Stdin:</label>
                    <textarea
                        rows="4"
                        cols="80"
                        value={stdin}
                        onChange={(e) => setStdin(e.target.value)}
                        placeholder="Optional input..."
                        className="p-2 border rounded"
                    />
                </div>

                <button onClick={handleRun} className="py-2 px-4 bg-slate-500 text-white hover:bg-slate-600 rounded">Run</button>
            </div>

            {output && (
                <div style={{ marginTop: "1rem" }}>
                    <div>
                        <h3>Output:</h3>
                    </div>
                    {output != "" && (
                        <div className="p-2 border rounded">
                            <pre>{output}</pre>
                        </div>
                    )}
                </div>
            )}
        </div>
    );
}
