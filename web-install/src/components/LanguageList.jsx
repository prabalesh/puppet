import { useEffect, useState } from 'react';
import {
    getLanguages
} from '../services/api';
import LanguageItem from './LanguageItem';

export default function LanguageList() {
    const [languages, setLanguages] = useState([]);

    async function refresh() {
        const data = await getLanguages();
        setLanguages(data);
    }

    useEffect(() => {
        refresh();
    }, []);

    return (
        <>
            {languages != null && languages.length > 0 ? (
                <div className="grid grid-cols-3 gap-2">
                    {languages.map((lang) => (
                        <LanguageItem key={lang.id} lang={lang} refresh={refresh} />
                    ))}
                </div>
            ) : (
                <p>No languages to show</p>
            )}
        </>
    );
}
