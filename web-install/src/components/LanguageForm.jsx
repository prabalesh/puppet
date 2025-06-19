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
        <form onSubmit={handleSubmit} style={{ marginBottom: "1rem" }}>
            <div>
                <label htmlFor="lname">Language Name</label>
                <input
                    id="lname"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    placeholder="Language name"
                    required
                />
            </div>
            <div>
                <label htmlFor="lversion">Language Version</label>
                <input
                    id="lversion"
                    value={version}
                    onChange={(e) => setVersion(e.target.value)}
                    placeholder="Version"
                    required
                />
            </div>
            <div>
                <label htmlFor="limagename">Language Image Name</label>
                <input
                    id="limagename"
                    value={imageName}
                    onChange={(e) => setImageName(e.target.value)}
                    placeholder="Docker image"
                    required
                />
            </div>
            <div>
                <label htmlFor="lfilename">Source File Name</label>
                <input
                    id="lfilename"
                    value={fileName}
                    onChange={(e) => setFileName(e.target.value)}
                    placeholder="e.g. main.cpp"
                    required
                />
            </div>
            <div>
                <label htmlFor="lcompilecmd">Compile Command</label>
                <input
                    id="lcompilecmd"
                    value={compileCommand}
                    onChange={(e) => setCompileCommand(e.target.value)}
                    placeholder="e.g. g++ -o main main.cpp"
                />
            </div>
            <div>
                <label htmlFor="lruncommand">Run Command</label>
                <input
                    id="lruncommand"
                    value={runCommand}
                    onChange={(e) => setRunCommand(e.target.value)}
                    placeholder="e.g. ./main"
                    required
                />
            </div>
            <button type="submit">Add Language</button>
        </form>
    );
}
