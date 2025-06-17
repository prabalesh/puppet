import { useEffect, useState } from 'react';
import { getLanguages, deleteLanguage, installLanguage, uninstallLanguage } from '../services/api';

export default function LanguageList() {
  const [languages, setLanguages] = useState([]);

  async function refresh() {
    const data = await getLanguages();
    setLanguages(data);
  }

  useEffect(() => { refresh(); }, []);

  return (
    <div>
        {languages != null ? (
            <>
      {languages.map(lang => (
        <div key={lang.id} style={{ border: '1px solid #ccc', padding: '0.5rem', marginBottom: '0.5rem' }}>
          <p>{lang.name} ({lang.version}) - {lang.installed ? 'Installed' : 'Not installed'}</p>
          <button onClick={() => { deleteLanguage(lang.id).then(refresh); }}>Delete</button>
          {lang.installed ? (
            <button onClick={() => { uninstallLanguage(lang.id).then(refresh); }}>Uninstall</button>
          ) : (
            <button onClick={() => { installLanguage(lang.id).then(refresh); }}>Install</button>
          )}
        </div>
      ))}
      </>
      ):(<p>No languages to show</p>)}
    </div>
  );
}
