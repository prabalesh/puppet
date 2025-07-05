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
                <div className="grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-3 gap-4">
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
