import { useState } from "react";
import { addLanguage } from "../services/api";

export default function LanguageForm({ onSuccess }) {
    const [name, setName] = useState("");
    const [version, setVersion] = useState("");
    const [imageName, setImageName] = useState("");
    const [fileName, setFileName] = useState("");
    const [compileCommand, setCompileCommand] = useState("");
    const [runCommand, setRunCommand] = useState("");

    async function handleSubmit(e) {
        e.preventDefault();
        await addLanguage({
            name,
            version,
            image_name: imageName,
            file_name: fileName,
            compile_command: compileCommand,
            run_command: runCommand,
        });
        onSuccess();
        // Reset form fields
        setName("");
        setVersion("");
        setImageName("");
        setFileName("");
        setCompileCommand("");
        setRunCommand("");
    }

    return (
        <div className="m-4 w-100">
            <div className="mb-4 text-2xl font-bold">
                <h1>Add Language</h1>
            </div>
            <form onSubmit={handleSubmit} className="flex flex-col gap-2">
                <div className="flex flex-col">
                    <label htmlFor="lname">Language Name</label>
                    <input
                        id="lname"
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        placeholder="Language name"
                        className="p-2 border rounded-sm"
                        required
                    />
                </div>
                <div className="flex flex-col">
                    <label htmlFor="lversion">Language Version</label>
                    <input
                        id="lversion"
                        value={version}
                        onChange={(e) => setVersion(e.target.value)}
                        placeholder="Version"
                        className="p-2 border rounded-sm"
                        required
                    />
                </div>
                <div className="flex flex-col">
                    <label htmlFor="limagename">Language Image Name</label>
                    <input
                        id="limagename"
                        value={imageName}
                        onChange={(e) => setImageName(e.target.value)}
                        placeholder="Docker image"
                        className="p-2 border rounded-sm"
                        required
                    />
                </div>
                <div className="flex flex-col">
                    <label htmlFor="lfilename">Source File Name</label>
                    <input
                        id="lfilename"
                        value={fileName}
                        onChange={(e) => setFileName(e.target.value)}
                        placeholder="e.g. main.cpp"
                        className="p-2 border rounded-sm"
                        required
                    />
                </div>
                <div className="flex flex-col">
                    <label htmlFor="lcompilecmd">Compile Command</label>
                    <input
                        id="lcompilecmd"
                        value={compileCommand}
                        onChange={(e) => setCompileCommand(e.target.value)}
                        placeholder="e.g. g++ -o main main.cpp"
                        className="p-2 border rounded-sm"
                    />
                </div>
                <div className="flex flex-col">
                    <label htmlFor="lruncommand">Run Command</label>
                    <input
                        id="lruncommand"
                        value={runCommand}
                        onChange={(e) => setRunCommand(e.target.value)}
                        placeholder="e.g. ./main"
                        className="p-2 border rounded-sm"
                        required
                    />
                </div>
                <button type="submit" className="w-100 py-2 rounded bg-slate-500 text-white hover:bg-slate-600">Add Language</button>
            </form>
        </div>
    );
}
