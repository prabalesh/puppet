import { useEffect, useState } from "react";
import { getLanguages, runCode } from "../services/api";
import Editor from "@monaco-editor/react";
import Spinner from "../components/Spinner";
import { Moon, Play, Sun } from "lucide-react";

export default function CodeExecution() {
    const [languages, setLanguages] = useState([]);
    const [selectedLanguage, setSelectedLanguage] = useState(null);
    const [code, setCode] = useState("// Write your code here");
    const [stdin, setStdin] = useState("");
    const [output, setOutput] = useState({
        stdout: "",
        stderr: "",
        duration: "",
    });

    const [loading, setLoading] = useState(false);
    const [theme, setTheme] = useState("vs-light");

    useEffect(() => {
        getLanguages().then(setLanguages);
    }, []);

    async function handleRun() {
        if (!code.trim()) {
            setOutput("Please enter some code.");
            return;
        }

        const languageId = selectedLanguage ? parseInt(selectedLanguage.id) : undefined;

        setLoading(true);
        setOutput("");

        try {
            const res = await runCode({
                languageId,
                code,
                stdin,
            });

            const result = await res.json();
            const { stdout, stderr, duration } = result;

            setOutput({ stdout, stderr, duration });
        } catch (err) {
            console.log(err);
            setOutput("An error occurred while executing the code.");
        } finally {
            setLoading(false);
        }
    }

    return (
        <div className="max-w-7xl mx-auto px-4 py-8">
            <div className="flex justify-between items-center mb-6">
                <h2 className="text-3xl font-bold">Code Runner</h2>
            </div>

            <div className="flex flex-col gap-6">
                <div className="flex gap-4 flex-wrap">
                    <div className="flex items-center gap-4">
                        <select
                            value={selectedLanguage?.id || ""}
                            onChange={(e) => {
                                const lang = languages.find(l => l.id.toString() === e.target.value);
                                setSelectedLanguage(lang || null);
                            }}
                            className="px-4 py-2 border rounded-md text-sm"
                        >
                            <option value="">Select a language</option>
                            {languages.map((lang) => (
                                <option key={lang.id} value={lang.id}>
                                    {lang.name} ({lang.version})
                                </option>
                            ))}
                        </select>
                    </div>
                    <div>
                        <button
                            onClick={() => setTheme(theme === "vs-dark" ? "vs-light" : "vs-dark")}
                            className="text-sm px-4 py-2 border rounded-md hover:bg-gray-200"
                            title="Toggle theme"
                        >
                            {theme === "vs-dark" ? <Sun size={18} /> : <Moon size={18} />}
                        </button>
                    </div>
                </div>

                {/* Editor and Output */}
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <div className="flex flex-col h-[500px]">
                        <label className="font-medium mb-2">Code:</label>
                        <div className="flex-1 border rounded-md overflow-hidden">
                            <Editor
                                height="100%"
                                language={getMonacoLanguage(selectedLanguage?.name)}
                                value={code}
                                onChange={(val) => setCode(val || "")}
                                theme={theme}
                                options={{
                                    minimap: { enabled: false },
                                    fontSize: 14,
                                    scrollBeyondLastLine: false,
                                }}
                            />
                        </div>
                    </div>

                    <div className="flex flex-col gap-4">
                        <div>
                            <label className="font-medium mb-2 block">Stdin (Input):</label>
                            <textarea
                                rows="8"
                                value={stdin}
                                onChange={(e) => setStdin(e.target.value)}
                                placeholder="Optional input..."
                                className="p-4 border rounded-md text-sm resize-y overflow-y-auto h-[150px] w-full"
                            />
                        </div>

                        <div className="p-4 border rounded-md bg-gray-100 h-[180px] overflow-y-auto text-sm font-mono whitespace-pre-wrap">
                            {loading ? (
                                <span className="text-gray-500"><Spinner /></span>
                            ) : (
                                <>
                                    {output.duration && (
                                        <div className="text-blue-600 font-semibold mb-2">
                                            Execution time: {output.duration} seconds
                                        </div>
                                    )}
                                    {output.stdout && (
                                        <div className="text-black">
                                            <strong className="block text-green-700 mb-1">STDOUT:</strong>
                                            {output.stdout}
                                        </div>
                                    )}
                                    {output.stderr && (
                                        <div className="text-red-600 mt-2">
                                            <strong className="block mb-1">STDERR:</strong>
                                            {output.stderr}
                                        </div>
                                    )}
                                    {!output.stdout && !output.stderr && !output.duration && (
                                        <span className="text-gray-500">No output.</span>
                                    )}
                                </>
                            )}
                        </div>

                    </div>
                </div>

                {/* Run Button */}
                <div>
                    <button
                        onClick={handleRun}
                        disabled={loading || !selectedLanguage}
                        className={`py-2 px-6 rounded-md text-white transition ${
                            loading
                                ? "bg-purple-400 cursor-not-allowed"
                                : "bg-purple-700 hover:bg-purple-800"
                        }`}
                        title="Run Code"
                    >
                        {loading ? <Spinner size="h-6 w-6" /> : (
                            <span className="flex items-center gap-2">
                                <Play size={20} />
                                Run
                            </span>
                        )}
                    </button>
                </div>
            </div>
        </div>
    );
}

function getMonacoLanguage(languageName = "") {
    const map = {
        javascript: "javascript",
        python: "python",
        java: "java",
        c: "c",
        cpp: "cpp",
        "c++": "cpp",
        "c#": "csharp",
        typescript: "typescript",
        php: "php",
        kotlin: "kotlin",
        ruby: "ruby",
        rust: "rust",
        go: "go",
        swift: "swift",
    };

    const key = languageName.toLowerCase();
    return map[key] || "plaintext";
}
