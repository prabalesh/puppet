import { useState } from "react";
import { addLanguage } from "../services/api";

export default function LanguageForm({ onSuccess }) {
    const [name, setName] = useState("");
    const [version, setVersion] = useState("");
    const [imageName, setImageName] = useState("");

    async function handleSubmit(e) {
        e.preventDefault();
        await addLanguage({ name, version, image_name: imageName });
        onSuccess();
        setName("");
        setVersion("");
        setImageName("");
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
            <button type="submit">Add Language</button>
        </form>
    );
}
